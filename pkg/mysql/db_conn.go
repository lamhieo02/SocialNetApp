package mysqlpkg

import (
	"fmt"

	"github.com/lamhieo02/socialnetapp/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMYSQLConnection(cfg *config.MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password,
	cfg.Host, cfg.Port, cfg.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}