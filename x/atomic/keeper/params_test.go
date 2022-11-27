package keeper_test

import (
	"testing"

	testkeeper "atomic/testutil/keeper"
	"atomic/x/atomic/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AtomicKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
