package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mn_gateway/dao"
	"mn_gateway/dto"
	"mn_gateway/middleware"
	"mn_gateway/public"
)

type AdminController struct{}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/info", adminLogin.AdminInfo)
	group.POST("/changePwd", adminLogin.AdminChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/info
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/info [get]
func (ac *AdminController) AdminInfo(c *gin.Context) {
	session := sessions.Default(c)
	getSessionInfo := session.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(getSessionInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 4010, err)
		return
	}
	out := &dto.AdminInfoOutput{
		Id:           adminSessionInfo.Id,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fb-ssl.duitang.com%2Fuploads%2Fitem%2F201905%2F28%2F20190528143150_fETNW.thumb.700_0.jpeg&refer=http%3A%2F%2Fb-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1620203840&t=0d80cb1c778cd414fe779b82625bde73",
		Introduction: "这是一个关于超级管理员的介绍",
		Roles:        []string{"超级管理员","管理员"},
	}
	middleware.ResponseSuccess(c, out)
}

// AdminChangePwd godoc
// @Summary 管理员密码修改
// @Description 管理员密码修改
// @Tags 管理员接口
// @ID /admin/changePwd
// @Accept json
// @Produce json
// @Param body body dto.ChangeAdmPwd true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/changePwd [post]
func (ac *AdminController) AdminChangePwd(c *gin.Context) {
	param := &dto.ChangeAdmPwd{}
	if err := param.BindValidateParam(c); err != nil {
		middleware.ResponseError(c, 4011, err)
		return
	}
	//读取session到结构体
	session := sessions.Default(c)
	getSessInfo := session.Get(public.AdminSessionInfoKey)
	adminSessInfo := &dto.AdminSessionInfo{}
	err := json.Unmarshal([]byte(fmt.Sprint(getSessInfo)), adminSessInfo)
	if err != nil {
		middleware.ResponseError(c, 4012, err)
		return
	}
	//从数据库中读取管理员信息
	db, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 4013, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.AdminFindFromDb(c, db, (&dao.Admin{UserName: adminSessInfo.UserName}))
	if err != nil {
		middleware.ResponseError(c, 4014, err)
		return
	}
	//对新密码加盐
	newSaltPwd := public.PasswordAddSalt(adminInfo.Salt, param.Password)
	adminInfo.Password = newSaltPwd
	//保存新密码到数据库
	if err := adminInfo.AdminSaveToDb(c, db); err != nil {
		middleware.ResponseError(c, 4015, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

