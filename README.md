# zlog
a golang log package base on https://github.com/marmotedu/gopractise-demo/tree/main/log/cuslog
## 功能特性
- [x] debug level
- [x] debug path
- [x] debug filename
- [x] format
- [x] facade
- [x] division
- [x] clean

## 快速开始

```go
go get -u github.com/JoyZF/zlog
```

```go
package main

import "github.com/JoyZF/zlog"

func main() {
    New(WithServiceName("test"),
        WithStdLevel(DebugLevel),
        WithFormatter(&JsonFormatter{}),
        WithOutputPath("./logs/", "app.log"),
        WithCleaner(&Clean{
            Interval: 24 * time.Hour,
            Reserve:  7 * 24 * time.Hour,
    }))
    Error("this is a error log")
}
```

