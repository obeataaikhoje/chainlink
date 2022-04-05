package telemetry

import (
	context "context"
	"crypto/ed25519"

	"github.com/smartcontractkit/chainlink-relay/core/services/telemetry/generated"
	"github.com/smartcontractkit/chainlink/core/logger"
	wsrpc "github.com/smartcontractkit/wsrpc"
)

type service struct {
	ctx              context.Context
	cancelCtx        context.CancelFunc
	serverURL        string
	clientPrivateKey ed25519.PrivateKey
	serverPublicKey  ed25519.PublicKey
	log              *logger.Logger
}

func NewService(
	serverURL string,
	clientPrivateKey ed25519.PrivateKey,
	serverPublicKey ed25519.PublicKey,
	log *logger.Logger,
) Service {
	ctx, cancelFunc := context.WithCancel(context.TODO())
	return &service{
		ctx,
		cancelFunc,
		serverURL,
		clientPrivateKey,
		serverPublicKey,
		log,
	}
}

const messageBufferCapacity = 100

func (s *service) Start() (Client, error) {
	conn, err := wsrpc.DialWithContext(
		s.ctx,
		s.serverURL,
		wsrpc.WithTransportCreds(
			s.clientPrivateKey,
			s.serverPublicKey,
		),
	)
	if err != nil {
		return &client{}, err
	}
	client := NewClient(s.ctx, generated.NewTelemetryClient(conn), messageBufferCapacity, s.log)
	return client, nil
}

func (s *service) Stop() {
	s.cancelCtx()
}
