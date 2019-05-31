<h1>GOAPP</h1>

<div>
 基于 Gin + GORM + Casbin + vue-element-admin 实现的权限管理系统 <br/>
 基于Casbin 实现RBAC权限管理 <br/>
 前端实现： vue-element-admin <br/>
</div>
<br/>

## 特性

- 基于 Casbin 的 RBAC 访问控制模型
- JWT 认证
- 前后端分离

## 下载并运行

### 获取代码

```
go get -v github.com/it234/goapp
```

### 运行

- 可以直接下载打包好的桌面客户端体验，下载地址: https://pan.baidu.com/s/1wDsHH-KMQHV5tMRUv50Q3w 提取码: 9u2d 
- 运行服务端：cd cmd/manageweb，go run main.go，运行成功后打开 127.0.0.1:8080
- 调试/运行web：cd website/manageweb，npm install，npm run dev
- 配置文件在(`cmd/manageweb/config.yaml`)中，用户默认为：admin/123456


#### 温馨提醒

1. 默认配置采用的是 sqlite 数据库，数据库文件(`自动生成`)在`cmd/manageweb/data/goapp.db`。如果想切换为`mysql`或`postgres`，请更改配置文件，并创建数据库（表会自动创建）。
2. 日志的配置为标准输出并写入文件。

## 前端实现

- website/manageweb：基于[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)的实现版本

## 项目结构概览

<details>
<summary>展开查看</summary>
<pre><code>.
├── cmd  项目的主要应用
├── internal  私有应用程序和库代码
├── pkg  外部应用程序可以使用的库代码
├── vendor  项目依赖的其他第三方库
├── website  vue-element-admin
</code></pre>
</details>


## 界面截图

<details>
<summary>展开查看</summary>
<pre><code>.
<img src="./login.jpeg" width="900" height="600" /><br/><br/>
<img src="./icon.jpeg" width="900" height="900" /><br/><br/>
<img src="./menu.jpeg" width="900" height="900" /><br/><br/>
<img src="./role.jpeg" width="900" height="900" /><br/><br/>
<img src="./admins.jpeg" width="900" height="900" /><br/><br/>
</code></pre>
</details>


## 感谢以下框架的开源支持

- [Gin] - [https://gin-gonic.com/](https://gin-gonic.com/)
- [GORM] - [http://gorm.io/](http://gorm.io/)
- [Casbin] - [https://casbin.org/](https://casbin.org/)
- [vue-element-admin] - [https://github.com/PanJiaChen/vue-element-admin/](https://panjiachen.github.io/vue-element-admin/)


## MIT License

    Copyright (c) 2019 Lyric

## 与作者对话

> 该项目是利用业余时间进行开发的，主要是对golang+vue-element-admin进行一个简单的实现，如果您有更好的想法和建议请与我进行沟通，我非常期待！我的微信号：it23456789，微信二维码：

<img src="./wechat.jpeg" width="256" height="256" />




