package keeper_test

import (
	"context"
	"testing"

	keepertest "atomic/testutil/keeper"
	"atomic/x/atomic/keeper"
	"atomic/x/atomic/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AtomicKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
