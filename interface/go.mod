module github.com/jordan-rash/wasmcloud-chat/interface

go 1.19

replace github.com/wasmcloud/interfaces/actor/tinygo => ../../../interfaces/actor/tinygo

require (
	github.com/wasmcloud/interfaces/actor/tinygo v0.0.0-00010101000000-000000000000
	github.com/wasmcloud/tinygo-cbor v0.1.0
	github.com/wasmcloud/tinygo-msgpack v0.1.4
)

require github.com/wasmcloud/interfaces v0.0.0-20220909155706-08cf517e404c // indirect
