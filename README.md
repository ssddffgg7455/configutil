### 读取配置

err := configutil.Init()
if err != nil {
    fmt.Printf("configutil init err [%v] \n", err)
}

默认读取
./config.toml
../config.toml
./config/config.toml
../config/config.toml
以下配置文件

传入路径
err := configutil.Init(configutil.WithFilePath(url))
if err != nil {
    fmt.Printf("configutil init err [%v] \n", err)
}
