# 在线练习/考试系统

> 基于 [gin-vue-admin](https://github.com/piexlmax/gin-vue-admin)

在线练习/考试系统是一个为用户提供在线做题的平台。学生可用于备考及考试，教师可进行试题管理及学生成绩分析。

## 快速开始

> 参考 [gin-vue-admin](https://github.com/piexlmax/gin-vue-admin)

`git clone git_url `

替换`git_url`为项目实际地址

### 前端

**Requirements：**`NodeJS`

```bash
$ cd web
$ npm install
$ npm run serve # for development
$ npm run build # for deployment
```

### 后端

**Requirements：**`Golang`

```bash
$ cd server
$ go generate
$ go run main.go
```

#### API文档

```bash
$ go get -u github.com/swaggo/swag/cmd/swag
$ cd server
$ swag init
```

## 技术架构

- 前端

  Vue、ElementUI

- 后端

  Golang

- 数据库

  MySQL
  MongoDB
  Redis

