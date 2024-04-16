package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	loopnet "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/net"
	ccippb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/ccip"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/relayer/pluginprovider/ext/ccip"
	looptest "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/test"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestStaticPriceRegistry(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)
	// static test implementation is self consistent
	assert.NoError(t, PriceRegistryReader.Evaluate(ctx, PriceRegistryReader))

	// error when the test implementation is evaluates something that differs from the static implementation
	botched := PriceRegistryReader
	botched.addressResponse = "not what we expect"
	err := PriceRegistryReader.Evaluate(ctx, botched)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not what we expect")
}

func TestPriceRegistryGRPC(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	scaffold := looptest.NewGRPCScaffold(t, setupPriceRegistryServer, setupPriceRegistryClient)
	roundTripPriceRegistryTests(ctx, t, scaffold.Client())
	// price registry implements dependency management, test that it closes properly
	t.Run("Dependency management", func(t *testing.T) {
		d := &looptest.MockDep{}
		scaffold.Server().AddDep(d)
		assert.False(t, d.IsClosed())
		scaffold.Client().Close()
		assert.True(t, d.IsClosed())
	})
}

// roundTripPriceRegistryTests tests the round trip of the client<->server.
// it should exercise all the methods of the client.
// do not add client.Close to this test, test that from the driver test
func roundTripPriceRegistryTests(ctx context.Context, t *testing.T, client cciptypes.PriceRegistryReader) {
	t.Run("Address", func(t *testing.T) {
		address, err := client.Address(ctx)
		require.NoError(t, err)
		assert.Equal(t, PriceRegistryReader.addressResponse, address)
	})

	t.Run("GetFeeTokens", func(t *testing.T) {
		price, err := client.GetFeeTokens(ctx)
		require.NoError(t, err)
		assert.Equal(t, PriceRegistryReader.getFeeTokensResponse, price)
	})

	t.Run("GetGasPriceUpdatesCreatedAfter", func(t *testing.T) {
		price, err := client.GetGasPriceUpdatesCreatedAfter(ctx,
			PriceRegistryReader.getGasPriceUpdatesCreatedAfterRequest.chainSelector,
			PriceRegistryReader.getGasPriceUpdatesCreatedAfterRequest.ts,
			PriceRegistryReader.getGasPriceUpdatesCreatedAfterRequest.confirmations,
		)
		require.NoError(t, err)
		assert.Equal(t, PriceRegistryReader.getGasPriceUpdatesCreatedAfterResponse, price)
	})

	t.Run("GetTokenPriceUpdatesCreatedAfter", func(t *testing.T) {
		price, err := client.GetTokenPriceUpdatesCreatedAfter(ctx,
			PriceRegistryReader.getTokenPriceUpdatesCreatedAfterRequest.ts,
			PriceRegistryReader.getTokenPriceUpdatesCreatedAfterRequest.confirmations,
		)
		require.NoError(t, err)
		assert.Equal(t, PriceRegistryReader.getTokenPriceUpdatesCreatedAfterResponse, price)
	})

	t.Run("GetTokenPrices", func(t *testing.T) {
		price, err := client.GetTokenPrices(ctx, PriceRegistryReader.getTokenPricesRequest)
		require.NoError(t, err)
		assert.Equal(t, PriceRegistryReader.getTokenPricesResponse, price)
	})

	t.Run("GetTokensDecimals", func(t *testing.T) {
		price, err := client.GetTokensDecimals(ctx, PriceRegistryReader.getTokensDecimalsRequest)
		require.NoError(t, err)
		assert.Equal(t, PriceRegistryReader.getTokensDecimalsResponse, price)
	})
}

func setupPriceRegistryServer(t *testing.T, server *grpc.Server, b *loopnet.BrokerExt) *ccip.PriceRegistryGRPCServer {
	priceRegistry := ccip.NewPriceRegistryGRPCServer(PriceRegistryReader)
	ccippb.RegisterPriceRegistryReaderServer(server, priceRegistry)
	return priceRegistry
}

// wrapper to enable use of the grpc scaffold
func setupPriceRegistryClient(b *loopnet.BrokerExt, conn grpc.ClientConnInterface) *ccip.PriceRegistryGRPCClient {
	return ccip.NewPriceRegistryGRPCClient(conn)
}
