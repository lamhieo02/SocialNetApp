package authenandpostservice

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
)

type AuthenAndPostStorage interface {
	CreateUser(context.Context, *authenandpostmodel.UserRegister)  (*uint, error)
	FindDFindDataWithCondition(ctx context.Context, condition map[string]any) (*authenandpostmodel.UserRegister, error)
	UpdateUser(ctx context.Context, condition map[string]any, user *authenandpostmodel.EditUser) error
}

type authenAndPostService struct {
	authenAndPostStorage AuthenAndPostStorage
}

func NewAuthenticateAndPostService(authenAndPostStorage AuthenAndPostStorage) *authenAndPostService {
	return &authenAndPostService{authenAndPostStorage: authenAndPostStorage}
}