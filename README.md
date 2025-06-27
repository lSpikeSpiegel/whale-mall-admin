# mall-admin

## 项目结构

ecommerce/
├── cmd/                # 主程序入口
│   └── main.go
├── config/             # 配置加载（YAML/env）
├── controller/         # 控制器层，处理 HTTP 请求
├── middleware/         # 自定义中间件
├── model/              # 数据模型（GORM struct）
├── repository/         # 数据访问层
├── router/             # 路由定义
├── service/            # 业务逻辑处理
├── utils/              # 工具函数、通用组件
├── docs/               # 接口文档（Swagger）
├── go.mod
└── README.md

## 功能模块划分

用户模块（注册/登录/鉴权）

商品模块（增删改查）

分类模块（树状结构）

购物车模块（添加/删除商品）

订单模块（生成订单、支付逻辑）

支付模块

库存模块（扣减、还原）

后台管理模块（商品上架、审核等）

## 技术栈

框架：Gin

ORM：GORM

数据库：MySQL

缓存：Redis

配置管理：viper

JWT 鉴权：github.com/dgrijalva/jwt-go

日志：zap 或 logrus

接口文档：Swaggo (Swagger UI)

依赖管理：Go Modules
