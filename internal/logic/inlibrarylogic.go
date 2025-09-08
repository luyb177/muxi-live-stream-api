package logic

import (
	"context"
	"fmt"

	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/tool"
	"muxi-live-stream-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InLibraryLogic {
	return &InLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InLibraryLogic) InLibrary(req *types.InLibraryRequest, cookie string) (resp *types.Response, err error) {
	if cookie == "" {
		return &types.Response{
			Code:    401,
			Message: "cookie 不能为空",
		}, nil
	}

	if req.Name == "" {
		return &types.Response{
			Code:    400,
			Message: "姓名不能为空",
		}, nil
	}
	if req.RoomIds == nil {
		return &types.Response{
			Code:    400,
			Message: "room_ids 不能为空",
		}, nil
	}

	// 初始化一个 grabber
	grabber := tool.NewGrabber(req.RoomIds, "08:00", "22:00")

	ot := grabber.IsInLibrary(req.Name, cookie)
	if ot != nil {
		fmt.Printf("%s 在图书馆的%s，%s - %s\n", req.Name, ot.Title, ot.Start, ot.End)
		return &types.Response{
			Code:    200,
			Message: fmt.Sprintf("%s 在图书馆的%s，%s - %s", req.Name, ot.Title, ot.Start, ot.End),
			Data: types.InLibraryResponse{
				IsInLibrary: true,
				Area:        ot.Title,
				Start:       ot.Start,
				End:         ot.End,
			},
		}, nil

	} else {
		fmt.Println("不在图书馆")
		return &types.Response{
			Code:    200,
			Message: "不在图书馆",
			Data: types.InLibraryResponse{
				IsInLibrary: false,
			},
		}, nil
	}

}
