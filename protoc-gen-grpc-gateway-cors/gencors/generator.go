package gencors

import (
	"fmt"
	"github.com/pkg/errors"
	"go/format"
	"path"
	"path/filepath"
	"strings"

	cors_options "github.com/dan-compton/grpc-gateway-cors/options"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	gen "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/generator"
)

var (
	errNoTargetService = errors.New("no target service defined in the file")
)

type generator struct {
	reg         *descriptor.Registry
	baseImports []descriptor.GoPackage
}

// New returns a new generator which generates grpc gateway files.
func New(reg *descriptor.Registry) gen.Generator {
	var imports []descriptor.GoPackage
	for _, pkgpath := range []string{
		"github.com/grpc-ecosystem/grpc-gateway/runtime",
	} {
		pkg := descriptor.GoPackage{
			Path: pkgpath,
			Name: path.Base(pkgpath),
		}
		if err := reg.ReserveGoPackageAlias(pkg.Name, pkg.Path); err != nil {
			for i := 0; ; i++ {
				alias := fmt.Sprintf("%s_%d", pkg.Name, i)
				if err := reg.ReserveGoPackageAlias(alias, pkg.Path); err != nil {
					continue
				}
				pkg.Alias = alias
				break
			}
		}
		imports = append(imports, pkg)
	}
	return &generator{reg: reg, baseImports: imports}
}

type ServiceCORS struct {
	Name string
	Opts *cors_options.CORS
}

func (g *generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	var files []*plugin.CodeGeneratorResponse_File
	for _, file := range targets {
		services := file.GetService()
		for _, targetSVC := range services {
			if !proto.HasExtension(targetSVC.Options, cors_options.E_Cors) {
				return nil, nil
			}

			svcExt, err := proto.GetExtension(targetSVC.Options, cors_options.E_Cors)
			if err != nil {
				return nil, err
			}

			opts, ok := svcExt.(*cors_options.CORS)
			if !ok {
				return nil, errors.Errorf("extension is %T; want a CORS object", svcExt)
			}

			code, err := applyTemplate(file, &ServiceCORS{targetSVC.GetName(), opts})
			if err != nil {
				if err == errNoTargetService {
					continue
				}
				return nil, err
			}

			formatted, err := format.Source([]byte(code))
			if err != nil {
				return nil, errors.Wrapf(err, "formatting %s", code)
			}

			name := strings.ToLower(file.GetName())
			ext := filepath.Ext(name)
			base := strings.TrimSuffix(name, ext)
			output := fmt.Sprintf("%s.pb.gw.cors.go", base)
			files = append(files, &plugin.CodeGeneratorResponse_File{
				Name:    proto.String(output),
				Content: proto.String(string(formatted)),
			})
		}
	}
	return files, nil
}
