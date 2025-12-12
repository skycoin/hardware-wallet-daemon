module github.com/skycoin/hardware-wallet-daemon

go 1.25.1

require (
	github.com/NYTimes/gziphandler v1.1.1
	github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883
	github.com/blang/semver v3.5.1+incompatible
	github.com/go-openapi/errors v0.22.5
	github.com/go-openapi/runtime v0.29.2
	github.com/go-openapi/strfmt v0.25.0
	github.com/go-openapi/swag v0.25.4
	github.com/go-openapi/validate v0.25.1
	github.com/gogo/protobuf v1.3.2
	github.com/rs/cors v1.11.1
	github.com/skycoin/hardware-wallet-go v1.1.1-0.20251212175623-4306ea9c2149
	github.com/skycoin/hardware-wallet-protob v0.0.0-20250805154629-410561e1bc2f
	github.com/skycoin/skycoin v0.28.1-0.20251012182647-a1a88ea0df8f
	github.com/spf13/cobra v1.10.2
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.24.1 // indirect
	github.com/go-openapi/jsonpointer v0.22.4 // indirect
	github.com/go-openapi/jsonreference v0.21.4 // indirect
	github.com/go-openapi/loads v0.23.2 // indirect
	github.com/go-openapi/spec v0.22.2 // indirect
	github.com/go-openapi/swag/cmdutils v0.25.4 // indirect
	github.com/go-openapi/swag/conv v0.25.4 // indirect
	github.com/go-openapi/swag/fileutils v0.25.4 // indirect
	github.com/go-openapi/swag/jsonname v0.25.4 // indirect
	github.com/go-openapi/swag/jsonutils v0.25.4 // indirect
	github.com/go-openapi/swag/loading v0.25.4 // indirect
	github.com/go-openapi/swag/mangling v0.25.4 // indirect
	github.com/go-openapi/swag/netutils v0.25.4 // indirect
	github.com/go-openapi/swag/stringutils v0.25.4 // indirect
	github.com/go-openapi/swag/typeutils v0.25.4 // indirect
	github.com/go-openapi/swag/yamlutils v0.25.4 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/google/gousb v1.1.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/stretchr/objx v0.5.3 // indirect
	go.mongodb.org/mongo-driver v1.17.6 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.39.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.39.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/term v0.38.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// IT IS FORBIDDEN TO USE REPLACE DIRECTIVES

// [error] The go.mod file for the module providing named packages contains one or
//	more replace directives. It must not contain directives that would cause
//	it to be interpreted differently than if it were the main module.

// Uncomment for tests with local sources
//replace github.com/skycoin/hardware-wallet-go => ../hardware-wallet-go
//replace github.com/skycoin/hardware-wallet-protob => ../hardware-wallet-protob
//replace github.com/skycoin/skycoin => ../skycoin

// Below should reflect current versions of the following deps
// To update deps to specific commit hash:
// 1) Uncomment one of the following lines and substituite version with desired commit hash:
//replace github.com/skycoin/skycoin => github.com/skycoin/skycoin v0.28.1-0.20251205225511-c088af7bbed1
//replace github.com/skycoin/hardware-wallet-go => github.com/skycoin/hardware-wallet-go v1.1.1-0.20251212175623-4306ea9c2149
//replace github.com/skycoin/hardware-wallet-protob => github.com/skycoin/hardware-wallet-protob v0.0.0-20250805154629-410561e1bc2f
// 2) Run `go mod tidy && go mod vendor`
// 3) Copy the populated version string to the correct place in require(...) above - replacing the specified version string
// 4) Re-comment the uncommented replace directive above
// 5) Save this file.
// 6) Run `go mod tidy && go mod vendor`
