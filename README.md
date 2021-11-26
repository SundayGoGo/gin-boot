### Gin Web 骨架 v2




#### 数据库连接
1. 内置多种连接器 比如postgresql mysql 等
2. 修改Enum 中对应配置即可

### GORM

> 自动创建表 手动添加model模型 使用逗号分开即可
```golang
	_db.AutoMigrate(&Models.User{}) // 自动创建表结构
```

#### 目录结构
- bean 
- context  mapper上下文，自动注入mapper接口，统一交给service调用
- controller  控制器handler
- databases 数据库连接器
- enums 枚举
- mapper 实现与数据库交互的接口
- middlewares 中间件
- models gorm模型
- router 路由
- services 业务逻辑层 调用mapper接口，实现具体业务
- utils 工具类

#### 开发规范


#### 开发流程
