package keeper_test

import (
	"testing"

	testkeeper "github.com/notional-labs/tinyport-scaffold/testutil/keeper"
	"github.com/notional-labs/tinyport-scaffold/x/tinyportscaffold/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TinyportscaffoldKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
