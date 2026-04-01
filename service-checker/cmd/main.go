package main

import (
	"flag"
	"fmt"
	"service-checker/internal/checker"
	"service-checker/internal/config"
	"service-checker/internal/model"
	"service-checker/internal/store"
)

func main() {
	// 1 从命令行参数中读取配置路径
	filepath := flag.String("c", "config/config.json", "Your Config File Path")
	outputpath := flag.String("o", "output/output.txt", "Output File Path")

	flag.Parse()

	// 2 解析配置文件
	targets, err := config.ParseConfig(*filepath)
	if err != nil {
		fmt.Println("Failed to Load Config:", err)
		return
	}
	fmt.Println("Successed to Load Config:", *filepath)

	// 3 定时启动检查进程
	checkResult := checker.CheckTargets(targets)

	// // 4 保存检查结果（以接口的方式遍历打印）
	consolePrinter := store.ConsolePrintr{}
	filePrinter := store.FilePrinter{Path: *outputpath}

	resultStore := []model.ResultStore{}
	resultStore = append(resultStore, consolePrinter)
	resultStore = append(resultStore, filePrinter)

	for _, s := range resultStore {
		err := s.WriteResult(checkResult)
		if err != nil {
			fmt.Println("Failed to Save Check Result:", err)
			return
		}
	}
}
