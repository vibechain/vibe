package keeper

import (
	"github.com/vibechain/vibe/x/vibe/types"
)

var _ types.QueryServer = Keeper{}
