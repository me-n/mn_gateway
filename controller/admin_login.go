package controller

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mn_gateway/dao"
	"mn_gateway/dto"
	"mn_gateway/middleware"
	"mn_gateway/public"
	"time"
)

type AdminLoginController struct{}
//管理员登陆注册，隶属admin——login注册组
func AdminLoginRegister(rGroup *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	rGroup.POST("/login", adminLogin.AdminLoginIn)
	rGroup.GET("/logout", adminLogin.AdminLoginOut)
}
// AdminLoginIn godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept json
// @Produce json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminLogin *AdminLoginController) AdminLoginIn(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BlindValidParam(c); err != nil {
		middleware.ResponseError(c, 4001, err)
		return
	}
	db, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 4002, err)
		return
	}
	admin := &dao.Admin{}
	admin, err = admin.AdminLoginCheckPwd(c, db, params)
	if err != nil {
		middleware.ResponseError(c, 4003, err)
		return
	}
	//设置session
	sessionInfo := &dto.AdminSessionInfo{
		Id:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessionInfo)
	if err != nil {
		middleware.ResponseError(c, 4004, err)
	}
	sess := sessions.Default(c)
	sess.Set(public.AdminSessionInfoKey, string(sessBts))
	sess.Save()
	output := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, output)
}

// AdminLoginIn godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /admin_login/logout
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/logout [get]
func (*AdminLoginController) AdminLoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(public.AdminSessionInfoKey)
	session.Save()
	middleware.ResponseSuccess(c, "登出完成")
}
