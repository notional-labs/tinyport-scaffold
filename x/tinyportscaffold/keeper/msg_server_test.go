package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/notional-labs/tinyport-scaffold/testutil/keeper"
	"github.com/notional-labs/tinyport-scaffold/x/tinyportscaffold/keeper"
	"github.com/notional-labs/tinyport-scaffold/x/tinyportscaffold/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TinyportscaffoldKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
