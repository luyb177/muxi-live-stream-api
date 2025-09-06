package logic

import (
	"context"

	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomIdsLogic {
	return &GetRoomIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomIdsLogic) GetRoomIds(cookie string) (resp *types.Response, err error) {
	if cookie == "" {
		return &types.Response{
			Code:    400,
			Message: "cookie 不能为空",
		}, nil
	}

	var roomIds = []string{
		"100455820", // 一楼综合学习室
		"100455822", // 二楼借阅室(一)
		"100671994", // 二楼借阅室(二)
		"100455824", // 三楼借阅室(三)
		"100455828", // 五楼借阅室(四)
		"100746476", // 五楼借阅室(五)
		"100746204", // 六楼阅览室(一)
		"100455830", // 六楼外文借阅室
		"100455832", // 七楼阅览室(二)
		"100746480", // 七楼阅览室(三)
		"100455834", // 九楼阅览室
		"101699179", // 南湖分馆一楼开敞座位区
		"101699187", // 南湖分馆一楼中庭开敞座位区
		"101699189", // 南湖分馆二楼开敞座位区
		"101699191", // 南湖分馆二楼卡座区
	}
	return &types.Response{
		Code:    200,
		Message: "获取成功",
		Data:    types.GetRoomIdsResponse{RoomIds: roomIds},
	}, nil
}
