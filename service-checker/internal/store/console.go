package store

import (
	"fmt"
	"service-checker/internal/model"
)

type ConsolePrinter struct{}

func (c ConsolePrinter) WriteResult(checkResult []model.CheckResult) error {
	fmt.Printf("Name \t Available \t Duration(ms) \t Fail Reason \n")
	for _, res := range checkResult {
		fmt.Printf("%s \t %v \t %d \t %s \n", res.Name, res.OK, res.DurationMS, res.FailReason)
	}

	return nil
}
