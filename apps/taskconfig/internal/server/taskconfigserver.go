// Code generated by goctl. DO NOT EDIT.
// Source: taskconfig.proto

package server

import (
	"context"

	"oh-my-helper-go/apps/taskconfig/internal/logic"
	"oh-my-helper-go/apps/taskconfig/internal/svc"
	"oh-my-helper-go/apps/taskconfig/taskconfig"
)

type TaskConfigServer struct {
	svcCtx *svc.ServiceContext
	taskconfig.UnimplementedTaskConfigServer
}

func NewTaskConfigServer(svcCtx *svc.ServiceContext) *TaskConfigServer {
	return &TaskConfigServer{
		svcCtx: svcCtx,
	}
}

func (s *TaskConfigServer) Ping(ctx context.Context, in *taskconfig.Request) (*taskconfig.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *TaskConfigServer) GetBiliConfig(ctx context.Context, in *taskconfig.BiliConfigRequest) (*taskconfig.BiliConfigResponse, error) {
	l := logic.NewGetBiliConfigLogic(ctx, s.svcCtx)
	return l.GetBiliConfig(in)
}

func (s *TaskConfigServer) EditBiliConfig(ctx context.Context, in *taskconfig.BiliConfigEditRequest) (*taskconfig.BiliConfigResponse, error) {
	l := logic.NewEditBiliConfigLogic(ctx, s.svcCtx)
	return l.EditBiliConfig(in)
}

func (s *TaskConfigServer) AddBiliConfig(ctx context.Context, in *taskconfig.BiliConfigAddRequest) (*taskconfig.BiliConfigResponse, error) {
	l := logic.NewAddBiliConfigLogic(ctx, s.svcCtx)
	return l.AddBiliConfig(in)
}
