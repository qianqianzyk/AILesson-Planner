package logs

import (
	"github.com/qianqianzyk/AILesson-Planner/internal/config"
	"go.uber.org/zap"
)

func ZapInit(c config.Config) error {
	zapInfo := InfoConfig{
		StacktraceLevel:   "warn",
		DisableStacktrace: c.Log.DisableStacktrace, // 是否禁用堆栈跟踪
		ConsoleLevel:      c.Log.Level,             // 日志级别
		Name:              c.Log.Name,              // 日志名称
		Writer:            c.Log.Writer,            // 日志输出方式
		LoggerDir:         c.Log.LoggerDir,         // 日志目录
		LogCompress:       c.Log.LogCompress,       // 是否压缩日志
		LogMaxSize:        c.Log.LogMaxSize,        // 日志文件最大大小（单位：MB）
		LogMaxAge:         c.Log.LogMaxAge,         // 日志保存天数
	}
	logger, err := Init(&zapInfo)
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	zap.L().Info("Logger initialized")
	return nil
}
