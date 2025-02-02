package masterTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/userTransaction"
)

type userTransactionDao struct {
	ShardConn *database.ShardConn
}

func NewUserTransactionDao(conn *database.SqlHandler) userTransaction.UserTransactionRepository {
	return &userTransactionDao{
		ShardConn: conn.User,
	}
}

func (d *userTransactionDao) Begin(ctx context.Context, shardKey string) (*gorm.DB, error) {
	tx := d.ShardConn.Shards[shardKey].WriteConn.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *userTransactionDao) Commit(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *userTransactionDao) Rollback(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
