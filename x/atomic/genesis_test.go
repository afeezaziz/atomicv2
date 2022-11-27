package atomic_test

import (
	"testing"

	keepertest "atomic/testutil/keeper"
	"atomic/testutil/nullify"
	"atomic/x/atomic"
	"atomic/x/atomic/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AtomicKeeper(t)
	atomic.InitGenesis(ctx, *k, genesisState)
	got := atomic.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
