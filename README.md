### Gin Web 骨架 v2


#### 开始
使用``go mod tiyd``下载依赖包


#### 数据库连接
1. 内置多种连接器 比如postgresql mysql 等
2. 修改Enum 中对应配置即可

### GORM

> 自动创建表 手动添加model模型 使用逗号分开即可
```go
_db.AutoMigrate(&Models.User{}) // 自动创建表结构
```

#### 目录结构
- bean 
- context  service上下文，自动注入服务
- handler  控制器handler
- databases 数据库连接器
- enums 枚举
- middlewares 中间件
- models gorm模型
- router 路由
- logic 业务逻辑层,实现具体业务
- utils 工具类

#### 开发规范


#### 开发流程


#### 交叉编译

- Mac 下编译 Linux 和 Windows 64位可执行程序

```cgo
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go //Linux
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go //windows
```

- Linux 下编译 Mac 和 Windows 64位可执行程序
```cgo
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

- Windows 下编译 Mac 和 Linux 64位可执行程序
```cgo
// mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

// linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go

// windows
go build main.go
```
