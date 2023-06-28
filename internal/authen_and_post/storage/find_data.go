package authenandpoststorage

import (
	"context"

	authenandpostmodel "github.com/lamhieo02/socialnetapp/internal/authen_and_post/model"
	"gorm.io/gorm"
)

func (s *authenAndPostStorage) FindDFindDataWithCondition(ctx context.Context, condition map[string]any) (*authenandpostmodel.UserRegister, error) {
	var user authenandpostmodel.UserRegister

	if err := s.db.Table(user.TableName()).Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}