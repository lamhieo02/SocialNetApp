package authenandposthandler

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
	authen_and_post "github.com/lamhieo02/socialnetapp/pkg/proto/authen_and_post/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (hdl *authenAndPostHandler) EditUser(ctx context.Context, req *authen_and_post.UserEdit) (*authen_and_post.EditUserResponse, error) {
	userEdit := &authenandpostmodel.EditUser{
		FirstName: req.FirstName,
		LastName: req.LastName,
		DateOfBirth: req.Dob,
	}
	userId := uint(req.Id)

	editUserResp := authen_and_post.EditUserResponse{}

	if err := hdl.authenAndPostService.EditProfile(ctx, &userId, userEdit); err != nil {
		editUserResp.Status = authen_and_post.UserStatus_NOT_FOUND
		return &editUserResp, status.Errorf(codes.InvalidArgument, err.Error())
	}

	editUserResp.Status = authen_and_post.UserStatus_UserStatusOK
	return &editUserResp, nil
	
} 