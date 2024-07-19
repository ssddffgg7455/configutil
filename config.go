package configutil

var config Config

// Config 配置参数接口
type Config interface {
	// GetConf 获取配置参数中的具体值 args 为 配置 keypath
	GetConf(args ...string) interface{}
	// GetAllConf 获取所有配置
	GetAllConf() map[string]interface{}
}

// Init 初始化
func Init(options ...Option) error {
	configFile := &ConfigFile{
		globalConfig: make(map[string]interface{}),
	}
	err := configFile.Init(options...)
	if err != nil {
		return err
	}
	config = configFile
	return nil
}

// Register 注入 Config
func Register(co Config) {
	config = co
}

// GetConf 获取配置参数中的具体值 args 为 配置 keypath
func GetConf(args ...string) interface{} {
	return config.GetConf(args...)
}

// GetAllConf 获取所有配置
func GetAllConf() map[string]interface{} {
	return config.GetAllConf()
}
