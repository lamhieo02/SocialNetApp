package authenandpostservice

import (
	"context"
	"fmt"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
	utilshasher "github.com/lamhieo02/socialnetapp/pkg/utils/hasher"
	"github.com/redis/go-redis/v9"
)

func (service *authenAndPostService) RegisterUser(ctx context.Context, req *authenandpostmodel.UserRegister) (*uint, error) {
	// TODO: check is duplicate email => solution is use redis to store username and email
	condition := map[string]interface{}{
		"email": req.Email,
	}

	// check is duplicate email in redis
	emailKey := fmt.Sprintf("useremail:%s", req.Email)
	val, err := service.authenUserCache.Get(ctx, emailKey)
	if err != nil {
		if err.Error() != redis.Nil.Error() {
			return nil, err
		}
	}

	if val != nil {
		return nil, fmt.Errorf("email %s is already exist", req.Email)
	}

	// check is duplicate email in db
	newUser, err := service.authenAndPostStorage.FindDFindDataWithCondition(ctx, condition)
	if err != nil {
		return nil, err
	}

	if newUser != nil {
		return nil, fmt.Errorf("email %s is already exist", req.Email)
	}

	// hash password before save to db
	req.Password, req.Salt, _ = utilshasher.HashPassword(req.Password)

	id, err := service.authenAndPostStorage.CreateUser(ctx, req)

	if err != nil {
		return nil, err
	}

	keyEmail := fmt.Sprintf("useremail:%s", req.Email)
	if err := service.authenUserCache.Set(ctx, keyEmail, true); err != nil {
		return nil, err
	}
	return id, nil
}
