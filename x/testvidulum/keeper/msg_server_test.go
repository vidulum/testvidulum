package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/vidulum/testvidulum/testutil/keeper"
	"github.com/vidulum/testvidulum/x/testvidulum/keeper"
	"github.com/vidulum/testvidulum/x/testvidulum/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestvidulumKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
