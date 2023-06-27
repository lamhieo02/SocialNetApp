package authenandpoststorage

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
)

func (storage *authenAndPostStorage) CreateUser(ctx context.Context, user * authenandpostmodel.UserRegister) (*uint, error) {
	db := storage.db.Begin()

	if err := db.Table(user.TableName()).Create(user).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, err
	}
	return &user.ID, nil
}