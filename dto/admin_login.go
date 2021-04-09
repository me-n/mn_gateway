package dto

import (
	"github.com/gin-gonic/gin"
	"mn_gateway/public"
	"time"
)
//管理员登陆
type AdminLoginInput struct {
	UserName string `json:"user_name" form:"user_name" comment:"用户名" example:"admin" validate:"required"`//管理员用户名
	Password string `json:"password" form:"password" comment:"密码" example:"1223456" validate:"required"`//密码
}
//管理员登出
type AdminLoginOutput struct {
	Token string `json:"token" form:"toke" comment:"token" example:"token" validate:""`//Token
}
//session结构体
type AdminSessionInfo struct {
	Id        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}

func (admIpt *AdminLoginInput) BlindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, admIpt)
}
