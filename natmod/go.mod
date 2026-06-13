module go.arpabet.com/sprint/natmod

go 1.23.0

require (
	github.com/huin/goupnp v1.2.0
	github.com/jackpal/go-nat-pmp v1.0.2
	github.com/pkg/errors v0.9.1
	go.arpabet.com/glue v1.5.0
	go.arpabet.com/sprint/nat v1.0.0
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/sync v0.15.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.arpabet.com/sprint/nat => ../nat
