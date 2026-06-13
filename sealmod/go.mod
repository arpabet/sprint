module go.arpabet.com/sprint/sealmod

go 1.25.0

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.11.1
	go.arpabet.com/sprint/seal v1.0.0
	golang.org/x/crypto v0.28.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/sys v0.46.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.arpabet.com/sprint/seal => ../seal
