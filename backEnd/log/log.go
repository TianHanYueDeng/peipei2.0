package log

import (
	log "github.com/sirupsen/logrus"

)


var logPath string


func init() {
	//pwd, _  := os.Getwd()
	//path    := filepath.Dir(pwd)
	//logPath = filepath.Join(path, "log.txt")
	//fmt.Println(logPath)
	//f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	//if err != nil{
	//	return
	//}
	//outPut := io.Writer(f)
	// 设置日志格式为json格式
	//log.SetFormatter(&log.TextFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	//log.SetOutput(outPut)

	// 增加hook,输出日行号与代码位置
	log.AddHook(lineHook{Field:"source"})
	log.AddHook(newLfsHook(5, 24))

	// 设置日志级别为warn以上
	//log.SetLevel(log.WarnLevel)
}
