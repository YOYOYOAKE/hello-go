# hello-go

这是我的 Go 学习仓库，用来记录练习项目、代码实验和阶段性成果。

我把不同主题的小项目放在各自的子目录中。每个项目通常会包含自己的 `go.mod`、README 和源代码。

## 项目列表

- `service-checker`
  
  一个简单的服务可用性检查工具。读取 JSON 配置，并发请求多个 URL，输出检查结果到控制台和文件。涉及到的内容有：

  - 结构体与切片：管理配置和状态
  - goroutine 与 channel：汇总结果
  - flag：命令行参数
  - json：处理配置文件
  - net/http：HTTP 访问
  - defer：资源释放
