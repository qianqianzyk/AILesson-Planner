package logs

import "go.uber.org/zap"

type Level uint8

const (
	LevelFatal  Level = 0
	LevelPanic  Level = 1
	LevelDpanic Level = 2
	LevelError  Level = 3
	LevelWarn   Level = 4
	LevelInfo   Level = 5
	LevelDebug  Level = 6
)

func GetLogFunc(level Level) func(string, ...zap.Field) {
	// 创建日志级别映射表
	logMap := map[Level]func(string, ...zap.Field){
		LevelFatal:  zap.L().Fatal,
		LevelPanic:  zap.L().Panic,
		LevelDpanic: zap.L().DPanic,
		LevelError:  zap.L().Error,
		LevelWarn:   zap.L().Warn,
		LevelInfo:   zap.L().Info,
		LevelDebug:  zap.L().Debug,
	}

	if logFunc, ok := logMap[level]; ok {
		return logFunc
	}
	return zap.L().Info
}
