package checker

import "service-checker/internal/model"

func CheckTargets(targets []model.Target) []model.CheckResult {
	// 1 创建结果容器和收集通道
	var Results []model.CheckResult
	resultChannel := make(chan model.CheckResult)

	// 2 并发启动检测协程
	for _, target := range targets {
		go func(t model.Target) {
			resultChannel <- checkOne(t)
		}(target)
	}

	// 3 阻塞运行直到所有收集齐所有结果
	for range targets {
		Results = append(Results, <-resultChannel)
	}

	return Results
}
