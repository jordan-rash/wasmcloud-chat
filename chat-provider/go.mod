module zxcv

go 1.18

require github.com/charmbracelet/bubbletea v0.22.1

require (
	github.com/charmbracelet/bubbles v0.14.0
	github.com/charmbracelet/keygen v0.3.0
	github.com/charmbracelet/lipgloss v0.5.0
	github.com/charmbracelet/wish v0.5.0
	github.com/gliderlabs/ssh v0.3.5
	github.com/google/uuid v1.3.0
	github.com/jordan-rash/wasmcloud-chat/interface v0.0.0-00010101000000-000000000000
	github.com/jordan-rash/wasmcloud-provider v0.0.0-20220901133242-6e3d105801c3
	github.com/sirupsen/logrus v1.9.0
	github.com/wasmcloud/interfaces/core/tinygo v0.0.0-00010101000000-000000000000
	github.com/wasmcloud/tinygo-msgpack v0.1.4
)

require (
	github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/caarlos0/sshmarshal v0.1.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.11.1-0.20220212125758-44cd13922739 // indirect
	github.com/nats-io/jwt/v2 v2.3.0 // indirect
	github.com/nats-io/nats.go v1.16.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/wasmcloud/interfaces/actor/tinygo v0.0.0-00010101000000-000000000000 // indirect
	github.com/wasmcloud/tinygo-cbor v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20220826181053-bd7e27e6170d // indirect
	golang.org/x/sys v0.0.0-20220825204002-c680a09ffe64 // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/jordan-rash/wasmcloud-provider => "/Users/jordan/Library/Mobile Documents/com~apple~CloudDocs/github/jordan-rash/wasmcloud/chat/wasmcloud-provider"

replace github.com/wasmcloud/interfaces/core/tinygo => "/Users/jordan/Library/Mobile Documents/com~apple~CloudDocs/github/jordan-rash/wasmcloud/interfaces/core/tinygo"

replace github.com/wasmcloud/interfaces/actor/tinygo => ../../../interfaces/actor/tinygo
