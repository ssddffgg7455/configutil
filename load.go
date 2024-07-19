package configutil

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// ConfigFile 本地文件配置
type ConfigFile struct {
	globalConfig map[string]interface{}
}

// Init 初始化
func (c *ConfigFile) Init(options ...Option) error {
	fmt.Println("Loading config ...")
	configData := new(ConfigData)
	for _, option := range options {
		option(configData)
	}

	if len(configData.filePath) == 0 {
		configData.filePath = getCfgPath()
	}

	if len(configData.filePath) == 0 {
		return fmt.Errorf("config file not exist")
	}

	fmt.Println("cfgPath: " + configData.filePath)

	_, err := toml.DecodeFile(configData.filePath, &c.globalConfig)
	if err != nil {
		return err
	}
	fmt.Println("Loading config complete.")
	return nil
}

// GetConf 获取配置参数中的具体值
func (c *ConfigFile) GetConf(args ...string) interface{} {
	return _getConf(c.globalConfig, args)
}

func (c *ConfigFile) GetAllConf() map[string]interface{} {
	return c.globalConfig
}

// 注意 config.toml文件只能生效一个 且存在优先级
func getCfgPath() string {
	var err error

	_, err = os.Stat("./config.toml")
	if err == nil {
		return "./config.toml"
	}
	_, err = os.Stat("../config.toml")
	if err == nil {
		return "../config.toml"
	}
	_, err = os.Stat("./config/config.toml")
	if err == nil {
		return "./config/config.toml"
	}
	_, err = os.Stat("../config/config.toml")
	if err == nil {
		return "../config/config.toml"
	}

	return ""
}

func _getConf(cfg map[string]interface{}, args []string) interface{} {
	if len(args) == 0 {
		return nil
	}
	switch v := (cfg[args[0]]).(type) {
	case map[string]interface{}:
		if len(args) > 1 {
			return _getConf(v, args[1:])
		}
		return v
	default:
		return v
	}
}
