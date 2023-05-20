package logger

import (
	"go.uber.org/zap"
)

// razbor
func New(version, description string) *zap.Logger {
	var log *zap.Logger
	var err error

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout", description + "-" + version + ".log"}

	if err != nil {
		log = zap.NewExample()
		log.Warn("Unable to set up the logger. Replaced with example one which shouldn't fail", zap.Error(err))
	}
	zap.ReplaceGlobals(log)
	err = log.Sync()
	if err != nil {
		log.Warn("Logger sync fail", zap.Error(err))
	}
	log.Debug("Logger is ready")
	return log
}
