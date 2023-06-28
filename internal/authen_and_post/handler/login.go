package authenandposthandler

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
	authen_and_post "github.com/lamhieo02/socialnetapp/pkg/proto/authen_and_post/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (hdl *authenAndPostHandler) CheckUserAuthentication(ctx context.Context, req *authen_and_post.UserInfoLogin) (*authen_and_post.UserLoginResponse, error) {

	UserLogin := &authenandpostmodel.UserLogin{
		Email: req.GetUserEmail(),
		Password: req.GetUserPassword(),
	}
	sessionId, err := hdl.authenAndPostService.LoginUser(ctx, UserLogin)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	resp := &authen_and_post.UserLoginResponse{
		SessionId: *sessionId,
		Status: authen_and_post.UserStatus_UserStatusOK,
	}
	return resp, nil
}
