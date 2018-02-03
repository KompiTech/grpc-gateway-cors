package gencors

import (
	"bytes"
	"text/template"

	gw_descriptor "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
)

func applyTemplate(file *gw_descriptor.File, opts *ServiceCORS) (string, error) {
	w := bytes.NewBuffer(nil)
	if err := headerTemplate.Execute(w, file); err != nil {
		return "", err
	}
	if err := corsTemplate.Execute(w, opts); err != nil {
		return "", err
	}
	return w.String(), nil
}

var (
	headerTemplate = template.Must(template.New("header").Parse(`
// CODE GENERATEB BY PROTOC-GEN-CORS
// DO NOT MODIFY
package {{ .GoPkg.Name }}
import cors "github.com/dan-compton/grpc-gateway-cors"
`))

	corsTemplate = template.Must(template.New("cors").Parse(`
func {{ .Name }}CORSOptions() []cors.Option {
	opts := []cors.Option{}
	opts = append(opts, cors.AllowedOrigins({{ range .Opts.AllowOrigin }} "{{ . }}", {{ end }}))
	opts = append(opts, cors.AllowedMethods({{ range .Opts.AllowMethod }} "{{ . }}", {{ end }}))
	opts = append(opts, cors.AllowedHeaders({{ range .Opts.AllowHeader }} "{{ . }}", {{ end }}))
	opts = append(opts, cors.ExposedHeaders({{ range .Opts.ExposeHeader }} "{{ . }}", {{ end }}))
	opts = append(opts, cors.MaxAge( {{ .Opts.MaxAge }}))
	opts = append(opts, cors.OptionsPassthrough( {{ .Opts.OptionsPassthrough }}))
	opts = append(opts, cors.AllowCredentials({{ .Opts.AllowCredentials }}))
	opts = append(opts, cors.Debug({{ .Opts.Debug }}))
	return opts
}
`))
)
