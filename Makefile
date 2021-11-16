.PHONY: all build_dir test bench pprof trace test_all clean analyze fmt vet tidy lint cyclo tools

all: build_dir test bench trace test_all

build_dir:
	mkdir -p ./build/

test: build_dir tidy analyze
	go test -covermode=count -coverprofile=./build/profile.out ./...
	if [ -f ./build/profile.out ]; then go tool cover -func=./build/profile.out; fi

bench: build_dir
	go test \
		-o=./build/bench.test \
		-bench=. \
		-cpuprofile=./build/cpuprofile.out \
		-memprofile=./build/memprofile.out \
		-blockprofile=./build/blockprofile.out \
		-mutexprofile=./build/mutexprofile.out \
		-benchtime=5s \
		./...

pprof_web:
	go tool pprof -http=localhost:8080 ./build/cpuprofile.out
	go tool pprof -http=localhost:8081 ./build/memprofile.out
	go tool pprof -http=localhost:8082 ./build/blockprofile.out
	go tool pprof -http=localhost:8083 ./build/mutexprofile.out

trace: build_dir
	go test -trace=./build/trace.out .
	if [ -f ./build/trace.out ]; then go tool trace ./build/trace.out; fi

test_all: test
	go test all

clean:
	rm -rf ./build
	go clean -cache

analyze: fmt vet lint cyclo

fmt:
	gofmt -w -s -d .

vet: lint cyclo
	go vet ./...

tidy:
	go mod tidy
	go mod verify

lint:
	golint ./...

cyclo:
	gocyclo -avg -over 15 -ignore "_test|Godeps|vendor/" -total .

tools:
	go get -u github.com/go-delve/delve/cmd/dlv@master
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u github.com/cweill/gotests/...
	go get -u golang.org/x/lint/golint
	go get -u github.com/fzipp/gocyclo/cmd/gocyclo
	go get -u github.com/amitbet/gorename
	go get -u golang.org/x/tools/cmd/benchcmp