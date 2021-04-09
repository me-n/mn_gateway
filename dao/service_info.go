package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"mn_gateway/dto"
	"mn_gateway/public"
	"time"
)

//服务信息
type ServiceInfo struct {
	Id          int64       `json:"id" gorm:"primary_key"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	CreatedAt    time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	UpdatedAt    time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	IsDelete    int       `json:"is_delete" gorm:"column:is_delete" description:"是否已删除 0：是 1：否"`
}

//数据库表单地址
func (*ServiceInfo) TableName() string {
	return "service_info"
}

//分页处理
func (so *ServiceInfo) PageList(c *gin.Context, db *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := int64(0)
	list := []ServiceInfo{}
	//计算偏移量
	offset := (param.PageNo - 1) * param.PageSize
	query := db.SetCtx(public.FromGinTraceContext(c))
	query.Table(so.TableName()).Where("is_delete=0")
	if param.Info != "" {
		query = query.Where("(service_name like ? or service_desc like ?)", "%"+param.Info+"%", "%"+param.Info+"%")
	}
	err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}

//查找
func (so *ServiceInfo) Find(c *gin.Context, db *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	outInfo := &ServiceInfo{}
	err := db.SetCtx(public.FromGinTraceContext(c)).Where(search).Find(outInfo).Error
	if err != nil {
		return nil, err
	}
	return outInfo, nil
}

//保存
func (so *ServiceInfo) Save(c *gin.Context, db *gorm.DB) error {
	return db.SetCtx(public.FromGinTraceContext(c)).Save(so).Error
}

func (so *ServiceInfo)ServiceDetail(c *gin.Context,db *gorm.DB, search *ServiceInfo) (*ServiceDetail,error) {
	if search.ServiceName==""{
		info ,err:=so.Find(c,db,search)
		if err!=nil{
			return nil, err
		}
		search=info
	}
	httpRule:=&HttpRule{ServiceId: search.Id}
	httpRule, err := httpRule.Find(c, db, httpRule)
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	tcpRule:=&TcpRule{ServiceId: search.Id}
	tcpRule, err = tcpRule.Find(c, db, tcpRule)
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	grpcRule:=&GrpcRule{ServiceId: search.Id}
	grpcRule, err = grpcRule.Find(c, db, grpcRule)
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	accessControl:=&AccessControl{ServiceId: search.Id}
	accessControl, err = accessControl.Find(c, db, accessControl)
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	loadBalance:=&LoadBalance{ServiceId: search.Id}
	loadBalance, err = loadBalance.Find(c, db, loadBalance)
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil, err
	}
	detail:=&ServiceDetail{
		Info: search,
		AccessControl: accessControl,
		HTTPRule: httpRule,
		TCPRule: tcpRule,
		GrpcRule: grpcRule,
	}
	return detail,nil
}
