package keeper

import (
	"encoding/binary"

	"github.com/cosmonaut/blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/**
* Get the store based on a store key and specific key. In this instance it is "blog" for the store key.
**/
func getStore(ctx sdk.Context, storeKey sdk.StoreKey, key string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(storeKey), []byte(key))
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := getStore(ctx, k.storeKey, types.PostCountKey)

	byteKey := []byte(types.PostCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := getStore(ctx, k.storeKey, types.PostCountKey)

	byteKey := []byte(types.PostCountKey)
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)

	post.Id = count

	store := getStore(ctx, k.storeKey, types.PostKey)

	byteKey := make([]byte, 8)

	binary.BigEndian.PutUint64(byteKey, post.Id)

	appendedValue := k.cdc.MustMarshal(&post)

	store.Set(byteKey, appendedValue)

	k.SetPostCount(ctx, count+1)
	return count
}
