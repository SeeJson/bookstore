package svc

import (
	"github.com/SeeJson/bookstore/rpc/add/internal/config"
	"github.com/SeeJson/bookstore/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Modele model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Modele: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
