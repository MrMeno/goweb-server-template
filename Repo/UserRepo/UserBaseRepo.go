package UserRepo

import (
	"MetaWebServer/DataReflect/User"
	"MetaWebServer/Utils"
)

type UserBaseRepo interface {
	GetAllUser() (*[]User.UserInfo, error)
	UpdateUserInfo(map[string]interface{}) (int64, error)
}

type UserBase struct {
	log *Utils.Log4g
}

func (s *UserBase) GetAllUser() (*[]User.UserInfo, error) {
	userModel := &User.UserInfo{}
	reslist := []User.UserInfo{}
	db, err := Utils.GetConnByKey("")
	if err != nil {
		s.log.Error("db connection error", err.Error())
		return nil, err
	}
	db.Find(userModel).Scan(&reslist)
	return &reslist, err
}

func (s *UserBase) UpdateUserInfo(conditon map[string]interface{}) (int64, error) {
	userModel := &User.UserInfo{}
	db, err := Utils.GetConnByKey("")
	if err != nil {
		s.log.Error("db connection error", err.Error())
		return -1, err
	}
	iSession := db.Begin()
	iSession = iSession.Model(userModel).
		Set("gorm:query_condition", "FOR UPDATE").
		Updates(conditon)
	if iSession.Error != nil {
		iSession.Rollback()
		return -1, iSession.Error
	}
	if err := iSession.Commit(); err != nil {
		return -1, iSession.Error
	}
	affectedRows := iSession.RowsAffected
	return affectedRows, err
}
