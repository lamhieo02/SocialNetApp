package authenandpostservice

import (
	"context"
	"errors"

	"github.com/google/uuid"
	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
	utilshasher "github.com/lamhieo02/socialnetapp/pkg/utils/hasher"
)

func (service *authenAndPostService) LoginUser(ctx context.Context, req *authenandpostmodel.UserLogin) (*string, error) {
	// B1: find user by email
	condition := map[string]interface{}{
		"email": req.Email,
	} 

	user, err := service.authenAndPostStorage.FindDFindDataWithCondition(ctx, condition)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// B2: compare password
	if err := utilshasher.ComparePassword(user.Password, req.Password, user.Salt); err != nil {
		return nil, err
	}

	// B3: generate session id and return; TODO: use redis to store session id
	// replace by token will be better
	sessionId := generateSessionId()
    return &sessionId, nil
}

func generateSessionId() string {
    sessionID := uuid.New().String()
    return sessionID
}