/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 2:21 下午
 */

package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
)

type Config struct {
	RedisConf RedisConf `toml:"redis_conf"`
	MysqlConf MysqlConf `toml:"database"`
	LogStruct LogStruct `toml:"log"`
}

var config = &Config{}

func GetConfig() *Config {
	return config
}

func init() {

	// 程序运行的时候，通过参数传过去 ./* -c ./etc/config.toml
	appConfigPath := ""
	flag.StringVar(&appConfigPath, "c", "./config-test.toml", "config path")
	flag.Parse()

	toml.DecodeFile(appConfigPath, config)
}
