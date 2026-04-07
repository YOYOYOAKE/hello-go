package checker

import (
	"net/http"
	"service-checker/internal/model"
	"time"
)

func checkOne(t model.Target) model.CheckResult {
	// 1 创建自定义请求对象
	req, err := http.NewRequest("GET", t.URL, nil)
	if err != nil {
		return model.CheckResult{
			Name:       t.Name,
			OK:         false,
			DurationMS: -1,
			FailReason: err.Error(),
		}
	}

	// 2 创建自定义客户端
	cli := http.Client{Timeout: time.Duration(t.TimeoutMS) * time.Millisecond}

	// 3 开始计时并发送自定义请求
	start := time.Now()

	res, err := cli.Do(req)
	if err != nil {
		return model.CheckResult{
			Name:       t.Name,
			OK:         false,
			DurationMS: -1,
			FailReason: err.Error(),
		}
	}

	defer res.Body.Close()

	// 4 结束计时并计算时间
	d := time.Since(start).Milliseconds()

	return model.CheckResult{
		Name:       t.Name,
		OK:         res.StatusCode >= 200 && res.StatusCode < 300,
		DurationMS: d,
		FailReason: "",
	}
}
