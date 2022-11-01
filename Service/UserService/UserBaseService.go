package UserService

import (
	"MetaWebServer/Repo/UserRepo"
	"MetaWebServer/Utils"
)

type UserBaseService interface {
	Get(map[string]interface{}) Utils.ResponseInfo
}

type UserBaseResponse struct {
	Model *UserRepo.UserBase
}

func (s *UserBaseResponse) Get(condition map[string]interface{}) Utils.ResponseInfo {
	var res Utils.ResponseInfo
	res.Code = Utils.SuccessCode
	res.Msg = Utils.CommonFailMsg
	repo := UserRepo.UserBaseRepo(s.Model)
	users, err := repo.GetAllUser(condition)
	if err != nil {
		res.Msg = err.Error()
		return res
	}
	if users != nil {
		if err == nil {
			res.Data = users
			res.Code = Utils.SuccessCode
			res.Msg = Utils.CommonSuccessMsg
		}
	}
	return res
}
