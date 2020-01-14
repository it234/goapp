package config

import (
	"github.com/it234/goapp/internal/pkg/config"

	"github.com/spf13/viper"
)

// 加载配置
func LoadConfig(fpath string) (c *config.Config, err error) {
	v := viper.New()
	v.SetConfigFile(fpath)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &config.Config{}
	c.Web.StaticPath = v.GetString("web.static_path")
	c.Web.Domain = v.GetString("web.domain")
	c.Web.Port = v.GetInt("web.port")
	c.Web.ReadTimeout = v.GetInt("web.read_timeout")
	c.Web.WriteTimeout = v.GetInt("web.write_timeout")
	c.Web.IdleTimeout = v.GetInt("web.idle_timeout")
	c.MySQL.Host = v.GetString("mysql.host")
	c.MySQL.Port = v.GetInt("mysql.port")
	c.MySQL.User = v.GetString("mysql.user")
	c.MySQL.Password = v.GetString("mysql.password")
	c.MySQL.DBName = v.GetString("mysql.db_name")
	c.MySQL.Parameters = v.GetString("mysql.parameters")
	c.Sqlite3.Path = v.GetString("sqlite3.path")
	c.Gorm.Debug = v.GetBool("gorm.debug")
	c.Gorm.DBType = v.GetString("gorm.db_type")
	c.Gorm.MaxLifetime = v.GetInt("gorm.max_lifetime")
	c.Gorm.MaxOpenConns = v.GetInt("gorm.max_open_conns")
	c.Gorm.MaxIdleConns = v.GetInt("gorm.max_idle_conns")
	c.Gorm.TablePrefix = v.GetString("gorm.table_prefix")
	return
}
