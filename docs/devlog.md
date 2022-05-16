# 开发日志

## 故障等级

共分为三个等级

严重：系统完全不可用，主要功能不可用，影响全部用户，需要立刻找人解决。

中等：系统部分不可用，次要功能不可用，影响部分用户，需要在1-2个工作日内找人解决。

一般：系统出现不易触发的BUG，影响少部分用户，正常排期解决。

## 详细记录

### 2021/09/13

故障等级：严重

故障修复：13:52-16:10

故障原因：后端 Gorm 连接 MySQL 数据库失败，未成功实例化 MySQL的情况下调用 Gorm 示例，导致出现空指针访问错误。

部分报错信息

```bash
runtime error: invalid memory address or nil pointer dereference

/backend/service/jwt_black_list.go:30 (0xd293af)
/backend/middleware/jwt.go:27 (0xd2e4c4)
```

### 2021/08/04

故障等级：中等

火花统一登录逻辑存在问题，修复时长 16:40 - 18:30

### 2021/04/26

VS Code 中 npm 安装依赖总是提示安装失败，可能是 VS Code 使用 node_modules。

关闭 VS Code，在 Terminal 中安装即可。

### 2021/04/01

Element UI v2.12 Popconfirm 组件在 table 组件中表现异常，按钮失效。

### 2021/03/26

Vue 或 Element UI 更新节点未全部更新，导致表单校验失败。

登录/注册功能两个表单混淆。

### 2021/03/25

gorm AutoMigrate 不支持创建数据库外键。
