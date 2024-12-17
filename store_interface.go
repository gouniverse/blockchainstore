package blockchainstore

import "context"

type StoreInterface interface {
	BlockCreate(ctx context.Context, block *Block) error
	BlockDelete(ctx context.Context, block *Block) error
	BlockDeleteByID(ctx context.Context, blockID string) error
	BlockFindByID(ctx context.Context, id string) (*Block, error)
	BlockList(ctx context.Context, options BlockQueryOptions) ([]Block, error)
	BlockUpdate(ctx context.Context, block *Block) error
}
