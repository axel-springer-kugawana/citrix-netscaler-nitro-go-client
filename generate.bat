go fmt ./generator
go build ./generator

generator.exe -i %GOPATH%/src/github.com/doubret/citrix-netscaler-nitro-11-yaml-specs/yml

go fmt ./nitro
go fmt ./tests/resources
go fmt ./tests/bindings
