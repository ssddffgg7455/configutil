# 读取toml配置

## 初始化
```
err := configutil.Init()
if err != nil {
    fmt.Printf("configutil init err [%v] \n", err)
}
```

### 默认读取以下配置文件
```
./config.toml
../config.toml
./config/config.toml
../config/config.toml
```

### 传入路径方法如下
```
err := configutil.Init(configutil.WithFilePath(url))
if err != nil {
    fmt.Printf("configutil init err [%v] \n", err)
}
```

## 读取配置
```
// [config]
// app = 1

// 读取某项配置
configutil.GetConf("config","app") // return interface{}
// 读取全部配置
configutil.GetAllConf() // interface{} // return map[string]interface{}
```
