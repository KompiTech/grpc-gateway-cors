GO_PLUGIN_PKG=github.com/golang/protobuf/protoc-gen-go
PKGMAP=Mgoogle/protobuf/descriptor.proto=$(GO_PLUGIN_PKG)/descriptor

CORS_PROTO=options/cors.proto
CORS_GO=$(CORS_PROTO:.proto=.pb.go)

$(CORS_GO):
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
 		--go_out=$(PKGMAP):. \
		$(CORS_PROTO)
