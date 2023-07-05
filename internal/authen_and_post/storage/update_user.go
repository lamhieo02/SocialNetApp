package authenandpoststorage

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
)

func (s *authenAndPostStorage) UpdateUser(ctx context.Context, condition map[string]any, user *authenandpostmodel.EditUser) error {
	
	if err := s.db.Table(user.TableName()).Where(condition).Updates(user).Error; err != nil {
		return err
	}

	return nil
}