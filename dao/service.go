package dao

type ServiceDetail struct {
	Info *ServiceInfo `json:"info" description:"基本信息"`
	AccessControl *AccessControl `json:"access_control" description:"access_control"`
	HTTPRule *HttpRule `json:"http_rule" description:"http_rule"`
	TCPRule *TcpRule `json:"tcp_rule" description:"tcp_rule"`
	GrpcRule *GrpcRule `json:"grpc_rule" description:"grpc_rule"`
	LoadBalance *LoadBalance `json:"load_balance" description:"load_balance"`
}

