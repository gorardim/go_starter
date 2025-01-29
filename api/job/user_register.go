package job

import "context"

//x:nsq travel.user_register 用户注册
type UserRegisterConsumer interface {
	//x:channel default 注册
	Default(ctx context.Context, request *UserRegisterRequest) error
}

type UserRegisterRequest struct {
	// 用户id
	UserId int `json:"user_id"`
}
