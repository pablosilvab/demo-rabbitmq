package rabbit

import (
	"time"

	"github.com/pablosilvab/elastic-lib"
)

type Log struct {
	Queue     string
	Timestamp time.Time
	Operation string
	Result    string
	Error     string
}

func WriteLog(queue string, operation string, status string, desc string) {
	log := Log{
		Queue:     queue,
		Timestamp: time.Now(),
		Operation: operation,
		Result:    status,
		Error:     desc,
	}
	elastic.Log("rabbit", log)
}
