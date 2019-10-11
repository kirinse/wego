package service

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/fsm"
	"github.com/9299381/wego/demo/src/repository"
)

type SqlService struct {
	next contracts.IService
}

func (it *SqlService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}
func (it *SqlService) Handle(ctx contracts.Context) error {

	repo := repository.NewUserRepo(ctx)
	req := make(map[string]interface{})
	req["id"] = "1189164474851006208"
	//req["user_name"] = "aaa"
	user, err := repo.First(req)
	if err != nil {
		return err
	}

	////初始化状态机
	sm := fsm.NewUserFSM(ctx, user)
	ctx.Log.Info(user.Status)
	//发送状态转换的事件
	if sm.Can("login") {
		err := sm.Event("login")
		if err != nil {
			return err
		}
		user.UserName = "aaaaaaaaa"
		ctx.Log.Info(user.Status)
		repo.Update(user)
	}
	ctx.Response("user", user)
	ctx.Response("request", ctx.Request())

	return it.next.Handle(ctx)
}
