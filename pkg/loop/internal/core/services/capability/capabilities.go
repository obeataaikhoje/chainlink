package capability

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	capabilitiespb "github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/net"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

type ActionCapabilityClient struct {
	*callbackExecutableClient
	*baseCapabilityClient
}

func NewActionCapabilityClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) capabilities.ActionCapability {
	return &ActionCapabilityClient{
		callbackExecutableClient: newCallbackExecutableClient(brokerExt, conn),
		baseCapabilityClient:     newBaseCapabilityClient(brokerExt, conn),
	}
}

type ConsensusCapabilityClient struct {
	*callbackExecutableClient
	*baseCapabilityClient
}

func NewConsensusCapabilityClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) capabilities.ConsensusCapability {
	return &ConsensusCapabilityClient{
		callbackExecutableClient: newCallbackExecutableClient(brokerExt, conn),
		baseCapabilityClient:     newBaseCapabilityClient(brokerExt, conn),
	}
}

type TargetCapabilityClient struct {
	*callbackExecutableClient
	*baseCapabilityClient
}

func NewTargetCapabilityClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) capabilities.TargetCapability {
	return &TargetCapabilityClient{
		callbackExecutableClient: newCallbackExecutableClient(brokerExt, conn),
		baseCapabilityClient:     newBaseCapabilityClient(brokerExt, conn),
	}
}

type TriggerCapabilityClient struct {
	*triggerExecutableClient
	*baseCapabilityClient
}

func NewTriggerCapabilityClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) capabilities.TriggerCapability {
	return &TriggerCapabilityClient{
		triggerExecutableClient: newTriggerExecutableClient(brokerExt, conn),
		baseCapabilityClient:    newBaseCapabilityClient(brokerExt, conn),
	}
}

type CallbackCapabilityClient struct {
	*callbackExecutableClient
	*baseCapabilityClient
}

type CallbackCapability interface {
	capabilities.CallbackExecutable
	capabilities.BaseCapability
}

func NewCallbackCapabilityClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) CallbackCapability {
	return &CallbackCapabilityClient{
		callbackExecutableClient: newCallbackExecutableClient(brokerExt, conn),
		baseCapabilityClient:     newBaseCapabilityClient(brokerExt, conn),
	}
}

func RegisterCallbackCapabilityServer(server *grpc.Server, broker net.Broker, brokerCfg net.BrokerConfig, impl CallbackCapability) error {
	bext := &net.BrokerExt{
		BrokerConfig: brokerCfg,
		Broker:       broker,
	}
	capabilitiespb.RegisterCallbackExecutableServer(server, newCallbackExecutableServer(bext, impl))
	capabilitiespb.RegisterBaseCapabilityServer(server, newBaseCapabilityServer(impl))
	return nil
}

func RegisterTriggerCapabilityServer(server *grpc.Server, broker net.Broker, brokerCfg net.BrokerConfig, impl capabilities.TriggerCapability) error {
	bext := &net.BrokerExt{
		BrokerConfig: brokerCfg,
		Broker:       broker,
	}
	capabilitiespb.RegisterTriggerExecutableServer(server, newTriggerExecutableServer(bext, impl))
	capabilitiespb.RegisterBaseCapabilityServer(server, newBaseCapabilityServer(impl))
	return nil
}

type baseCapabilityServer struct {
	capabilitiespb.UnimplementedBaseCapabilityServer

	impl capabilities.BaseCapability
}

func newBaseCapabilityServer(impl capabilities.BaseCapability) *baseCapabilityServer {
	return &baseCapabilityServer{impl: impl}
}

var _ capabilitiespb.BaseCapabilityServer = (*baseCapabilityServer)(nil)

func (c *baseCapabilityServer) Info(ctx context.Context, request *emptypb.Empty) (*capabilitiespb.CapabilityInfoReply, error) {
	info, err := c.impl.Info(ctx)
	if err != nil {
		return nil, err
	}

	var ct capabilitiespb.CapabilityType
	switch info.CapabilityType {
	case capabilities.CapabilityTypeTrigger:
		ct = capabilitiespb.CapabilityType_CAPABILITY_TYPE_TRIGGER
	case capabilities.CapabilityTypeAction:
		ct = capabilitiespb.CapabilityType_CAPABILITY_TYPE_ACTION
	case capabilities.CapabilityTypeConsensus:
		ct = capabilitiespb.CapabilityType_CAPABILITY_TYPE_CONSENSUS
	case capabilities.CapabilityTypeTarget:
		ct = capabilitiespb.CapabilityType_CAPABILITY_TYPE_TARGET
	}

	return &capabilitiespb.CapabilityInfoReply{
		Id:             info.ID,
		CapabilityType: ct,
		Description:    info.Description,
		Version:        info.Version,
	}, nil
}

