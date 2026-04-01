package model

type Target struct {
	Name       string `json:"name"`
	Url        string `json:"url"`
	Timeout_ms int64  `json:"timeout_ms"`
}

type CheckResult struct {
	Name        string
	Ok          bool
	Duration_ms int64
	FailReason  string
}

type ResultStore interface {
	WriteResult(checkResult []CheckResult) error
}
