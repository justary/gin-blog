package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	Sync         bool
	TablePrefix  string
	ParseTime    bool
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

func (s *Setting) ReadSection(key string, v interface{}) error {

	if err := s.vp.UnmarshalKey(key, v); err != nil {
		return err
	}

	return nil
}
