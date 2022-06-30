## 说明
基于 Gin 进行模块化设计的 API 框架，封装了常用功能，操作简单，致力于进行快速的业务研发。

## 技术选型
- 框架：gin
- 数据库：gorm.io/driver/mysql
- 模型：gorm.io/gorm
- 配置：github.com/spf13/viper
- 校检：github.com/go-playground/validator/v10 已支持自定义错误内容
- 验证：github.com/golang-jwt/jwt/v4（原https://github.com/dgrijalva/jwt-go）

## 目录结构

```
├── App 项目核心目录
│   └── Api api接口目录
│       ├── Middlewares 中间件
│       │   ├── Auth 鉴权
│       │   │   └── Auth.go
│       │   ├── Corss
│       │   │   └── Corss.go
│       │   └── Recover
│       │       └── Recover.go
│       └── Router 路由
│           └── router.go
├── Config 配置
│   ├── config.go
│   └── mysql.go
├── Cores 核心库
│   ├── cores.go 初始化核心库
│   ├── mysql 数据库初始化
│   │   └── mysql.go
│   └── viper viper初始化
│       └── viper.go
├── README.md
├── config.yaml 配置文件
├── go.mod
├── go.sum
└── main.go 入口文件
```
