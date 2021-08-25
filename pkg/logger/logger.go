package logger

import "log"

// LogError 当存在错误时记录日志
func LogError(err error) string {
	if err != nil {
		log.Println(err)

		return err.Error()
	} else {
		return "请求成功"
	}
}
