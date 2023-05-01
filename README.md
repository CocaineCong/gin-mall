# gin-mall

**基于 gin+gorm+mysql读写分离 的一个电子商场**

本项目改自于作者[Congz](https://github.com/congz666)的[电子商城](https://github.com/congz666/cmall-go)
去除了一些如第三方登录，极验，第三方支付等功能，新增了MySQL读写分离、ELK日志体系、AES对称加密进行数据脱敏等。
在此也非常感谢作者开源！

此项目比较全面，比较适合小白入门`web开发`

# 更新说明
**V2版本，结构较比V1版本有很大的改动**
**全部转化成 controller、dao、service 模式，更加符合企业开发**

考虑到部分同学的基础，所以V2版本的技术栈只有mysql，redis，docker。

而 ELK，MQ，Jaeger，Prometheus 这部分都集成在V3版本，并且V3版本的项目结构进行部分重构。

由于整合上传oss和上传到本地，需要在 conf 中进行配置 `UploadModel` 字段，上传到 oss 则配置 oss，上传本地则配置 local

其中我个人用到的测试sql数据都放在了`doc/mall_sql`文件当中

# 开源合作
欢迎大家把自己的想法 pr 到这个项目中。

**说明：**
1. 大家可以根据自己的需要进行分支的合并，不要直接合main分支⚠️，尽量合去最新的版本。现在最新版本是v3版本。
2. CR 通过之后，就会到合并到 main 分支。

# 项目的主要功能介绍

- 用户注册登录(JWT-Go鉴权)
- 用户基本信息修改，解绑定邮箱，修改密码
- 商品的发布，浏览等
- 购物车的加入，删除，浏览等
- 订单的创建，删除，支付等
- 地址的增加，删除，修改等
- 各个商品的浏览次数，以及部分种类商品的排行
- 设置了支付密码，对用户的金额进行了对称加密
- 支持事务，支付过程发送错误进行回退处理
- 可以将图片上传到对象存储，也可以切换分支上传到本地static目录下
- 添加ELK体系，方便日志查看和管理

# 项目需要完善的地方

## P0
-[ ] 考虑加入kafka或是rabbitmq，新增一个秒杀专场 \
-[x] 优化 service 返回的参数，加上返回值 error，因为go的函数返回都是要有error的，这才是go的代码风格（我也不懂go为啥要这样设置，很多优秀的开源项目都是这样写函数的返回值） \
-[x] 抽离 service 的结构体到 types，引入 sync.Once 模块，重构 service 层 \
-[ ] 优化鉴权模块，加上 refreshToken，将 token 改成 accessToken \
-[ ] 抽离登陆，引入SSO\
-[x] 优化日志输出，统一用日志对象 \
-[x] 考虑 cmd 和 loading 这两个文件夹是否合并\
-[ ] 加入 Jaeger 进行链路追踪\
-[ ] 加入 Prometheus 监控中间件\
-[ ] 优化ToC应用的 SQL JOIN 语句\
-[ ] MySQL到ES的数据同步，将搜索改成查找ES（注意一下，这里最好引入kafka，mysql推到kafka，kafka再推到es，确保一下ack）


# 项目的主要依赖：
Golang V1.18
- gin
- gorm
- mysql
- redis
- ini
- jwt-go
- crypto
- logrus
- qiniu-go-sdk
- dbresolver

# 项目结构
```
gin-mall/
├── api
├── cmd
├── conf
├── doc
├── middleware
├── model
├── pkg
│  ├── e
│  └── util
├── repository
│  ├── cache
│  ├── db
│  ├── es
│  ├── mq
│  └── redis
├── routes
├── serializer
├── service
└── static
```
- api : 用于定义接口函数，也就是controller的作用
- conf : 用于存储配置文件
- dao : 对持久层进行操作
- doc : 存放接口文档
- loading : 需要加载的应用
- middleware : 应用中间件
- model : 应用数据库模型
- pkg/e : 封装错误码
- pkg/util : 工具函数
- repository : 存放存储仓库
- repository/cache : 放置redis缓存
- repository/db : 放置持久层的mysql
- repository/db/dao : dao层，对db进行操作
- repository/db/model : 定义mysql的模型
- repository/es : 放置es，形成elk体系
- repository/mq : 放置各种mq，kafka，rabbitmq等等...
- routes : 路由逻辑处理
- serializer : 将数据序列化为 json 的函数，便于返回给前端
- service : 接口函数的实现
- static : 存放静态文件

# 配置文件
`conf/config.ini` 文件配置

```ini
#debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000

[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = root
DbName =

[redis]
RedisDb = redis
RedisAddr = 127.0.0.1:6379
RedisPw =
RedisDbName =

[qiniu]
AccessKey =
SerectKey =
Bucket =
QiniuServer =

[email]
ValidEmail=http://localhost:8080/#/vaild/email/
SmtpHost=smtp.qq.com
SmtpEmail=
SmtpPass=
#SMTP服务的通行证

[es]
EsHost = 127.0.0.1
EsPort = 9200
EsIndex = mylog
```

## 简要说明
1. `mysql` 是存储主要数据。
2. `redis` 用来存储商品的浏览次数。
3. 由于使用的是AES对称加密算法，这个算法并不保存在数据库或是文件中，是第一次登录的时候需要给的值，因为第一次登录系统会送1w作为初始金额进行购物，所以对其的加密，后续支付必须要再次输入，否则无法进行购物。
4. 本项目运用了gorm的读写分离，所以要保证mysql的数据一致性。
5. 引入了ELK体系，可以通过docker-compose全部up起来，也可以本地跑(确保ES和Kibana都开启)
6. 用户创建默认金额为 **1w** ，默认头像为 `static/imgs/avatar/avatar.JPG` 
# 导入接口文档

打开postman，点击导入

![postman导入](doc/1.点击import导入.png)

选择导入文件
![选择导入接口文件](doc/2.选择文件.png)

![导入](doc/3.导入.png)

效果

![展示](doc/4.效果.png)


这里是用postman查询es，Kibana也可以查看es！

![postman-es](doc/5.postman-es.png)

# 项目运行
**本项目采用Go Mod管理依赖**

**下载依赖**

```shell
go mod tidy
```

**下载依赖**
```shell
go run main.go
```
