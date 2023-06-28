package authenandpostmodel

type UserRegister struct {
	ID          uint   `gorm:"primary"`
	Password    string `gorm:"size:50;not null"`
	FirstName   string `gorm:"size:50;not null"`
	LastName    string `gorm:"size:50;not null"`
	DateOfBirth string `gorm:"not null"`
	Email       string `gorm:"size:50;not null"`
	UserName    string `gorm:"size:50;not null;index:idx_username"`
}

func (UserRegister) TableName() string {
	return "user"
}