# gapp
go gin框架应用脚手架

### gapp使用了那些第三方库
> [**框架: gin**](https://github.com/gin-gonic/gin)

> [**数据库: gorm**](https://github.com/jinzhu/gorm)

> [**.env配置: godotenv**](https://github.com/joho/godotenv)


### gapp目录结构
```php
gapp                  应用根目录
├─dev                  开发者存放项目相关开发信息
├─gapp                  应用根目录
│  ├─langs              核心语言包目录
│  ├─routers            路由目录
│  │  ├─api_router.go       接口路由配置文件
│  │  ├─app_router.go       网页应用路由配置文件
│  │  ├─router.go           路由配置文件
│  │  ├─api                 接口目录
│  │  │  ├─v1                   v1接口开发目录
│  │  │  │  ├─demo                示例接口目录
│  │  │  │  │  ├─simple.go          简单请求示例
│  │  │  │  │  ├─gorm.go            gorm操作数据库(db)示例
│  │  │  ├─ ...                 vn接口开发目录
│  │  ├─app                 网页应用目录
│  │  │  ├─v1                   v1网页应用开发目录
│  │  │  │  ├─demo                  示例网页应用开发目录
│  │  │  │  │  ├─simple.go      简单请求示例
│  │  │  ├─ ...                 vn网页应用开发目录
│  │  ├─Vendor                  第三方类库目录
│  ├─.env               配置文件
│  ├─example.env        示例配置文件
│  ├─LICENSE.txt        授权协议文件
│  ├─README.txt         README文件
│  └─main.go            入口文件
```