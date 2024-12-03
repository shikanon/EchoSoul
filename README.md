# EchoSoul
EchoSoul（音心），AI音乐播客，代码版权属于哂码科技

目录：
```
EchoSoul
├── api
├── cmd
│   └── main.go
├── config
├── docs
├── handlers
├── internal
├── models
├── pkg
├── scripts
└── utils
```
功能包括：
- 通过手机号码进行用户登录
- 获取播客节目列表
- 根据播客节目获取音频列表
- 查询用户已订阅播客节目
- 查询用户自己创建的播客节目


## 部署

初始化：
```
// 安装swag
go get -u github.com/swaggo/swag/cmd/swag
// 引入依赖
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```
生成swag文档:
```
export SWAGGO_DEBUG_MODE=true
swag init --output ./swagger_docs  -d ./cmd 
```

部署MySQL：
```bash
docker run --name mysql-server -e MYSQL_ROOT_PASSWORD=your_password -p 13306:3306 -d mysql:5.7
```
创建数据库:
```bash
docker exec -it mysql-server mysql -p
```
sql:
```sql
CREATE DATABASE echosouldb;
CREATE USER 'podcast_user'@'%' IDENTIFIED BY 'your_password123';
GRANT ALL PRIVILEGES ON echosouldb.* TO 'podcast_user'@'%';
FLUSH PRIVILEGES;
```

部署Redis:
```bash
docker run -d --name redis -p 16379:6379 redis redis-server --requirepass "your_password123"
```