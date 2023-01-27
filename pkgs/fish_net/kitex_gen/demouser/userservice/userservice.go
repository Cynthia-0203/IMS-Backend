// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	demouser "fish_net/kitex_gen/demouser"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*demouser.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"MGetUser":   kitex.NewMethodInfo(mGetUserHandler, newUserServiceMGetUserArgs, newUserServiceMGetUserResult, false),
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newUserServiceCheckUserArgs, newUserServiceCheckUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "demouser",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demouser.UserServiceCreateUserArgs)
	realResult := result.(*demouser.UserServiceCreateUserResult)
	success, err := handler.(demouser.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return demouser.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return demouser.NewUserServiceCreateUserResult()
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demouser.UserServiceMGetUserArgs)
	realResult := result.(*demouser.UserServiceMGetUserResult)
	success, err := handler.(demouser.UserService).MGetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceMGetUserArgs() interface{} {
	return demouser.NewUserServiceMGetUserArgs()
}

func newUserServiceMGetUserResult() interface{} {
	return demouser.NewUserServiceMGetUserResult()
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demouser.UserServiceCheckUserArgs)
	realResult := result.(*demouser.UserServiceCheckUserResult)
	success, err := handler.(demouser.UserService).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCheckUserArgs() interface{} {
	return demouser.NewUserServiceCheckUserArgs()
}

func newUserServiceCheckUserResult() interface{} {
	return demouser.NewUserServiceCheckUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, req *demouser.CreateUserRequest) (r *demouser.CreateUserResponse, err error) {
	var _args demouser.UserServiceCreateUserArgs
	_args.Req = req
	var _result demouser.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, req *demouser.MGetUserRequest) (r *demouser.MGetUserResponse, err error) {
	var _args demouser.UserServiceMGetUserArgs
	_args.Req = req
	var _result demouser.UserServiceMGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, req *demouser.CheckUserRequest) (r *demouser.CheckUserResponse, err error) {
	var _args demouser.UserServiceCheckUserArgs
	_args.Req = req
	var _result demouser.UserServiceCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