type baseCapabilityClient struct {
	grpc capabilitiespb.BaseCapabilityClient
	*net.BrokerExt
}

var _ capabilities.BaseCapability = (*baseCapabilityClient)(nil)

func newBaseCapabilityClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) *baseCapabilityClient {
	return &baseCapabilityClient{grpc: capabilitiespb.NewBaseCapabilityClient(conn), BrokerExt: brokerExt}
}

func (c *baseCapabilityClient) Info(ctx context.Context) (capabilities.CapabilityInfo, error) {
	resp, err := c.grpc.Info(ctx, &emptypb.Empty{})
	if err != nil {
		return capabilities.CapabilityInfo{}, err
	}

	var ct capabilities.CapabilityType
	switch resp.CapabilityType {
	case capabilitiespb.CapabilityTypeTrigger:
		ct = capabilities.CapabilityTypeTrigger
	case capabilitiespb.CapabilityTypeAction:
		ct = capabilities.CapabilityTypeAction
	case capabilitiespb.CapabilityTypeConsensus:
		ct = capabilities.CapabilityTypeConsensus
	case capabilitiespb.CapabilityTypeTarget:
		ct = capabilities.CapabilityTypeTarget
	case capabilitiespb.CapabilityTypeUnknown:
		return capabilities.CapabilityInfo{}, fmt.Errorf("invalid capability type: %s", ct)
	}

	return capabilities.CapabilityInfo{
		ID:             resp.Id,
		CapabilityType: ct,
		Description:    resp.Description,
		Version:        resp.Version,
	}, nil
}

type triggerExecutableServer struct {
	capabilitiespb.UnimplementedTriggerExecutableServer
	*net.BrokerExt

	impl capabilities.TriggerExecutable

	cancelFuncs map[string]func()
}

func newTriggerExecutableServer(brokerExt *net.BrokerExt, impl capabilities.TriggerExecutable) *triggerExecutableServer {
	return &triggerExecutableServer{
		impl:        impl,
		BrokerExt:   brokerExt,
		cancelFuncs: map[string]func(){},
	}
}

var _ capabilitiespb.TriggerExecutableServer = (*triggerExecutableServer)(nil)

func (t *triggerExecutableServer) GetRequestConfigJsonSchema(ctx context.Context, _ *emptypb.Empty) (*capabilitiespb.CapabilityResponse, error) {
	resp := t.impl.GetRequestConfigJsonSchema()

	// NOTE: Am i doing this right?
	return &capabilitiespb.CapabilityResponse{
		Value: values.Proto(resp.Value),
		Error: resp.Err.Error(),
	}, nil
}

func (t *triggerExecutableServer) RegisterTrigger(ctx context.Context, request *capabilitiespb.RegisterTriggerRequest) (*emptypb.Empty, error) {
	ch := make(chan capabilities.CapabilityResponse)

	conn, err := t.Dial(request.CallbackId)
	if err != nil {
		return nil, err
	}

	connCtx, connCancel := context.WithCancel(context.Background())
	go callbackIssuer(connCtx, capabilitiespb.NewCallbackClient(conn), ch, t.Logger)

	req := pb.CapabilityRequestFromProto(request.CapabilityRequest)
	err = t.impl.RegisterTrigger(ctx, ch, req)
	if err != nil {
		connCancel()
		return nil, err
	}

	t.cancelFuncs[request.CapabilityRequest.Metadata.WorkflowId] = connCancel
	return &emptypb.Empty{}, nil
}

func (t *triggerExecutableServer) UnregisterTrigger(ctx context.Context, request *capabilitiespb.UnregisterTriggerRequest) (*emptypb.Empty, error) {
	req := pb.CapabilityRequestFromProto(request.CapabilityRequest)
	err := t.impl.UnregisterTrigger(ctx, req)
	if err != nil {
		return nil, err
	}

	cancelFunc := t.cancelFuncs[request.CapabilityRequest.Metadata.WorkflowId]
	if cancelFunc != nil {
		cancelFunc()
	}

	return &emptypb.Empty{}, nil
}

