package dto

import (
	"github.com/gin-gonic/gin"
	"mn_gateway/public"
	"time"
)

type AdminInfoOutput struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangeAdmPwd struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`//密码
}

func (param *ChangeAdmPwd) BindValidateParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
