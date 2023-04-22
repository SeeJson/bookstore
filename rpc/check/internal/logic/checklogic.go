package logic

import (
	"context"
	pb "github.com/SeeJson/bookstore/rpc/check/pb"

	"github.com/SeeJson/bookstore/rpc/check/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *pb.CheckReq) (*pb.CheckResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Model.FindOne(l.ctx, in.Book)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
}
