package controller

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"mn_gateway/dao"
	"mn_gateway/dto"
	"mn_gateway/middleware"
	"mn_gateway/public"
)

type ServiceController struct{}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
	group.GET("/service_delete", service.ServiceDelete)
	group.GET("/service_detail", service.ServiceDetail)
	group.GET("/service_stat", service.ServiceStat)

	group.POST("/service_add_http", service.ServiceAddHTTP)
	group.POST("/service_update_http", service.ServiceUpdateHTTP)
	group.POST("/service_add_tcp", service.ServiceAddTcp)
	group.POST("/service_update_tcp", service.ServiceUpdateTcp)
	group.POST("/service_add_grpc", service.ServiceAddGrpc)
	group.POST("/service_update_grpc", service.ServiceUpdateGrpc)
}
// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (sc *ServiceController) ServiceList(c *gin.Context) {
	param := &dto.ServiceListInput{}
	if err := param.BindValidParam(c);err!=nil{
		middleware.ResponseError(c,4021,err)
		return
	}
	db, err := lib.GetGormPool("default")
	if err!=nil{
		middleware.ResponseError(c,4022,err)
		return
	}
	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(c, db, param)
	if err!=nil{
		middleware.ResponseError(c,4023,err)
		return
	}
	outList:=[]dto.ServiceListItemOutput{}
	for _,listItem:=range list{
		serviceDetail, err := listItem.ServiceDetail(c, db, &listItem)
		if err!=nil{
			middleware.ResponseError(c,4024,err)
			return
		}
		serviceAddr:="unknow"
		clusterIp:=lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort:=lib.GetStringConf("base.cluster.cluster_port")
		clusterSSLPort:=lib.GetStringConf("base.cluster.cluster_ssl_port")
		if serviceDetail.Info.LoadType==public.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL &&
			serviceDetail.HTTPRule.NeedHttps == 1 {
			serviceAddr =fmt.Sprintf("%s:%s%s",clusterIp,clusterSSLPort,serviceDetail.HTTPRule.Rule)
		}
		if serviceDetail.Info.LoadType==public.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL &&
			serviceDetail.HTTPRule.NeedHttps == 0 {
			serviceAddr =fmt.Sprintf("%s:%s%s",clusterIp,clusterPort,serviceDetail.HTTPRule.Rule)
		}
		if serviceDetail.Info.LoadType==public.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType ==public.HTTPRuleTypeDomain{
			serviceAddr=serviceDetail.HTTPRule.Rule
		}
		if serviceDetail.Info.LoadType==public.LoadTypeTCP{
			serviceAddr=fmt.Sprintf("%s:%d",clusterIp,serviceDetail.TCPRule.Port)
		}
		if serviceDetail.Info.LoadType==public.LoadTypeGRPC{
			serviceAddr=fmt.Sprintf("%s:%d",clusterIp,serviceDetail.GrpcRule.Port)
		}
		ipList := serviceDetail.LoadBalance.GetIpListByModel()
		counter, err := public.FlowCounterHander.GetCounter(public.FlowServicePrefix + listItem.ServiceName)
		if err!=nil{
			middleware.ResponseError(c,4025,err)
			return
		}
		output := dto.ServiceListItemOutput{
			Id: listItem.Id,
			LoadType:  listItem.LoadType,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			ServiceAddr: serviceAddr,
			Qps: counter.QPS,
			Qpd: counter.TotalCount,
			TotalNode: len(ipList),
		}
		outList=append(outList,output)
	}
	out:=&dto.ServiceListOutput{
		Total: total,
		List: outList,
	}
	middleware.ResponseSuccess(c,out)
}

func (sc *ServiceController) ServiceDelete(c *gin.Context) {

}

func (sc *ServiceController) ServiceDetail(c *gin.Context) {

}

func (sc *ServiceController) ServiceStat(c *gin.Context) {

}

func (sc *ServiceController) ServiceAddHTTP(c *gin.Context) {

}

func (sc *ServiceController) ServiceUpdateHTTP(c *gin.Context) {

}

func (sc *ServiceController) ServiceAddTcp(c *gin.Context) {

}

func (sc *ServiceController) ServiceUpdateTcp(c *gin.Context) {

}

func (sc *ServiceController) ServiceAddGrpc(c *gin.Context) {

}

func (sc *ServiceController) ServiceUpdateGrpc(c *gin.Context) {

}
