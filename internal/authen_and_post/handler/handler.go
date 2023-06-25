package authenandposthandler

import (
	"context"

	"github.com/lamhieo02/socialnetapp/internal/entities"
	authen_and_post "github.com/lamhieo02/socialnetapp/pkg/proto/authen_and_post/pkg/proto"
)

type AuthenAndPostService interface {
	RegisterUser(ctx context.Context, req *entities.UserRegister) (*uint, error)
}

type authenAndPostHandler struct {
	authen_and_post.UnimplementedAuthenticateAndPostServer
	authenAndPostService AuthenAndPostService
}

func (hdl *authenAndPostHandler) CreateUser(ctx context.Context, req *authen_and_post.UserDetailInfo) (*authen_and_post.CreateUserResponse, error) {
	userInfo := &entities.UserRegister{
		UserName: req.UserName,
		DateOfBirth: req.Dob,
		Password: req.Password,
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
	}

	id, err := hdl.authenAndPostService.RegisterUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	resp := &authen_and_post.CreateUserResponse{
		UserId: int64(*id),
		Status: *authen_and_post.CreateUserStatus_CreateUserStatusOK.Enum(),
	}
	return resp, nil
}


func NewAuthenticateAndPostHandler(auAndPostSrv AuthenAndPostService) *authenAndPostHandler {
	return &authenAndPostHandler{authenAndPostService: auAndPostSrv}
}