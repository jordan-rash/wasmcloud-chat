package main

import (
	"fmt"
	"math/rand"
	"time"

	chat "github.com/jordan-rash/wasmcloud-chat/interface"
	msgpack "github.com/wasmcloud/tinygo-msgpack"
)

func randomColor() string {
	hex := func(num int) string {
		h := fmt.Sprintf("%x", num)
		if len(h) == 1 {
			h = "0" + h
		}
		return h
	}

	rand.Seed(time.Now().UnixNano())
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	h := "#" + hex(red) + hex(green) + hex(blue)
	return h
}

func encodeMsg(m chat.Msg) []byte {
	var sizeri msgpack.Sizer
	size_enci := &sizeri
	m.MEncode(size_enci)
	buf := make([]byte, sizeri.Len())
	encoderi := msgpack.NewEncoder(buf)
	enci := &encoderi
	m.MEncode(enci)
	return buf
}
