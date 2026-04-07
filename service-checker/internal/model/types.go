package model

type Target struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	TimeoutMS int64  `json:"timeout_ms"`
}

type CheckResult struct {
	Name       string
	OK         bool
	DurationMS int64
	FailReason string
}

type ResultStore interface {
	WriteResult(checkResult []CheckResult) error
}
