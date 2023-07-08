package authenandposthandler

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
	authen_and_post "github.com/lamhieo02/socialnetapp/pkg/proto/authen_and_post/pkg/proto"
)

type AuthenAndPostService interface {
	RegisterUser(ctx context.Context, req *authenandpostmodel.UserRegister) (*uint, error)
	LoginUser(ctx context.Context, req *authenandpostmodel.UserLogin) (*string, error)
	EditProfile(ctx context.Context, userId *uint, userEdit *authenandpostmodel.EditUser) (err error)

}

type authenAndPostHandler struct {
	authen_and_post.UnimplementedAuthenticateAndPostServer
	authenAndPostService AuthenAndPostService
}

func NewAuthenticateAndPostHandler(auAndPostSrv AuthenAndPostService) *authenAndPostHandler {
	return &authenAndPostHandler{authenAndPostService: auAndPostSrv}
}
