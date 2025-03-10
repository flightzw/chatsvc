### Chatsvc - 聊天室项目服务端

![chatsvc](https://github.com/flightzw/chatsvc/actions/workflows/docker-ci.yml/badge.svg?branch=master&event=push) ![chatsvc](https://github.com/flightzw/chatsvc/actions/workflows/golang-lint.yml/badge.svg?branch=master&event=push)

项目结构：

```bash
.
├── api
│   └── chatsvc
│       ├── errno    # 错误码
│       ├── model    # 模型定义
│       └── v1       # api 定义
├── cmd
│   └── chatsvc      # 程序入口
├── configs          # 本地配置
├── internal
│   ├── biz          # 业务层
│   ├── conf         # 配置定义
│   ├── data         # 数据访问层
│   │   ├── model    # gorm/gen model
│   │   └── query    # gorm/gen query
│   ├── entity       # 工具 model
│   ├── enum         # 枚举
│   ├── middleware   # http 中间件
│   ├── server       # http 实例配置
│   ├── service      # 服务层，模型转换
│   ├── utils        # 工具函数
│   ├── vo           # view model
│   └── ws           # websocket 会话管理
│       ├── client   # 对外提供调用 api
│       └── server   # 会话管理模块
└── third_party      # 第三方 proto 依赖
    ├── buf
    │   └── validate
    ├── errors
    ├── google
    │   ├── api
    │   └── protobuf
    └── openapi
        └── v3
```



