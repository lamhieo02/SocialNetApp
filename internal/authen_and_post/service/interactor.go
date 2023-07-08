package authenandpostservice

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
)

type AuthenAndPostStorage interface {
	CreateUser(context.Context, *authenandpostmodel.UserRegister) (*uint, error)
	FindDFindDataWithCondition(ctx context.Context, condition map[string]any) (*authenandpostmodel.UserRegister, error)
	UpdateUser(ctx context.Context, condition map[string]any, user *authenandpostmodel.EditUser) error
}

type AuthenUserCache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (*string, error)
}

type authenAndPostService struct {
	authenAndPostStorage AuthenAndPostStorage
	authenUserCache      AuthenUserCache
}

func NewAuthenticateAndPostService(authenAndPostStorage AuthenAndPostStorage, authenUserCache AuthenUserCache) *authenAndPostService {
	return &authenAndPostService{authenAndPostStorage: authenAndPostStorage, authenUserCache: authenUserCache}
}
