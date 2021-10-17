package keeper

import (
	"github.com/vidulum/testvidulum/x/testvidulum/types"
)

var _ types.QueryServer = Keeper{}
