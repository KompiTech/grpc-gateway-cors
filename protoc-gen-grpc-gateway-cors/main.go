// Command protoc-gen-grpc-gateway-cors is a plugin for Google protocol buffer
// compiler to generate a CORS patterns for GRPC gateway.
// You rarely need to run this program directly. Instead, put this program
// into your $PATH with a name "protoc-gen-grpc-gateway-cors" and run
//   protoc --grpc-gateway-cors_out=output_directory path/to/input.proto
//
// See README.md for more details.
package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dan-compton/grpc-gateway-cors/protoc-gen-grpc-gateway-cors/gencors"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/pkg/errors"
)

var (
	importPrefix = flag.String("import_prefix", "", "prefix to be added to go package paths for imported proto files")
)

func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "reading code generator request")
	}
	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		return nil, errors.Wrap(err, "unmarshaling code generator request")
	}
	return req, nil
}

func main() {
	flag.Parse()

	reg := descriptor.NewRegistry()

	req, err := parseReq(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if req.Parameter != nil {
		for _, p := range strings.Split(req.GetParameter(), ",") {
			spec := strings.SplitN(p, "=", 2)
			if len(spec) == 1 {
				if err := flag.CommandLine.Set(spec[0], ""); err != nil {
					log.Fatal("cannot set flag ", p)
				}
				continue
			}
			name, value := spec[0], spec[1]
			if strings.HasPrefix(name, "M") {
				reg.AddPkgMap(name[1:], value)
				continue
			}
			if err := flag.CommandLine.Set(name, value); err != nil {
				log.Fatal("cannot set flag ", p)
			}
		}
	}

	g := gencors.New(reg)

	reg.SetPrefix(*importPrefix)
	if err := reg.Load(req); err != nil {
		emitError(err)
		return
	}

	var targets []*descriptor.File
	for _, target := range req.FileToGenerate {
		f, err := reg.LookupFile(target)
		if err != nil {
			log.Fatal(err)
		}
		targets = append(targets, f)
	}

	out, err := g.Generate(targets)
	if err != nil {
		emitError(err)
		return
	}
	emitFiles(out)
}

func emitFiles(out []*plugin.CodeGeneratorResponse_File) {
	emitResp(&plugin.CodeGeneratorResponse{File: out})
}

func emitError(err error) {
	emitResp(&plugin.CodeGeneratorResponse{Error: proto.String(err.Error())})
}

func emitResp(resp *plugin.CodeGeneratorResponse) {
	buf, err := proto.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatal(err)
	}
}
