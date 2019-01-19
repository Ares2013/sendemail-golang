package dao

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/prometheus/common/log"
	"sync"
)

type DbType struct {
	MasterDbConfig DbConfig
	SlaveDbConfig DbConfig
	DriverName string
}
type DbConfig struct{
	User string
	Pwd	 string
	Host string
	Port int
	DbName string
	Charset string
	DriverName string
}
// db 连接的配置
var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)
// Load json config file
//config.LoadFile("/tmp/config.json")
var conf = DbType{MasterDbConfig:DbConfig{User:"ws_cbs",Pwd:"wscbs123",Host:"192.168.10.240",Port:3306,DbName:"cms_global",Charset:"utf8",DriverName:"mysql"},SlaveDbConfig:DbConfig{User:"ws_cbs",Pwd:"wscbs123",Host:"192.168.10.240",Port:3306,DbName:"cms_global",Charset:"utf8",DriverName:"mysql"},DriverName:"mysql"}

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()
	if masterEngine != nil {
		return masterEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("DbConnection connect db error=", err)
		return nil
	} else {
		masterEngine = engine
		return masterEngine
	}
}
func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()
	if slaveEngine != nil {
		return slaveEngine
	}
	c := conf.SlaveDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("DbConnection connect db error=", err)
		return nil
	} else {
		slaveEngine = engine
		return slaveEngine
	}
}
