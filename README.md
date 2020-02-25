goconf
========

## 描述

使用 `goconf` 更简单的读取 `go` 的 `ini` 配置文件以及根据特定格式的各种配置文件。

## 安装方法

```bash
go get github.com/auroralzdf/goconf
```

## 使用方法

### ini配置文件格式样列
```bash
[database]
hostname = localhost
username = root
password = 123456
port     = 3306

[redis]
username = root
password = 123456
port     = 6379

```
### 初始化
```go
conf := goconf.InitConfig("./conf/conf.ini") // "./conf/conf.ini" 是你配置文件的位置
```

### 获取单个配置信息
```go
username := conf.GetValue("database", "username") // database 是你的 [section]， username 是你要获取值的 key 名称
fmt.Println(username) // root
```

### 删除一个配置信息
```go
conf.DeleteValue("database", "username")	//username 是你删除的 key
username = conf.GetValue("database", "username")
if len(username) == 0 {
    fmt.Println("username is not exists") // this stdout username is not exists
}
```

### 添加一个配置信息
```go
conf.SetValue("database", "username", "root")
username = conf.GetValue("database", "username")
fmt.Println(username) //root 添加配置信息如果存在 [section] 则添加或者修改对应的值，如果不存在则添加 section
```

### 获取所有配置信息
```go
conf.GetAllSetion() //返回 map[string]map[string]string 的格式 即 setion => key -> value
```

---

# example
=========
```go
package main

import (
	"fmt"
	"github.com/auroraLZDF/goconf"
)

func main() {

	conf := goconf.InitConfig("./config.ini")

	for key,value :=range conf.Conflist {
		fmt.Println(key)
		for k,v := range value{
			fmt.Println(k,":",v)
		}
	}
	
	fmt.Println()
	
	fmt.Println(conf.GetValue("database","hostname"))

	conf.SetValue("database","hostname","127.100.100.100")

	fmt.Println(conf.GetValue("database","hostname"))
}
```
>output
```
database
hostname : localhost
username : root
password : 123456
port     : 3306

redis
username : root
password : 123456
port     : 6379

localhost
127.100.100.100

Process finished with exit code 0
```


