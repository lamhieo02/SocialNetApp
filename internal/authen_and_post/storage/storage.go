package authenandpoststorage

import (
	"github.com/lamhieo02/socialnetapp/config"
	"github.com/lamhieo02/socialnetapp/internal/entities"
	mysqlpkg "github.com/lamhieo02/socialnetapp/pkg/proto/mysql"
	"gorm.io/gorm"
)

type authenAndPostStorage struct {
	db *gorm.DB
}

type AuthenAndPostStorageInterface interface {
	CreateUser(*entities.UserRegister)  error
}

func NewAuthenticateAndPostStorage(conf *config.Config) (*authenAndPostStorage, error) {
	db, err := mysqlpkg.NewMYSQLConnection(&conf.Mysql)
	if err != nil {
		return nil, err
	}
	return &authenAndPostStorage{
		db: db,
	}, nil
}