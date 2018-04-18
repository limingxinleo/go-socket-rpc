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