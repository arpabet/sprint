module go.arpabet.com/sprint/fsmod

go 1.25.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.11.1
	go.arpabet.com/sprint/fs v1.0.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto v0.0.0-20230303212802-e74f57abe488 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.arpabet.com/sprint/fs => ../fs
