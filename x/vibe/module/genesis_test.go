package vibe_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vibechain/vibe/testutil/keeper"
	"github.com/vibechain/vibe/testutil/nullify"
	vibe "github.com/vibechain/vibe/x/vibe/module"
	"github.com/vibechain/vibe/x/vibe/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VibeKeeper(t)
	vibe.InitGenesis(ctx, k, genesisState)
	got := vibe.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
