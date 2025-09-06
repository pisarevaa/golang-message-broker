.PHONY: help unit_test integration_test e2e_test test lint build_api
help:
	cat Makefile

test:
	go clean -testcache
	go test -v ./internal/... -race -coverpkg=./...  -coverprofile=coverage.out

# benchmark_test:
# 	go clean -testcache
# 	go test -v ./benchmark_tests/... -race -bench=. -benchtime=100x -benchmem

# benchmark_test_with_export_result:
# 	go clean -testcache
# 	go test -v ./benchmark_tests/... -race -bench=. -benchtime=100x -benchmem > bench_results.out

# benchmark_compare_results:
# 	benchstat bench_results.out bench_results2.out

# benchmark_test_with_profiles:
# 	go clean -testcache
# 	go test -v ./benchmark_tests/... -race -bench=. -benchtime=100x -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof
# 	# go tool pprof cpu.prof

# mock_generate:
# 	mockgen -source=internal/api/storage/kafka/types.go -destination=internal/api/mocks/producer.go -package=mock Producer
# 	mockgen -source=internal/api/storage/db/types.go -destination=internal/api/mocks/storage.go -package=mock Storage

swag_generate:
	swag init --dir cmd,internal --output ./docs

swag_format:
	swag fmt

generate_enums:
	go generate ./...

lint:
	go fmt ./...
	find . -name '*.go' -exec goimports -w {} +
	find . -name '*.go' -exec golines -w {} -m 120 \;
	golangci-lint run ./...

run_api:
	go run ./cmd

build_api:
	CGO_ENABLED=0 go build -o api ./cmd

profiling_cpu:
	CGO_ENABLED=0 go build -o api ./cmd
	go tool pprof api cpu.prof
