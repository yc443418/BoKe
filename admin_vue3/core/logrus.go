// 日志配置
// @author chen

package core

import (
	"admin_vue3/config"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	log := config.Config.Logger
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-0102 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s: %s\n",
			log.Prefix, timestamp, levelColor, entry.Level, funcVal, fileVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n",
			log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// INitLogger 初始化
func InitLogger() *logrus.Logger {
	mlog := logrus.New()                                // 新建一个实例
	mlog.SetOutput(os.Stdout)                           // 设置输出类型
	mlog.SetReportCaller(config.Config.Logger.ShowLine) // 开启返回函数名和行号
	mlog.SetFormatter(&LogFormatter{})                  // 设置自己定义的Formatter
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mlog.SetLevel(level) // 设置最低的level
	InitDefaultLogger()  // 不注释即启用全局log
	return mlog
}

// InitDefaultLogger 全局log
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)                           //设置输出类型
	logrus.SetReportCaller(config.Config.Logger.ShowLine) // 开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                  // 设置自己定义的Formatter
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低的level
}
