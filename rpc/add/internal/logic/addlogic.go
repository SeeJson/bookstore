package logic

import (
	"context"

	"github.com/SeeJson/bookstore/rpc/add/internal/svc"
	pb "github.com/SeeJson/bookstore/rpc/add/pb"
	"github.com/SeeJson/bookstore/rpc/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.AddReq) (*pb.AddResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.Modele.Insert(l.ctx, &model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddResp{
		Ok: true,
	}, nil
}
