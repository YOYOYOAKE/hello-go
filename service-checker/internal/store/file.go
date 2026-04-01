package store

import (
	"fmt"
	"os"
	"path/filepath"
	"service-checker/internal/model"
)

type FilePrinter struct {
	Path string
}

func (f FilePrinter) WriteResult(checkResult []model.CheckResult) {
	// 1 创建本地文件及目录
	os.MkdirAll(filepath.Dir(f.Path), 0755)
	file, err := os.Create(f.Path)
	if err != nil {
		fmt.Println("Failed to Write Local File:", err)
		return
	}

	// 2 注册关闭文件的 Defer
	defer file.Close()

	fmt.Fprintf(file, "Name \t Available \t Duration(ms) \n")
	for _, res := range checkResult {
		fmt.Fprintf(file, "%s \t %v \t %d \n", res.Name, res.Ok, res.Duration_ms)
	}
}
