<h1>
<a href="https://www.dtapp.net/">Golang Decimal</a>
</h1>

📦 Golang 小数点处理

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/godecimal?status.svg)](https://pkg.go.dev/github.com/dtapps/godecimal)
[![goproxy.cn](https://goproxy.cn/stats/github.com/dtapps/godecimal/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/godecimal)
[![goreportcard.com](https://goreportcard.com/badge/github.com/dtapps/godecimal)](https://goreportcard.com/report/github.com/dtapps/godecimal)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgodecimal)

#### 安装

```go
go get -v -u github.com/dtapps/godecimal
```

#### 使用

```go
package main

import (
	"github.com/dtapps/godecimal"
	"log"
	"reflect"
)

func main() {
	log.Println("加：", godecimal.Float64Add(10, 3), reflect.TypeOf(godecimal.Float64Add(10, 3)))
	log.Println("减", godecimal.Float64Sub(10, 3), reflect.TypeOf(godecimal.Float64Sub(10, 3)))
	log.Println("乘：", godecimal.Float64Mul(10, 3), reflect.TypeOf(godecimal.Float64Mul(10, 3)))
	log.Println("除：", godecimal.Float64Quo(10, 3), reflect.TypeOf(godecimal.Float64Quo(10, 3)))
}
```