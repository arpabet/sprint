module go.arpabet.com/sprint/sprintframework

go 1.25.0

//replace go.arpabet.com/glue => ../../go.arpabet.com/glue

//replace go.arpabet.com/sprint/sprint => ../../go.arpabet.com/sprint/sprint

require (
	github.com/fsnotify/fsnotify v1.6.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/hashicorp/go-hclog v1.5.0
	github.com/mailgun/mailgun-go/v4 v4.8.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.11.1
	go.arpabet.com/base62 v1.1.0
	go.arpabet.com/glue v1.5.0
	go.arpabet.com/properties v1.0.0
	go.arpabet.com/sprint/cert v1.0.0
	go.arpabet.com/sprint/dns v1.0.0 // indirect
	go.arpabet.com/sprint/nat v1.0.0
	go.arpabet.com/sprint/sprint v1.0.0
	go.arpabet.com/store v1.1.0
	go.arpabet.com/uuid v1.0.0
	go.uber.org/atomic v1.10.0
	go.uber.org/zap v1.28.0
	golang.org/x/crypto v0.28.0
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.15.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.36.11
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	go.arpabet.com/sprint/sprintpb v1.0.0
	go.arpabet.com/store/providers/badger v1.1.0
	go.arpabet.com/store/providers/bbolt v1.1.0
	go.arpabet.com/store/providers/bolt v1.1.0
	go.arpabet.com/store/providers/mem v1.1.0
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/badger/v4 v4.9.2 // indirect
	github.com/dgraph-io/ristretto/v2 v2.4.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/flatbuffers v25.12.19+incompatible // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jellydator/ttlcache/v3 v3.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.arpabet.com/sprint/certpb v1.0.0 // indirect
	go.etcd.io/bbolt v1.4.3 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/oauth2 v0.9.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/term v0.25.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto v0.0.0-20230303212802-e74f57abe488 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.arpabet.com/sprint/cert => ../cert

replace go.arpabet.com/sprint/certpb => ../certpb

replace go.arpabet.com/sprint/dns => ../dns

replace go.arpabet.com/sprint/nat => ../nat

replace go.arpabet.com/sprint/sprint => ../sprint

replace go.arpabet.com/sprint/sprintpb => ../sprintpb
