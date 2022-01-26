package writer

import (
	"time"

	"github.com/anchore/grype/internal/log"
	"gorm.io/gorm/logger"
)

type logAdapter struct {
}

func (l *logAdapter) LogMode(level logger.LogLevel) logger.Interface {
	//TODO implement me
	panic("implement me")
}

func (l *logAdapter) Info(ctx context.Context, s string, i ...interface{}) {
	log.Info(s, i)
}

func (l *logAdapter) Warn(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *logAdapter) Error(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *logAdapter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
	panic("implement me")
}

func (l *logAdapter) Print(v ...interface{}) {
	log.Error(v...)
}
