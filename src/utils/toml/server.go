/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 2:22 下午
 */

package conf

// redis config
type RedisConf struct {
	Host   string `toml:"host"`
	Port   string `toml:"port"`
	PassWd string `toml:"passWd"`
	Db     int    `toml:"db"`
}

// mysql config
type MysqlConf struct {
	Db         string `toml:"db"`
	DbHost     string `toml:"dbHost"`
	DbPort     string `toml:"dbPort"`
	DbUser     string `toml:"dbUser"`
	DbPassWord string `toml:"dbPassWord"`
}

// log config
type LogStruct struct {
	LogPath    string `json:"logPath"`
	MaxAge     int    `json:"maxAge"`
	RotateTime int    `json:"rotateTime"`
	Level      int    `json:"level"`
}
