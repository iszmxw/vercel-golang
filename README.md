# vercel-golang

- 开发 golang 服务并且部署到 vercel

[![Deploy with Vercel](https://vercel.com/button?utm_source=busiyi&utm_campaign=oss)](https://vercel.com/new/clone?utm_source=busiyi&utm_campaign=oss&repository-url=https://github.com/iszmxw/vercel-golang&env=CONFIG)


## 配置文件

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
    "DB_HOST": "db4free.net",
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
        "logo": "https://avatars.githubusercontent.com/u/135967790?s=400&u=0efb3cc947e9f0c2165c11f65f374524cb48915d&v=4"
    },
    "wechat": {
        "botUrl": "http://106.52.198.173:8000",
        "guid": "68d4ebef-f854-3387-ab40-7832d51dab25",
        "BelongWx": "xiaomg_zs"
    }
}

```
