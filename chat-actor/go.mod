module github.com/jordan-rash/wasmcloud-chat/actor

go 1.18

replace github.com/wasmcloud/interfaces/actor/tinygo => ../../../interfaces/actor/tinygo

require (
	github.com/jordan-rash/wasmcloud-chat/interface v0.0.0-00010101000000-000000000000
	github.com/wasmcloud/interfaces/actor/tinygo v0.0.0-00010101000000-000000000000
	github.com/wasmcloud/interfaces/logging/tinygo v0.0.0-20220909155706-08cf517e404c
	github.com/wasmcloud/interfaces/messaging/tinygo v0.0.0-20220909155706-08cf517e404c
	github.com/wasmcloud/tinygo-msgpack v0.1.4
)

require (
	github.com/wasmcloud/actor-tinygo v0.1.1 // indirect
	github.com/wasmcloud/interfaces v0.0.0-20220909155706-08cf517e404c // indirect
	github.com/wasmcloud/tinygo-cbor v0.1.0 // indirect
)