type triggerExecutableClient struct {
	grpc capabilitiespb.TriggerExecutableClient
	*net.BrokerExt
}

var _ capabilities.TriggerExecutable = (*triggerExecutableClient)(nil)

func (t *triggerExecutableClient) GetRequestConfigJsonSchema() (*capabilities.CapabilityResponse) {
	resp, err := t.grpc.GetRequestConfigJsonSchema(context.Background(), &emptypb.Empty{})
	if err != nil {
		return &capabilities.CapabilityResponse{
			Err: err,
		}
	}

	return &capabilities.CapabilityResponse{
		Value: values.FromProto(resp.Value),
		Err: errors.New(resp.Error),
	}
}

func (t *triggerExecutableClient) RegisterTrigger(ctx context.Context, callback chan<- capabilities.CapabilityResponse, req capabilities.CapabilityRequest) error {
	cid, res, err := t.ServeNew("Callback", func(s *grpc.Server) {
		capabilitiespb.RegisterCallbackServer(s, newCallbackServer(callback))
	})
	if err != nil {
		return err
	}

	r := &capabilitiespb.RegisterTriggerRequest{
		CallbackId:        cid,
		CapabilityRequest: pb.CapabilityRequestToProto(req),
	}
	_, err = t.grpc.RegisterTrigger(ctx, r)
	if err != nil {
		t.CloseAll(res)
	}
	return err
}

func (t *triggerExecutableClient) UnregisterTrigger(ctx context.Context, req capabilities.CapabilityRequest) error {
	r := &capabilitiespb.UnregisterTriggerRequest{
		CapabilityRequest: pb.CapabilityRequestToProto(req),
	}
	_, err := t.grpc.UnregisterTrigger(ctx, r)
	return err
}

func newTriggerExecutableClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) *triggerExecutableClient {
	return &triggerExecutableClient{grpc: capabilitiespb.NewTriggerExecutableClient(conn), BrokerExt: brokerExt}
}

type callbackExecutableServer struct {
	capabilitiespb.UnimplementedCallbackExecutableServer
	*net.BrokerExt

	impl capabilities.CallbackExecutable

	cancelFuncs map[string]func()
}

func newCallbackExecutableServer(brokerExt *net.BrokerExt, impl capabilities.CallbackExecutable) *callbackExecutableServer {
	return &callbackExecutableServer{
		impl:        impl,
		BrokerExt:   brokerExt,
		cancelFuncs: map[string]func(){},
	}
}

var _ capabilitiespb.CallbackExecutableServer = (*callbackExecutableServer)(nil)

func (c *callbackExecutableServer) RegisterToWorkflow(ctx context.Context, req *capabilitiespb.RegisterToWorkflowRequest) (*emptypb.Empty, error) {
	config := values.FromProto(req.Config)

	err := c.impl.RegisterToWorkflow(ctx, capabilities.RegisterToWorkflowRequest{
		Metadata: capabilities.RegistrationMetadata{
			WorkflowID: req.Metadata.WorkflowId,
		},
		Config: config.(*values.Map),
	})
	return &emptypb.Empty{}, err
}

func (c *callbackExecutableServer) UnregisterFromWorkflow(ctx context.Context, req *capabilitiespb.UnregisterFromWorkflowRequest) (*emptypb.Empty, error) {
	config := values.FromProto(req.Config)

	err := c.impl.UnregisterFromWorkflow(ctx, capabilities.UnregisterFromWorkflowRequest{
		Metadata: capabilities.RegistrationMetadata{
			WorkflowID: req.Metadata.WorkflowId,
		},
		Config: config.(*values.Map),
	})
	return &emptypb.Empty{}, err
}

