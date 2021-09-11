# 通用后端Go框架

## 一、目录组织结构
    ├─apps             // 业务应用
    │  ├─comm          // 多个相似业务公用模块
    │  │  ├─actions    // 公用控制器
    │  │  ├─defines    // 公用接口，类型，常量等定义
    │  │  ├─models     // 公用model，数据抽象
    │  │  ├─router     // 公共路由
    │  │  └─services   // 公用service对数据库或其他数据源操作的包装
    │  ├─routers       // 应用路由包装  
    │  └─demo          // 演示应用
    │      ├─actions   // 应用控制器
    │      ├─defines   // 应用接口，类型，常量等定义
    │      ├─models    // 应用model，数据抽象
    │      ├─router    // 应用路由
    │      └─services  // 应用service，对数据库或其他数据源操作的包装
    ├─config           // 服务配置管理
    ├─docs             // 工程文档 
    ├─errs             // 错误抽象管理
    ├─mwares           // 中间件
    ├─resp             // 返回结构抽象
    ├─scripts          // 脚本目录
    │  ├─docker        // docker脚本目录
    │  └─swagger       // swagger脚本目录
    ├─server           // gin server和gorm连接包装
    ├─statics          // 静态资源目录(前端)
    └─utils            // 工具集
        ├─crypto       // 加密和Hash相关函数包装
        ├─file         // 文件和目录相关函数包装
        ├─helper       // 通用助手函数
        └─logger       // 日志函数包装

## 二、swagger文件生成方式
**注意: 以下两种方式在同一个工程内不要混用！！！**

### 2.1 接口注释方式

在按照swagger规范编写完接口注释后，在终端执行如下命令：

`$ scripts/swagger/gen`

### 2.2 通过工具生成(推荐使用Stoplight Studio)

通过工具编辑好swagger文档后将其保存为docs/swagger.json，然后在终端执行如下命令：

`$ scripts/swagger/renew`

### 2.3 代码自动生成
在scripts/coder目录包含代码自动生成工具，主要用来生成和初始化应用和模块

生成代码：
`$  scripts/swagger/gen APP_NAME MODULE_NAME`

  - APP_NAME    - 是apps下单个应用的根目录
  - MODULE_NAME - 模块名称，会自动据此生成actions、services、models、routers下对应的文件和基础代码。必须是全小写加"_"（package name）。
  
删除代码（注意是物理删除）：
`$  scripts/swagger/del APP_NAME MODULE_NAME`
## 三、注意事项

 - 所有的model字段需要显式指定json key
 - 相似度或者可重用性很高的应用可考虑使用同工程多个app实现，而不是多个branch并部署多个服务


## 四、改进计划

### 4.1 完善中间件及集成其他库
增加Redis，GRPC，WebSocket，GraphQL支持

### 五、swagger接口
http://localhost:8999/swagger/index.html