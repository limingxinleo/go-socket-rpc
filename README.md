# GoRPC


## 安装
~~~
go get github.com/limingxinleo/gorpc
~~~

## 使用
~~~go
package main

import (
	"github.com/limingxinleo/gorpc/app"
	"github.com/limingxinleo/gorpc/tests/handlers"
)

func main() {
	rpc := &app.Rpc{Address: "0.0.0.0:11521"}
	rpc.RegisterHandler("test", &handlers.Test{})
	rpc.Run()
}
~~~

## 测试
对应的PHP单元测试 https://github.com/limingxinleo/gorpc-unit-test

测试代码位于tests中
ORM使用：https://github.com/jinzhu/gorm
Mysql使用：https://github.com/go-sql-driver/mysql