package internal

import (
	"blog-backend/global"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

func (r *fileRotateLogs) GetWriterSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.YAGAMI_CONFIG.Zap.Director, "%Y-%m-%d"+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.YAGAMI_CONFIG.Zap.MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.YAGAMI_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
