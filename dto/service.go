package dto

import (
	"github.com/gin-gonic/gin"
	"mn_gateway/public"
)

//服务列表输出结构体
type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键字" example:"" validate:""`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"`
}

//服务输出列表具体项目结构体
type ServiceListItemOutput struct {
	Id          int64  `json:"id" form:"id"`
	LoadType    int    `json:"load_type" form:"load_type"`
	ServiceName string `json:"service_name" form:"service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	ServiceAddr string `json:"service_addr" form:"service_addr"`
	Qps         int64  `json:"qps" form:"qps"`
	Qpd         int64  `json:"qpd" form:"qpd"`
	TotalNode   int    `json:"total_node" form:"total_node"`
}

//服务输出列表结构体
type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" example:"" validate:""`
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表" example:"" validate:""`
}

//删除服务输入结构体
type ServiceDeleteInput struct {
	Id int64 `json:"id" form:"id" comment:"服务ID" example:"" validate:"required"`
}

//服务流量输出结构体
type ServiceStatOutput struct {
	Today     []int64 `json:"today" form:"today" comment:"今日流量" example:"" validate:""`
	Yesterday []int64 `json:"yesterday" form:"yesterday" comment:"昨日流量" example:"" validate:""`
}

//添加http 输入
type ServiceAddHttpInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,min=1,max=255"`

	RuleType       int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"" validate:"min=0,max=1"`
	Rule           string `json:"rule" form:"rule" comment:"接入路径：前缀或域名" example:"" validate:"required,valid_rule"`
	NeedHttps      int    `json:"need_https" form:"need_https" comment:"支持https" example:"" validate:"min=0,max=1"` //是否开启支持https
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"" validate:"min=0,max=1"`
	NeedStripUri   int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri" example:"" validate:"min=0,max=1"`
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"url重写" example:"" validate:"valid_url_rewrite"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header转换" example:"" validate:"valid_header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"开启权限" example:"" validate:"min=0,max=1"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP" example:"" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP" example:"" validate:""`
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:"min=0,max=3"`
	IpList                 string `json:"ip_list" form:"ip_list" comment:"IP列表" example:"" validate:"required,valid_ipportlist"`
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重轮询" example:"" validate:"required,valid_weightlist"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时，单位s" example:"" validate:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时，单位s" example:"" validate:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间，单位s" example:"" validate:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"" validate:"min=0"`
}

//更新http 输入
type ServiceUpdateHttpInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,min=1,max=255"`

	RuleType       int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"" validate:"min=0,max=1"`
	Rule           string `json:"rule" form:"rule" comment:"接入路径：前缀或域名" example:"" validate:"required,valid_rule"`
	NeedHttps      int    `json:"need_https" form:"need_https" comment:"支持https" example:"" validate:"min=0,max=1"` //是否开启支持https
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"" validate:"min=0,max=1"`
	NeedStripUri   int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri" example:"" validate:"min=0,max=1"`
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"url重写" example:"" validate:"valid_url_rewrite"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header转换" example:"" validate:"valid_header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"开启权限" example:"" validate:"min=0,max=1"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP" example:"" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP" example:"" validate:""`
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:"min=0,max=3"`
	IpList                 string `json:"ip_list" form:"ip_list" comment:"IP列表" example:"" validate:"required,valid_ipportlist"`
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重轮询" example:"" validate:"required,valid_weightlist"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时，单位s" example:"" validate:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时，单位s" example:"" validate:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间，单位s" example:"" validate:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"" validate:"min=0"`
}

//添加Tcp 输入
type ServiceAddTcpInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" example:"" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口设置（8001-8999）" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"header转换" example:"" validate:""`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" example:"" validate:""`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，逗号间隔，优先级：白名单高于黑名单" example:"" validate:"valid_iplist"`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，逗号间隔，优先级：白名单高于黑名单" example:"" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，逗号间隔" validate:"valid_iplist"`
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" example:"" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" example:"" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

//更新Tcp 输入
type ServiceUpdateTcpInput struct {
	Id                int    `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" example:"" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口设置（8001-8999）" validate:"required,min=8001,max=8999"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" example:"" validate:""`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，逗号间隔，优先级：白名单高于黑名单" example:"" validate:"valid_iplist"`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，逗号间隔，优先级：白名单高于黑名单" example:"" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，逗号间隔" validate:"valid_iplist"`
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" example:"" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" example:"" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

//添加grpc 输入
type ServiceAddGrpcInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

//更新grpc 输入
type ServiceUpdateGrpcInput struct {
	Id                int    `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

//服务输入列表绑定到上下文
func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

//删除服务输入绑定到上下文
func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

//添加http绑定到上下文
func (param *ServiceAddHttpInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

//更新http绑定到上下文
func (param *ServiceUpdateHttpInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

//添加tcp绑定到上下文
func (param *ServiceAddTcpInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

//更新tcp绑定到上下文
func (param *ServiceUpdateTcpInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

//添加grpc绑定到上下文
func (param *ServiceAddGrpcInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

//更新grpc绑定到上下文
func (param *ServiceUpdateGrpcInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
