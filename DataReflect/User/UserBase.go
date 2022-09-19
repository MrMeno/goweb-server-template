package User

/*
tb_user_info
*/
type UserInfo struct {
	UserID   string `gorm:"userid"`
	Name     string `gorm:"name"`
	Password string `gorm:"password"`
	CNName   string `gorm:"cnname"`
}

func (tb_user_info *UserInfo) TableName() string {
	return "tb_user_info"
}
