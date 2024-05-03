package global

import (
	"github.com/justary/gin-blog/pkg/logger"
	"github.com/justary/gin-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
