include .env


local-build-run: statistic-local-build statistic-run

statistic-local-build:
	@echo "  >  Building statictic service ...."
	go build -o ./cmd/statistic/service ./cmd/statistic


statistic-run:
	@echo "  >  Building statistic service ...."
	go build -o ./cmd/statistic/service ./cmd/statistic
	@echo "  >  statistic run ...."
	go run ./cmd/statistic/main.go statistic

#поиск состояний гонки
statistic-run-race:
	@echo "  >  Building statistic service ...."
	go build -race -o ./cmd/statistic/service ./cmd/statistic
	@echo "  >  statistic run ...."
	go run -race ./cmd/statistic/main.go statistic


grpc:
	protoc --go_out=pkg/grpc --go_opt=paths=source_relative  --go-grpc_out=pkg/grpc/  --go-grpc_opt=paths=source_relative protos/*.proto  && mv ./pkg/grpc/protos/* ./pkg/grpc/


#https://httpd.apache.org/docs/2.4/programs/ab.html
bench:
	ab -m POST -T application/json  -c 12 -n 400 "http://0.0.0.0:5313/collect_statistic/get/history_client/0/50"

pprof:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/profile
allocs:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/allocs
block:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/block
goroutine:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/goroutine
heap:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/heap
mutex:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/mutex
threadcreate:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/threadcreate
trace:
	go tool pprof goprofex http://127.0.0.1:10001/debug/pprof/trace