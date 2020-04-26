emit-log:
	go run sender/emit.go

emit-error:
	go run sender/emit.go error

receive-all:
	go run receiver/receive.go info warning error

receive-warning-error:
	go run receiver/receive.go warning error > logs_from_rabbit.log