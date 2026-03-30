package store

import (
	"fmt"
	"service-checker/internal/model"
)

type ConsolePrintr struct{}

func (c ConsolePrintr) WriteResult(checkResult []model.CheckResult) {
	fmt.Printf("Name \t Available \t Duration(ms) \n")
	for _, res := range checkResult {
		fmt.Printf("%s \t %v \t %d \n", res.Name, res.Ok, res.Duration_ms)
	}
}
