package main

import (
	"errors"
	"strings"

	chat "github.com/jordan-rash/wasmcloud-chat/interface"

	actor "github.com/wasmcloud/interfaces/actor/tinygo"
	log "github.com/wasmcloud/interfaces/logging/tinygo"
	messaging "github.com/wasmcloud/interfaces/messaging/tinygo"
	msgpack "github.com/wasmcloud/tinygo-msgpack"
)

func main() {
	me := ChatActor{}
	actor.RegisterHandlers(
		chat.ChatSubscriberHandler(&me),
		messaging.MessageSubscriberHandler(&me),
	)
}

type ChatActor struct{}

func (c *ChatActor) HandleMessage(ctx *actor.Context, arg messaging.SubMessage) error {
	logger := log.NewProviderLogging()
	client := chat.NewProviderChat()
	d := msgpack.NewDecoder(arg.Body)
	msg, err := chat.MDecodeMsg(&d)
	if err != nil {
		return err
	}

	logger.WriteLog(ctx, log.LogEntry{Level: "debug", Text: "GOT A MSG on sub" + arg.Subject})
	logger.WriteLog(ctx, log.LogEntry{Level: "debug", Text: "GOT A MSG" + msg.Value})

	sub := strings.Split(arg.Subject, ".")
	switch len(sub) {
	case 3: // should be a chat message
		client.UpdateChat(ctx, msg)
		return nil
	case 4: // should be a user message with action
		s := chat.UserChannelMsg{
			User:   msg.Owner,
			Action: sub[3],
		}

		client.UpdateUserList(ctx, s)
		return nil
	}
	return errors.New("Unexpected Nats subject")
}

// arg should decode to chat.Msg
func (c *ChatActor) HandleChatMessage(ctx *actor.Context, arg chat.Msg) error {
	logger := log.NewProviderLogging()

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	msg := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(msg)
	enc := &encoder
	arg.MEncode(enc)

	pm := messaging.PubMessage{
		Body:    msg,
		Subject: "wcchat.msg." + arg.Owner.Name,
	}
	client := messaging.NewProviderMessaging()
	err := client.Publish(ctx, pm)
	if err != nil {
		return err
	}

	logger.WriteLog(ctx, log.LogEntry{Level: "debug", Text: "Subject: " + pm.Subject})
	return nil
}

// arg should decode to chat.UserChannelMsg
func (c *ChatActor) HandleUserMessage(ctx *actor.Context, arg chat.Msg) error {
	logger := log.NewProviderLogging()
	logger.WriteLog(ctx, log.LogEntry{Level: "debug", Text: "Inside UserMessageHandler"})

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	msg := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(msg)
	enc := &encoder
	arg.MEncode(enc)

	subj := ""

	switch arg.Value {
	case "join":
		subj = "wcchat.msg." + arg.Owner.Name + ".join"
	case "leave":
		subj = "wcchat.msg." + arg.Owner.Name + ".leave"
	default:
		return errors.New("Invalid user action")
	}

	pm := messaging.PubMessage{
		Body:    []byte(arg.Owner.Name + "|" + arg.Owner.Color),
		Subject: subj,
	}

	logger.WriteLog(ctx, log.LogEntry{Level: "debug", Text: "Subject: " + subj})

	client := messaging.NewProviderMessaging()
	err := client.Publish(ctx, pm)
	if err != nil {
		return err
	}

	return nil
}
