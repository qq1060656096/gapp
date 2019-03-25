# gapp
go gin框架应用脚手架, 帮助你快速搭建golang项目


## gapp使用说明
> 1. gapp使用了那些golang第三方库
> 2. gapp目录结构
> 3. 安装gapp
> 4. gapp配置
> 5. 运行gapp
> 6. 示例文件

### 1. gapp使用了那些第三方库
> [**框架: gin**](https://github.com/gin-gonic/gin) -> [文档](https://gin-gonic.com/zh-cn/docs/)

> [**数据库: gorm**](https://github.com/jinzhu/gorm) -> [文档](http://gorm.io/zh_CN/docs/)

> [**.env配置: godotenv**](https://github.com/joho/godotenv) -> [文档](https://github.com/joho/godotenv)


### 2. gapp目录结构
```php
gapp                  应用根目录
├─dev                  开发者存放项目相关开发信息
├─gapp                  应用根目录
│  ├─langs              核心语言包目录
│  ├─models             模型目录
│  ├─pkg                公共库目录
│  │  ├─util                通用包目录
│  ├─resources          资源目录
│  │  ├─themes              主题(themes)目录
│  │  │  ├─default              默认主题(default)目录
│  │  │  ├─default_mobile       默认手机主题(default_mobile)目录
│  ├─routers            路由目录
│  │  ├─api_router.go       接口路由配置文件
│  │  ├─app_router.go       网页应用路由配置文件
│  │  ├─router.go           路由配置文件
│  │  ├─api                 接口目录
│  │  │  ├─v1                   v1接口开发目录
│  │  │  │  ├─demo                示例接口目录
│  │  │  │  │  ├─simple.go          简单请求示例
│  │  │  │  │  ├─gorm.go            gorm操作数据库(db)示例
│  │  │  │  │  ├─gorm_raw_sql.go    gorm raw sql 原生sql操作数据库(db)示例
│  │  │  ├─ ...                 vn接口开发目录
│  │  ├─app                 网页应用目录
│  │  │  ├─v1                   v1网页应用开发目录
│  │  │  │  ├─demo                  示例网页应用开发目录
│  │  │  │  │  ├─simple_html.go         简单网页应用示例
│  │  │  ├─ ...                 vn网页应用开发目录
│  │  ├─Vendor                  第三方类库目录
│  ├─.env               配置文件
│  ├─example.env        示例配置文件
│  ├─LICENSE.txt        授权协议文件
│  ├─README.txt         README文件
│  └─main.go            入口文件
```

### 3. 安装gapp
> 1. 查看GOPATH路径: go env 命令查看 GOPATH 路径
> 2. 下载gapp: git clone https://github.com/qq1060656096/gapp.git

### 4. 配置文件
> 1. 进入目录: cd gapp
> 2. 创建配置: cp example.env .env 
> 3. 配置数据

### 5. 运行gapp
```sh
# 运行gapp
go run main.go
```

### 6. gapp示例

```
# gorm model操作数据库
gapp/routers/api/v1/demo/gorm.go

# gorm 执行原生sql
gapp/routers/api/v1/demo/gorm_raw_sql.go

# 接口演示get post put delete
gapp/routers/api/v1/demo/simple.go

# 网页模板示例
gapp/routers/app/v1/demo/simple_html.go
```