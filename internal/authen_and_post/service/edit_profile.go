package authenandpostservice

import (
	"context"
	"errors"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
)

func (service *authenAndPostService) EditProfile(ctx context.Context, userId *uint, userEdit *authenandpostmodel.EditUser) (err error) {
	// condition
	condition := map[string]any{
		"id":userId,
	}

	// check exist by userId
	data, _ := service.authenAndPostStorage.FindDFindDataWithCondition(ctx, condition)
	if data == nil {
		return errors.New("user not found")
	}

	// update user
	err = service.authenAndPostStorage.UpdateUser(ctx, condition, userEdit)
	if err != nil {
		return err
	}
	return nil
}