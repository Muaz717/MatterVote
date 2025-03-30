package tarantool

import (
	"MatterVote/internal/config"
	"context"
	"github.com/tarantool/go-tarantool/v2"
	"github.com/tarantool/go-tarantool/v2/pool"
)

type Storage struct {
	db tarantool.NetDialer
}

func New(ctx context.Context, cfg config.Tarantool) {
	const op = "storage.tarantool.New"

	dialer := tarantool.NetDialer{
		Address:  cfg.Host + ":" + cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
	}

	con, err := pool.Connect(ctx, dialer)
}
