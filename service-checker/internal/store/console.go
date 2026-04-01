package store

import (
	"fmt"
	"service-checker/internal/model"
)

type ConsolePrintr struct{}

func (c ConsolePrintr) WriteResult(checkResult []model.CheckResult) error {
	fmt.Printf("Name \t Available \t Duration(ms) \t Fail Reason \n")
	for _, res := range checkResult {
		fmt.Printf("%s \t %v \t %d \t %s \n", res.Name, res.Ok, res.Duration_ms, res.FailReason)
	}

	return nil
}
