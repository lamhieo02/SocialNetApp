package authenandposthandler

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
	authen_and_post "github.com/lamhieo02/socialnetapp/pkg/proto/authen_and_post/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (hdl *authenAndPostHandler) CreateUser(ctx context.Context, req *authen_and_post.UserDetailInfo) (*authen_and_post.CreateUserResponse, error) {
	userInfo := &authenandpostmodel.UserRegister{
		UserName: req.UserName,
		DateOfBirth: req.Dob,
		Password: req.Password,
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
	}

	id, err := hdl.authenAndPostService.RegisterUser(ctx, userInfo)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	resp := &authen_and_post.CreateUserResponse{
		UserId: int64(*id),
		Status: *authen_and_post.CreateUserStatus_CreateUserStatusOK.Enum(),
	}
	return resp, nil
}
