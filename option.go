package configutil

// ConfigData 配置数据
type ConfigData struct {
	filePath string
}

type Option func(*ConfigData)

// WithFilePath 设置文件路径
func WithFilePath(filePath string) Option {
	return func(cd *ConfigData) {
		cd.filePath = filePath
	}
}
