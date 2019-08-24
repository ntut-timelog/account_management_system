package config

import "github.com/spf13/viper"

const (
	mysqlHost = "MysqlHost"
	mysqlUser = "MysqlUser"
	mysqlPwd  = "MysqlPwd"
	mysqlDB   = "MysqlDB"
)

func init() {
	viper.SetDefault(mysqlHost, "localhost")
	viper.SetDefault(mysqlUser, "root")
	viper.SetDefault(mysqlPwd, "password")
	viper.SetDefault(mysqlDB, "database")

	viper.BindEnv(mysqlHost, "MYSQL_HOST")
	viper.BindEnv(mysqlUser, "MYSQL_USER")
	viper.BindEnv(mysqlPwd, "MYSQL_PWD")
	viper.BindEnv(mysqlDB, "MYSQL_DB")
}

type mySQL struct {
	Host     string
	User     string
	Password string
	Database string
}

var mysql *mySQL

func MySQL() *mySQL {
	if mysql == nil {
		mysql = &mySQL{
			Host:     viper.Get(mysqlHost).(string),
			User:     viper.Get(mysqlUser).(string),
			Password: viper.Get(mysqlPwd).(string),
			Database: viper.Get(mysqlDB).(string),
		}
	}
	return mysql
}
