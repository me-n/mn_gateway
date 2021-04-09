package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"mn_gateway/middleware"
	"mn_gateway/public"
)

type TcpRule struct {
	Id             int64  `json:"id" gorm:"primary_key"`
	ServiceId      int64  `json:"service_id" gorm:"column:service_id" description:"服务Id"`
	Port           int    `json:"port" gorm:"column:port" description:"端口"`
	HeaderTransfor string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换，增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (*TcpRule) TableName() string {
	return "service_tcp_rule"
}

func (*TcpRule) Find(c *gin.Context, db *gorm.DB, search *TcpRule) (*TcpRule, error) {
	tpr := &TcpRule{}
	err := db.SetCtx(public.FromGinTraceContext(c)).Where(search).Find(tpr).Error
	if err != nil {
		middleware.ResponseError(c, 4060, err)
		return nil, err
	}
	return tpr, nil
}

func (tpr *TcpRule) Save(c *gin.Context, db *gorm.DB) error {
	err := db.SetCtx(public.FromGinTraceContext(c)).Save(tpr).Error
	if err != nil {
		middleware.ResponseError(c, 4061, err)
		return err
	}
	return nil
}

func (tpr *TcpRule) ListByServiceId(c *gin.Context, db *gorm.DB, serviceId int64) ([]TcpRule, int64, error) {
	var list []TcpRule
	var count int64
	find := db.SetCtx(public.FromGinTraceContext(c)).Table(tpr.TableName()).Select("*").Where("service_id=?",serviceId).Order("id desc").Find(&list)
	if err := find.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	err := find.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
