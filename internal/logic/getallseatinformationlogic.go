package logic

import (
	"context"
	"muxi-live-stream-api/internal/tool"

	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllSeatInformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllSeatInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllSeatInformationLogic {
	return &GetAllSeatInformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllSeatInformationLogic) GetAllSeatInformation(req *types.GetAllSeatInformationRequest, cookie string) (resp *types.Response, err error) {
	if cookie == "" {
		return &types.Response{
			Code:    401,
			Message: "cookie 不能为空",
		}, nil
	}

	if req.RoomIds == nil {
		return &types.Response{
			Code:    400,
			Message: "room_ids 不能为空",
		}, nil
	}

	grabber := tool.NewGrabber(req.RoomIds, "08:00", "22:00")
	allInfo := grabber.SearchAllSeats(cookie)

	return &types.Response{
		Code:    200,
		Data:    allInfo,
		Message: "获取成功",
	}, nil
}
