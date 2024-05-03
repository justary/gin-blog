package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {

	vp := viper.New()

	// 设置配置文件名和路径
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("configs/")

	// 读取配置文件
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp: vp}, nil
}
