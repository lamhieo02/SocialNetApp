package authenandpostservice

import (
	"context"

	"github.com/lamhieo02/socialnetapp/internal/entities"
)

type AuthenAndPostStorage interface {
	CreateUser(context.Context, *entities.UserRegister)  (*uint, error)
}

type authenAndPostService struct {
	authenAndPostStorage AuthenAndPostStorage
}

func (service *authenAndPostService) RegisterUser(ctx context.Context, req *entities.UserRegister) (*uint, error) {
	id, err := service.authenAndPostStorage.CreateUser(ctx, req)
	if err != nil {	
		return nil, err
	}
	return id, nil
}


func NewAuthenticateAndPostService(authenAndPostStorage AuthenAndPostStorage) *authenAndPostService {
	return &authenAndPostService{authenAndPostStorage: authenAndPostStorage}
}