package public

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/garyburd/redigo/redis"
)

func RedisConfPipline(pip ...func(c redis.Conn)) error {
	conn, err := lib.RedisConnFactory("default")
	if err!=nil{
		return err
	}
	defer conn.Close()
	for _,f:=range pip{
		f(conn)
	}
	conn.Flush()
	return nil
}

func RedisConfDo(commonName string,args ...interface{}) (interface{},error) {
	conn, err := lib.RedisConnFactory("default")
	if err!=nil{
		return nil, err
	}
	defer conn.Close()
	return conn.Do(commonName,args...)
}
