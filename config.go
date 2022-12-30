// author: asydevc <asydev@163.com>
// date: 2021-02-13

package db

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
	"xorm.io/xorm/log"

	"github.com/asydevc/log/v2"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// DB配置.
type configuration struct {
	Driver      string   `yaml:"driver"`
	Dsn         []string `yaml:"dsn"`
	MaxIdle     int      `yaml:"max-idle"`
	MaxOpen     int      `yaml:"max-open"`
	MaxLifetime int      `yaml:"max-lifetime"`
	Mapper      string   `yaml:"mapper"`
	engines     *xorm.EngineGroup
	slaveEnable bool
}

// 读取YAML文件.
func (o *configuration) LoadYaml(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, o); err != nil {
		return err
	}

	log.Debugf("parse configuration from %s.", path)
	o.parse()
	return nil
}

// 初始化配置.
func (o *configuration) initialize() {
	for _, path := range []string{"./tmp/db.yaml", "./config/db.yaml", "../config/db.yaml", "../../configs/db.yaml"} {
		if o.LoadYaml(path) == nil {
			break
		}
	}
}

// 应用配置.
func (o *configuration) parse() {
	var err error
	if o.engines, err = xorm.NewEngineGroup(o.Driver, o.Dsn); err != nil {
		panic(err)
	}

	log.Debugf("config %d dsn, max idle is %d, max open is %d, mapper is %s.", len(o.Dsn), o.MaxIdle, o.MaxOpen, o.Mapper)

	o.engines.SetConnMaxLifetime(time.Duration(o.MaxLifetime) * time.Second)
	o.engines.SetMaxIdleConns(o.MaxIdle)
	o.engines.SetMaxOpenConns(o.MaxOpen)
	o.engines.SetLogger(plugins.NewXOrm())
	o.slaveEnable = len(o.Dsn) > 1

	if o.Mapper == "same" {
		o.engines.SetColumnMapper(names.SameMapper{})
	} else if o.Mapper == "snake" {
		o.engines.SetColumnMapper(names.SnakeMapper{})
	}
}
