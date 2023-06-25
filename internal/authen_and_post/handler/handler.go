package authenandposthandler

import (
	"context"

	"github.com/lamhieo02/socialnetapp/internal/entities"
	authen_and_post "github.com/lamhieo02/socialnetapp/pkg/proto/authen_and_post/pkg/proto"
)

type AuthenAndPostService interface {
	RegisterUser(ctx context.Context, req *entities.UserRegister)  error
}

type authenAndPostHandler struct {
	authen_and_post.UnimplementedAuthenticateAndPostServer
	authenAndPostService AuthenAndPostService
}

func (hdl *authenAndPostHandler) CreateUser(ctx context.Context, req *authen_and_post.UserDetailInfo) (*authen_and_post.CreateUserResponse, error) {
	return nil, nil
}


func NewAuthenticateAndPostHandler(auAndPostSrv AuthenAndPostService) (*authenAndPostHandler, error) {
	return &authenAndPostHandler{authenAndPostService: auAndPostSrv}, nil
}