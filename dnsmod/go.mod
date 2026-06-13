module go.arpabet.com/sprint/dnsmod

go 1.23

toolchain go1.23.4

require (
	github.com/likexian/whois v1.14.8
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.11.1
	go.arpabet.com/glue v1.5.0
	go.arpabet.com/sprint/dns v1.0.0
	go.uber.org/zap v1.28.0
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.arpabet.com/sprint/dns => ../dns
