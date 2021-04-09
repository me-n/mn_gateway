package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"mn_gateway/public"
)

type HttpRule struct {
	Id             int64  `json:"id" gorm:"primary_key"`
	ServiceId      int64  `json:"service_id" gorm:"column:service_id" description:"服务ID"`
	RuleType       int    `json:"rule_type" gorm:"column:rule_type" description:"匹配类型 domain=域名, url_prefix=url前缀"`
	Rule           string `json:"rule" gorm:"column:rule" description:"type=domain表示域名，type=url_prefix时表示url前缀"`
	NeedHttps      int    `json:"need_https" gorm:"column:need_https" description:"type=支持https 1=支持"`
	NeedWebsocket  int    `json:"need_websocket" gorm:"column:need_websocket" description:"启用websocket 1=启用"`
	NeedStripUri   int    `json:"need_strip_uri" gorm:"column:need_strip_uri" description:"启用strip_uri 1=启用"`
	UrlRewrite     string `json:"url_rewrite" gorm:"column:url_rewrite" description:"url重写功能，每行一个"`
	HeaderTransfor string `json:"header_transfor" gorm:"header_transfor" description:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (*HttpRule) TableName() string {
	return "service_http_rule"
}

func (hr *HttpRule) Find(c *gin.Context, db *gorm.DB, search *HttpRule) (*HttpRule, error) {
	model := &HttpRule{}
	err := db.SetCtx(public.FromGinTraceContext(c)).Where(search).Find(&model).Error
	return model, err
}

func (hr *HttpRule) Save(c *gin.Context, db *gorm.DB) error {
	err := db.SetCtx(public.FromGinTraceContext(c)).Save(hr).Error
	if err != nil {
		return err
	}
	return nil
}

func (hr *HttpRule) ListByServiceId(c *gin.Context, db *gorm.DB, serviceId int64) ([]HttpRule, int64, error) {
	var list []HttpRule
	var count int64
	find := db.SetCtx(public.FromGinTraceContext(c)).Table(hr.TableName()).Select("*").Where("service_id=?", serviceId).Order("id desc").Find(&list)
	err := find.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	err = find.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
