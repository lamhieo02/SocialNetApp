package authenandpostservice

import (
	"context"

	"github.com/lamhieo02/socialnetapp/internal/entities"
)

type AuthenAndPostStorage interface {
	CreateUser(context.Context, *entities.UserRegister)  error
}

type authenAndPostService struct {
	authenAndPostStorage AuthenAndPostStorage
}

func (service *authenAndPostService) RegisterUser(ctx context.Context, req *entities.UserRegister)  error {
	if err := service.authenAndPostStorage.CreateUser(ctx, req); err != nil {	
		return err
	}
	return nil
}


func NewAuthenticateAndPostService(authenAndPostStorage AuthenAndPostStorage) (*authenAndPostService, error) {
	return &authenAndPostService{authenAndPostStorage: authenAndPostStorage}, nil
}