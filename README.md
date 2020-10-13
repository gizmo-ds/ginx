# ginx

一个基于 gin 的简单中间件, 能更方便地处理返回数据跟处理数据绑定.

> 没做过性能测试, 可能会存在性能上的问题.

[![Build Status](https://drone.liuli.lol/api/badges/GizmoOAO/ginx/status.svg?ref=refs/heads/main)](https://drone.liuli.lol/GizmoOAO/ginx)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/GizmoOAO/ginx/main)
[![GitHub](https://img.shields.io/github/license/GizmoOAO/ginx)](./LICENSE)

# 安装

下载并安装 ginx:

```bash
go get -u github.com/GizmoOAO/ginx
```

将 ginx 引入到代码:

```go
import "github.com/GizmoOAO/ginx"
```

# 使用

```go
app := gin.New()
app.Use(ginx.Ginx())
```

# 许可证

MIT
