package conf

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"io/ioutil"
	"time"
)

type DBConfig struct {
	DBuser     string `json:"db_user"`
	DBpassword string `json:"db_password"`
	DBip       string `json:"db_ip"`
	DBport     string `json:"db_port"`
	DBname     string `json:"db_name"`
}

type Sysconfig struct {
	DBConfig    DBConfig           `json:"db_config"`
	IrisConfig  iris.Configuration `json:"iris_config"`
	RedisConfig redis.Config       `json:"redis_config"`
}

var (
	Config = &Sysconfig{}
	//pwd, _  = os.Getwd()
	//path    = filepath.Dir(pwd)
)

func init() {
	configPath := "config.json"
	//configPath := filepath.Join(path,"config.json")
	//读取json配置
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, Config)
	if err != nil {
		fmt.Println(err)
		return
	}
	//log.Println("Start loading config")
	Config.RedisConfig.Clusters = nil
	Config.RedisConfig.Timeout = time.Duration(30) * time.Second
	Config.RedisConfig.Driver = redis.Redigo()

	//// 日志格式化为 JSON 而不是默认的 ASCII
	//log.SetFormatter(&log.JSONFormatter{})
	//// 输出 stdout 而不是默认的 stderr，也可以是一个文件
	//log.SetOutput(os.Stdout)
	//// 只记录严重或以上警告
	//log.SetLevel(log.WarnLevel)

}
