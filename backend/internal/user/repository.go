package user

import (
	"github.com/Andylixunan/insta/pkg/dbcontext"
	"github.com/Andylixunan/insta/pkg/log"
)

func NewRepository(logger *log.Logger, db *dbcontext.DB) Repository {
	return &repository{
		logger: logger,
		db:     db,
	}
}

type Repository interface {
}

type repository struct {
	logger *log.Logger
	db     *dbcontext.DB
}
