package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gliderlabs/ssh"
	"github.com/google/uuid"
	chat "github.com/jordan-rash/wasmcloud-chat/interface"
	provider "github.com/jordan-rash/wasmcloud-provider"
	wccore "github.com/jordan-rash/wasmcloud-provider/core"
	"github.com/sirupsen/logrus"
	core "github.com/wasmcloud/interfaces/core/tinygo"
	msgpack "github.com/wasmcloud/tinygo-msgpack"

	"github.com/charmbracelet/keygen"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

const useHighPerformanceRenderer = false

func init() {
	os.Setenv("RUST_LOG", "debug")

	file, err := os.OpenFile("/tmp/chat_wc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)
}

var (
	log               = logrus.New()
	wasmcloudProvider = provider.WasmcloudProvider{}
	m                 = model{}
	localUser         = chat.User{}
	aID               string
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	hostDataRaw, _ := reader.ReadString('\n')

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error

	wasmcloudProvider, err = provider.Init(ctx, hostDataRaw)
	if err != nil {
		log.Error(err)
		cancel()
	}

	// Listen for Shutdown request
	s := &ssh.Server{}
	go func() {
		<-wasmcloudProvider.Shutdown

		msg := chat.Msg{
			Owner: &localUser,
			Value: "leave",
			Id:    uuid.NewString(),
		}
		sendDownLattice(msg, SEND_USER_UPDATE)

		s.Shutdown(ctx)
		close(wasmcloudProvider.ProviderAction)
		cancel()
	}()

	// Wait for a valid link definiation
	log.Println("Ready for link definitions")
	actorData := <-wasmcloudProvider.Links
	aID = actorData.ActorID

	k, err := keygen.New("", nil, keygen.Ed25519)
	if err != nil {
		return
	}
	log.Printf("Private Key: %v", string(k.PrivateKeyPEM()))
	log.Printf("Public Key: %s", string(k.PublicKey()))

	// Configure and start ssh server
	config := validateConfig(actorData.ActorConfig)

	s, err = wish.NewServer(
		wish.WithAddress("127.0.0.1:"+config.Port),
		wish.WithHostKeyPEM(k.PrivateKeyPEM()),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			logging.Middleware(),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	// Start SSH Server
	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Start listening on topic for requests from actor
	wasmcloudProvider.ListenForActor(actorData.ActorID)

	//Wait for valid requests
	log.Print("Waiting for actor requests")
	for actorRequest := range wasmcloudProvider.ProviderAction {
		go evaluateRequest(actorRequest, config)
	}
}

func evaluateRequest(actorRequest provider.ProviderAction, actorConfig ChatClientConfig) error {
	log.Printf("[MAIN] GOT COMMAND " + actorRequest.Operation)

	resp := provider.ProviderResponse{}

	switch actorRequest.Operation {
	case UPDATE_USER_LIST.String():
		d := msgpack.NewDecoder(actorRequest.Msg)
		msg, err := chat.MDecodeUserChannelMsg(&d)
		if err != nil {
			resp.Error = err.Error()
		}
		cm := chat.Msg{
			Owner: msg.User,
			Value: "NOOP",
			Id:    uuid.NewString(),
		}
		switch msg.Action {
		case "join":
			cm.Value = "USERJOIN"
		case "leave":
			cm.Value = "USERLEAVE"
		default:
			return errors.New("Unknown user action")
		}
		if msg.User.Name != localUser.Name {
			m.Update(&cm)
		}

	case UPDATE_CHAT.String():
		d := msgpack.NewDecoder(actorRequest.Msg)
		msg, err := chat.MDecodeMsg(&d)
		if err != nil {
			resp.Error = err.Error()
		}
		if msg.Owner.Name != localUser.Name {
			m.Update(msg)
		}

	default:
		resp.Error = "Invalid SSHChat Operation"
	}

	actorRequest.Respond <- resp
	return nil
}

func sendDownLattice(msg chat.Msg, op Operation) {
	if *msg.Owner == localUser {
		buf := encodeMsg(msg)
		guid := GenGuid()

		var capId core.CapabilityContractId = "jordanrash:wcchat"

		i := core.Invocation{
			Origin: core.WasmCloudEntity{
				PublicKey:  wasmcloudProvider.HostData.ProviderKey,
				LinkName:   wasmcloudProvider.HostData.LinkName,
				ContractId: capId,
			},
			Target: core.WasmCloudEntity{
				PublicKey:  aID,
				LinkName:   wasmcloudProvider.HostData.LinkName,
				ContractId: capId,
			},
			Operation:     op.String(),
			Msg:           buf,
			Id:            guid,
			HostId:        wasmcloudProvider.HostData.HostId,
			ContentLength: uint64(len(buf)),
		}

		provider.EncodeClaims(&i, wasmcloudProvider.HostData, guid)

		natsBody := wccore.EncodeInvocation(i)

		// NC Request
		subj := fmt.Sprintf("wasmbus.rpc.%s.%s", wasmcloudProvider.HostData.LatticeRpcPrefix, aID)
		log.
			WithField("subj", subj).
			WithField("op", op).
			WithField("msg", msg).
			Debug("msg sent")
		wasmcloudProvider.NatsConnection.Request(subj, natsBody, 2*time.Second)
	}
}
