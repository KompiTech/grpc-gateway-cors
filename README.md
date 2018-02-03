# github.com/dan-compton/grpc-gateway-cors

This repository contains a protoc plugin, options proto, and cors implementation which allows
for CORS options to be specified within a grpc-spec.  It is intended for usage in [github.com/grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
enabled services.  


The underlying CORS implementation is [github.com/rs/cors](https://github.com/rs/cors). 

## Installation

`go get -u github.com/dan-compton/grpc-gateway-cors/...`

## Usage

```
syntax = "proto3";
option go_package = "api";

import "google/api/annotations.proto";
import "grpc-gateway-cors/options/cors.proto";
import "api/credentials/credentials.proto";

service ExampleService {
    option (grpc.gateway.cors) = {
        allow_origin: "*";
        allow_method: "PUT";
        allow_method: "GET";
        allow_method: "OPTIONS";
        max_age: 3200;
        allow_credentials: false,
    };

    // CheckHealth is a server health check.
    rpc CheckHealth(HealthCheckRequest) returns (HealthCheckResponse){
        option (google.api.http) = {
            get: "/_status"
        };
    }
}

```

The resultant `example.pb.gw.cors.go` would contain: 

```
func ExampleCORSOptions() []cors.Option {
	opts := []cors.Option{}
	opts = append(opts, cors.AllowedOrigins("*"))
	opts = append(opts, cors.AllowedMethods("GET", "PUT", "OPTIONS"))
	opts = append(opts, cors.AllowedHeaders("*"))
	opts = append(opts, cors.ExposedHeaders())
	opts = append(opts, cors.MaxAge(3200))
	opts = append(opts, cors.OptionsPassthrough(false))
	opts = append(opts, cors.AllowCredentials(true))
	opts = append(opts, cors.Debug(true))
	return opts
}
```

You can then wrap a handler, thereby enabling CORS.

```
cors.Wrap(myhttpHandler, ExampleCORSOptions()...)
```
