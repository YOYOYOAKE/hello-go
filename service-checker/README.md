# Service Checker

该项目用于个人练习，用于巩固 Go 核心内容。

本项目涉及到的内容有：

- 结构体与切片：管理配置和状态
- goroutine 与 channel：汇总结果
- flag：命令行参数
- json：处理配置文件
- net/http：HTTP 访问
- defer：资源释放

## 第一阶段 最小可用实现

程序启动后读取 `config.json`，结构类似于：

```json
[
    {
        "name": "github",
        "url": "https://github.com",
        "timeout_ms": 3000
    },
    {
        "name": "local-api",
        "url": "http://127.0.0.1:8080/health",
        "timeout_ms": 1000
    }
]
```

程序并发检查所有目标，把检测结果打印成表格。

```text
Name 	 Available 	 Duration(ms) 
local-api 	 false 	 -1 
github 	 true 	 335 
```

项目结构：

```text
service-checker/
├─ cmd/main.go 程序入口
├─ internal/config/ 配置文件处理
├─ internal/checker/ 服务状态检查
├─ internal/store/ 结果存储
├─ internal/model/ 数据模型定义
├─ config/config.json 配置文件
└─ README.md
```