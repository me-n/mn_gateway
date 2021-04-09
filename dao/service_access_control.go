package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"mn_gateway/middleware"
	"mn_gateway/public"
)

type AccessControl struct {
	Id                int64  `json:"id" gorm:"primary_key"`
	ServiceId         int64  `json:"service_id" gorm:"column:service_id" description:"服务ID"`
	OpenAuth          int    `json:"open_auth" gorm:"column:open_auth" description:"是否开启权限 开启=1"`
	WhiteList         string `json:"white_list" gorm:"column:white_list" description:"白名单ip"`
	BlackList         string `json:"black_list" gorm:"column:black_list" description:"黑名单ip"`
	WhiteHostName     string `json:"white_host_name" gorm:"column:white_host_name" description:"白名单主机"`
	ClientIpFlowLimit int    `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流"`
	ServiceFlowLimit  int    `json:"service_flow_limit" gorm:"column:service_flow_limit" description:"服务端限流"`
}

func (*AccessControl) TableName() string {
	return "service_access_control"
}

func (*AccessControl) Find(c *gin.Context, db *gorm.DB, search *AccessControl) (*AccessControl, error) {
	acl := &AccessControl{}
	err := db.SetCtx(public.FromGinTraceContext(c)).Where(search).Error
	if err != nil {
		middleware.ResponseError(c, 4040, err)
		return nil, err
	}
	return acl, nil
}

func (acl *AccessControl) Save(c *gin.Context, db *gorm.DB) error {
	err := db.SetCtx(public.FromGinTraceContext(c)).Save(acl).Error
	if err != nil {
		middleware.ResponseError(c, 4041, err)
		return err
	}
	return nil
}

func (acl *AccessControl) ListByServiceId(c *gin.Context, db *gorm.DB, serviceId int64) ([]AccessControl, int64, error) {
	aclList := make([]AccessControl, 0)
	var count int64
	query := db.SetCtx(public.FromGinTraceContext(c)).Table(acl.TableName()).Select("*").Where("service_id=?", serviceId).Order("id desc").Find(&aclList)
	err := query.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	err = query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return aclList, count, nil
}
