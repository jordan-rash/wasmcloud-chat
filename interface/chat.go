package chat

import (
	// org.jordanrash.chat
	actor "github.com/wasmcloud/interfaces/actor/tinygo" //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"
	msgpack "github.com/wasmcloud/tinygo-msgpack"
)

type Msg struct {
	Id    string `json:"id"`
	Owner *User  `json:"owner"`
	Time  string `json:"time"`
	Value string `json:"value"`
}

// MEncode serializes a Msg using msgpack
func (o *Msg) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("id")
	encoder.WriteString(o.Id)
	encoder.WriteString("owner")
	if o.Owner == nil {
		encoder.WriteNil()
	} else {
		o.Owner.MEncode(encoder)
	}
	encoder.WriteString("time")
	encoder.WriteString(o.Time)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// MDecodeMsg deserializes a Msg using msgpack
func MDecodeMsg(d *msgpack.Decoder) (Msg, error) {
	var val Msg
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "id":
			val.Id, err = d.ReadString()
		case "owner":
			fval, err := MDecodeUser(d)
			if err != nil {
				return val, err
			}
			val.Owner = &fval
		case "time":
			val.Time, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Msg using cbor
func (o *Msg) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("id")
	encoder.WriteString(o.Id)
	encoder.WriteString("owner")
	if o.Owner == nil {
		encoder.WriteNil()
	} else {
		o.Owner.CEncode(encoder)
	}
	encoder.WriteString("time")
	encoder.WriteString(o.Time)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// CDecodeMsg deserializes a Msg using cbor
func CDecodeMsg(d *cbor.Decoder) (Msg, error) {
	var val Msg
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "id":
			val.Id, err = d.ReadString()
		case "owner":
			fval, err := CDecodeUser(d)
			if err != nil {
				return val, err
			}
			val.Owner = &fval
		case "time":
			val.Time, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type User struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

// MEncode serializes a User using msgpack
func (o *User) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("color")
	encoder.WriteString(o.Color)
	encoder.WriteString("name")
	encoder.WriteString(o.Name)

	return encoder.CheckError()
}

// MDecodeUser deserializes a User using msgpack
func MDecodeUser(d *msgpack.Decoder) (User, error) {
	var val User
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "color":
			val.Color, err = d.ReadString()
		case "name":
			val.Name, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a User using cbor
func (o *User) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("color")
	encoder.WriteString(o.Color)
	encoder.WriteString("name")
	encoder.WriteString(o.Name)

	return encoder.CheckError()
}

// CDecodeUser deserializes a User using cbor
func CDecodeUser(d *cbor.Decoder) (User, error) {
	var val User
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "color":
			val.Color, err = d.ReadString()
		case "name":
			val.Name, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type UserChannelMsg struct {
	Action string `json:"action"`
	User   *User  `json:"user"`
}

// MEncode serializes a UserChannelMsg using msgpack
func (o *UserChannelMsg) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("action")
	encoder.WriteString(o.Action)
	encoder.WriteString("user")
	if o.User == nil {
		encoder.WriteNil()
	} else {
		o.User.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeUserChannelMsg deserializes a UserChannelMsg using msgpack
func MDecodeUserChannelMsg(d *msgpack.Decoder) (UserChannelMsg, error) {
	var val UserChannelMsg
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "action":
			val.Action, err = d.ReadString()
		case "user":
			fval, err := MDecodeUser(d)
			if err != nil {
				return val, err
			}
			val.User = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a UserChannelMsg using cbor
func (o *UserChannelMsg) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("action")
	encoder.WriteString(o.Action)
	encoder.WriteString("user")
	if o.User == nil {
		encoder.WriteNil()
	} else {
		o.User.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeUserChannelMsg deserializes a UserChannelMsg using cbor
func CDecodeUserChannelMsg(d *cbor.Decoder) (UserChannelMsg, error) {
	var val UserChannelMsg
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "action":
			val.Action, err = d.ReadString()
		case "user":
			fval, err := CDecodeUser(d)
			if err != nil {
				return val, err
			}
			val.User = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type Chat interface {
	UpdateUserList(ctx *actor.Context, arg UserChannelMsg) error
	UpdateChat(ctx *actor.Context, arg Msg) error
}

// ChatHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func ChatHandler(actor_ Chat) actor.Handler {
	return actor.NewHandler("Chat", &ChatReceiver{}, actor_)
}

// ChatContractId returns the capability contract id for this interface
func ChatContractId() string { return "jordanrash:wcchat" }

// ChatReceiver receives messages defined in the Chat service interface
type ChatReceiver struct{}

func (r *ChatReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(Chat)
	switch message.Method {

	case "UpdateUserList":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeUserChannelMsg(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.UpdateUserList(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "Chat.UpdateUserList", Arg: buf}, nil
		}
	case "UpdateChat":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeMsg(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.UpdateChat(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "Chat.UpdateChat", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "Chat."+message.Method)
	}
}

// ChatSender sends messages to a Chat service
type ChatSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a Chat provider
// implementing the 'jordanrash:wcchat' capability contract, with the "default" link
func NewProviderChat() *ChatSender {
	transport := actor.ToProvider("jordanrash:wcchat", "default")
	return &ChatSender{transport: transport}
}

// NewProviderChatLink constructs a client for sending to a Chat provider
// implementing the 'jordanrash:wcchat' capability contract, with the specified link name
func NewProviderChatLink(linkName string) *ChatSender {
	transport := actor.ToProvider("jordanrash:wcchat", linkName)
	return &ChatSender{transport: transport}
}

func (s *ChatSender) UpdateUserList(ctx *actor.Context, arg UserChannelMsg) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "Chat.UpdateUserList", Arg: buf})
	return nil
}
func (s *ChatSender) UpdateChat(ctx *actor.Context, arg Msg) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "Chat.UpdateChat", Arg: buf})
	return nil
}

type ChatSubscriber interface {
	HandleChatMessage(ctx *actor.Context, arg Msg) error
	HandleUserMessage(ctx *actor.Context, arg Msg) error
}

// ChatSubscriberHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func ChatSubscriberHandler(actor_ ChatSubscriber) actor.Handler {
	return actor.NewHandler("ChatSubscriber", &ChatSubscriberReceiver{}, actor_)
}

// ChatSubscriberContractId returns the capability contract id for this interface
func ChatSubscriberContractId() string { return "jordanrash:wcchat" }

// ChatSubscriberReceiver receives messages defined in the ChatSubscriber service interface
type ChatSubscriberReceiver struct{}

func (r *ChatSubscriberReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(ChatSubscriber)
	switch message.Method {

	case "HandleChatMessage":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeMsg(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.HandleChatMessage(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "ChatSubscriber.HandleChatMessage", Arg: buf}, nil
		}
	case "HandleUserMessage":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeMsg(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.HandleUserMessage(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "ChatSubscriber.HandleUserMessage", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "ChatSubscriber."+message.Method)
	}
}

// ChatSubscriberSender sends messages to a ChatSubscriber service
type ChatSubscriberSender struct{ transport actor.Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorChatSubscriberSender(actor_id string) *ChatSubscriberSender {
	transport := actor.ToActor(actor_id)
	return &ChatSubscriberSender{transport: transport}
}

func (s *ChatSubscriberSender) HandleChatMessage(ctx *actor.Context, arg Msg) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "ChatSubscriber.HandleChatMessage", Arg: buf})
	return nil
}
func (s *ChatSubscriberSender) HandleUserMessage(ctx *actor.Context, arg Msg) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "ChatSubscriber.HandleUserMessage", Arg: buf})
	return nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.5.1