func (c *callbackExecutableServer) Execute(ctx context.Context, req *capabilitiespb.ExecuteRequest) (*emptypb.Empty, error) {
	ch := make(chan capabilities.CapabilityResponse)

	conn, err := c.Dial(req.CallbackId)
	if err != nil {
		return nil, err
	}

	connCtx, connCancel := context.WithCancel(context.Background())
	go callbackIssuer(connCtx, capabilitiespb.NewCallbackClient(conn), ch, c.Logger)

	r := pb.CapabilityRequestFromProto(req.CapabilityRequest)
	err = c.impl.Execute(ctx, ch, r)
	if err != nil {
		connCancel()
		return nil, err
	}

	c.cancelFuncs[req.CapabilityRequest.Metadata.WorkflowId] = connCancel
	return &emptypb.Empty{}, nil
}

type callbackExecutableClient struct {
	grpc capabilitiespb.CallbackExecutableClient
	*net.BrokerExt
}

func newCallbackExecutableClient(brokerExt *net.BrokerExt, conn *grpc.ClientConn) *callbackExecutableClient {
	return &callbackExecutableClient{
		grpc:      capabilitiespb.NewCallbackExecutableClient(conn),
		BrokerExt: brokerExt,
	}
}

var _ capabilities.CallbackExecutable = (*callbackExecutableClient)(nil)

func (c *callbackExecutableClient) Execute(ctx context.Context, callback chan<- capabilities.CapabilityResponse, req capabilities.CapabilityRequest) error {
	cid, res, err := c.ServeNew("Callback", func(s *grpc.Server) {
		capabilitiespb.RegisterCallbackServer(s, newCallbackServer(callback))
	})
	if err != nil {
		return err
	}

	r := &capabilitiespb.ExecuteRequest{
		CallbackId:        cid,
		CapabilityRequest: pb.CapabilityRequestToProto(req),
	}

	_, err = c.grpc.Execute(ctx, r)
	if err != nil {
		c.CloseAll(res)
	}
	return err
}

func (c *callbackExecutableClient) UnregisterFromWorkflow(ctx context.Context, req capabilities.UnregisterFromWorkflowRequest) error {
	config := &values.Map{Underlying: map[string]values.Value{}}
	if req.Config != nil {
		config = req.Config
	}

	r := &capabilitiespb.UnregisterFromWorkflowRequest{
		Config: values.Proto(config),
		Metadata: &capabilitiespb.RegistrationMetadata{
			WorkflowId: req.Metadata.WorkflowID,
		},
	}

	_, err := c.grpc.UnregisterFromWorkflow(ctx, r)
	return err
}

func (c *callbackExecutableClient) RegisterToWorkflow(ctx context.Context, req capabilities.RegisterToWorkflowRequest) error {
	config := &values.Map{Underlying: map[string]values.Value{}}
	if req.Config != nil {
		config = req.Config
	}

	r := &capabilitiespb.RegisterToWorkflowRequest{
		Config: values.Proto(config),
		Metadata: &capabilitiespb.RegistrationMetadata{
			WorkflowId: req.Metadata.WorkflowID,
		},
	}

	_, err := c.grpc.RegisterToWorkflow(ctx, r)
	return err
}

type callbackServer struct {
	capabilitiespb.UnimplementedCallbackServer
	ch chan<- capabilities.CapabilityResponse

	isClosed bool
	mu       sync.RWMutex
}

func newCallbackServer(ch chan<- capabilities.CapabilityResponse) *callbackServer {
	return &callbackServer{ch: ch}
}

func (c *callbackServer) SendResponse(ctx context.Context, resp *capabilitiespb.CapabilityResponse) (*emptypb.Empty, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.isClosed {
		return nil, errors.New("cannot send response: the underlying channel has been closed")
	}
	c.ch <- pb.CapabilityResponseFromProto(resp)
	return &emptypb.Empty{}, nil
}

func (c *callbackServer) CloseCallback(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	close(c.ch)
	c.isClosed = true
	return &emptypb.Empty{}, nil
}

func callbackIssuer(ctx context.Context, client capabilitiespb.CallbackClient, callbackChannel chan capabilities.CapabilityResponse, logger logger.Logger) {
	for {
		select {
		case <-ctx.Done():
			return
		case resp, isOpen := <-callbackChannel:
			if !isOpen {
				_, err := client.CloseCallback(ctx, &emptypb.Empty{})
				if err != nil {
					logger.Error("could not close upstream callback", err)
				}
				return
			}

			cr := pb.CapabilityResponseToProto(resp)
			_, err := client.SendResponse(ctx, cr)
			if err != nil {
				logger.Error("error sending callback response", err)
			}
		}
	}
}
