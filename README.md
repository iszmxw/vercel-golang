# vercel-golang

- 开发 golang 服务并且部署到 vercel
- 请复制底部的 `json` 文件修改后开始部署体验

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https://github.com/iszmxw/vercel-golang&project-name=vercel-golang&repository-name=vercel-golang&env=CONFIG)


## 配置文件

- [application.yaml](application.yaml) 为项目的主要配置文件，当不依赖 `vercel` 的时候可以使用此配置文件加载配置

```bash
# 进入项目根目录
# 修改 application.yaml
# 数据库相关信息修改
# 很重要的一步
echo "DEV=1" > .env
echo "TLS=false" > .env
# 有上面的配置才会走 application.yaml
# 修改好后就可以通过根目录 main.go 文件来启动开发
go run main.go
```

- [config.yml](config.yml) 为 【[gormt.go](bin%2Fgormt.go)】 的配置文件 【主要用来管理 orm 的，自动维护 gorm 模型 使用方法如下】

```shell
# 进入项目根目录
go run bin/gormt.go
# 运行完后会根据链接的数据表刷新 ./app/models 中的 model 文件
```

- [.env.example](.env.example) 为 `vercel` 开发模式下的配置文件通过环境变量的方式加载引入

```bash
# 进入项目根目录
cp .env.example .env
# 修改 TLS true/false
# TLS=false
# 修改 CONFIG
# CONFIG={"APP_NAME":"FundStock","APP_ENV":"dev","APP_URL":"http://127.0.0.1","APP_PORT":8888,"APP_PPROF":true,"HTTPS":0,"ADDRESS_LIMIT":true,"DB_CONNECTION":"mysql","DB_HOST":"localhost","DB_PORT":3306,"DB_DATABASE":"maccms","DB_USERNAME":"root","DB_PASSWORD":"123456","DB_PREFIX":"t_","REDIS_HOST":"127.0.0.1","REDIS_PASSWORD":"","REDIS_PORT":6379,"REDIS_DB":0,"PW_SALT":"xw2023-06-08","LOG_MYSQL_DEBUG":true,"LOG_MYSQL_ERROR":true,"LOG_MYSQL_WARN":true,"NOTIFY_EMAIL":"mail@54zm.com","bark":{"url":"https://api.day.app/","key":{"mac":"YvemvqzPAgvc2ASqwcP2KH","iphone":"nPuFoXfYsMpVrUaqU8to99"},"logo":"https://avatars.githubusercontent.com/u/135967790?s=400&u=0efb3cc947e9f0c2165c11f65f374524cb48915d&v=4"},"wechat":{"botUrl":"http://106.52.198.173:8000","guid":"68d4ebef-f854-3387-ab40-7832d51dab25","BelongWx":"xiaomg_zs"}}
# 修改好后就可以本地开发了
vercel dev
```

## 本地开发

> 需要安装 [Golang](https://go.dev/dl/)（建议安装 `go1.18.10` 及以上版本）及 [pnpm](https://pnpm.io/zh/installation)

首先把 `vercel-golang` 仓库 `fork` 一份到自己的 `Github`，然后从个人仓库把项目 `clone` 到本地，项目默认是 `main` 分支。

然后依次在项目根目录运行以下命令：

```bash
# 安装依赖
go mod tidy
# 安装 vercel 方便 vercel本地开发调试
pnpm i -g vercel
# 登录 vercel ,按照提示进行登录
vercel login
```

运行完上述命令后，环境已经准备好，此时可以新拉一条分支进行开发。


## vercel 上配置方法
- 配置文件文 `json` 格式，`vercel` 后台设置环境变量 `CONFIG` 字段，字段内容如下
- 请自行替换成自己的配置后再进行使用

```json

{
    "APP_NAME": "GoApi",
    "APP_ENV": "dev",
    "APP_URL": "http://127.0.0.1",
    "APP_PORT": 80,
    "APP_PPROF": true,
    "HTTPS": 0,
    "ADDRESS_LIMIT": true,
    "DB_CONNECTION": "mysql",
    "DB_HOST": "127.0.0.1",
    "DB_PORT": 3306,
    "DB_DATABASE": "vercel_golang",
    "DB_USERNAME": "vercel_golang",
    "DB_PASSWORD": "123456",
    "DB_PREFIX": "t_",
    "REDIS_HOST": "ethical-walrus-49379.kv.vercel-storage.com",
    "REDIS_PASSWORD": "9f9659as78sdf90as78qw7fsd45cfa3b918",
    "REDIS_PORT": 6379,
    "REDIS_DB": 0,
    "PW_SALT": "xw2023-06-08",
    "LOG_MYSQL_DEBUG": true,
    "LOG_MYSQL_ERROR": true,
    "LOG_MYSQL_WARN": true,
    "NOTIFY_EMAIL": "mail@54zm.com",
    "bark": {
        "url": "https://api.day.app/",
        "key": {
            "mac": "YaemvezPAqvc2ASewcP2KH",
            "iphone": "nPuDoXfXsMpVrDaqU9to88"
        },
        "logo": "https://avatars.githubusercontent.com/u/31272102"
    },
    "wechat": {
        "botUrl": "http://127.0.0.1:8000",
        "guid": "68d4ebef-f854-3387-ab40-7832d51dab25",
        "BelongWx": "xiaomg_zs"
    }
}

```
