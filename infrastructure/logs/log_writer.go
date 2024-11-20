package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"strings"
	"time"
)

// 使用rotate 进行日志的分割
func getLogWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.logs.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"%Y%m%d.%H.log", // 没有使用go风格反人类的format格式
		rotatelogs.WithRotationTime(time.Hour),                    // 设置每小时日志
		rotatelogs.WithMaxAge(time.Hour*30*24),                    // 设置30天有效期
	)
	if err != nil {
		panic(err)
	}
	return hook
}
