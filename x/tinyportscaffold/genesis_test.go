package tinyportscaffold_test

import (
	"testing"

	keepertest "github.com/notional-labs/tinyport-scaffold/testutil/keeper"
	"github.com/notional-labs/tinyport-scaffold/testutil/nullify"
	"github.com/notional-labs/tinyport-scaffold/x/tinyportscaffold"
	"github.com/notional-labs/tinyport-scaffold/x/tinyportscaffold/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by tinyport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TinyportscaffoldKeeper(t)
	tinyportscaffold.InitGenesis(ctx, *k, genesisState)
	got := tinyportscaffold.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by tinyport scaffolding # genesis/test/assert
}
