package util

import (
	log "github.com/cihub/seelog"
)

var logger log.LoggerInterface

// Log 获取日志
func Log() log.LoggerInterface {
	return logger
}

// InitialLog 初始化日志
func InitialLog(logPath string) error {
	var err error
	logger, err = log.LoggerFromConfigAsFile(FileUtil().AbsPath(logPath))
	if err != nil {
		return err
	}
	return nil
}
