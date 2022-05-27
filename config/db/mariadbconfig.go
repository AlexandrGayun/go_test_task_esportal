package mariadbconfig

import (
	"github.com/go-sql-driver/mysql"
	"net"
	"os"
	"strings"
)

func newConfig(envPrefix string) *mysql.Config {
	envPrefix = strings.ToUpper(envPrefix)
	envName := func(env string) string { return envPrefix + env }
	config := mysql.NewConfig()
	config.User = os.Getenv(envName("DB_USER"))
	config.Passwd = os.Getenv(envName("DB_PASSWORD"))
	config.DBName = os.Getenv(envName("DB_NAME"))
	config.ParseTime = true
	config.Net = "tcp"
	host := os.Getenv(envName("DB_HOST"))
	port := os.Getenv(envName("DB_PORT"))
	config.Addr = net.JoinHostPort(host, port)
	return config
}

func GetDSN(envPrefix string) string {
	return newConfig(envPrefix).FormatDSN()
}
