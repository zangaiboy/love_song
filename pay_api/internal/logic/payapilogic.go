package logic

import (
	"context"

	"love_song/pay_api/internal/svc"
	"love_song/pay_api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Pay_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPay_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) Pay_apiLogic {
	return Pay_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Pay_apiLogic) Pay_api(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
