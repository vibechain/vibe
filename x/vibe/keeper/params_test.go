package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/vibechain/vibe/testutil/keeper"
	"github.com/vibechain/vibe/x/vibe/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VibeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
