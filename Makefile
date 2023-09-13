BINARY_NAME=go-mini-kv-store
 
build:
	go build -o ${BINARY_NAME} ./cmd/go-mini-kv-store/main.go
 
clean:
	go clean
	rm ${BINARY_NAME}

vet:
	go vet ./...

test-e2e:
	export URL=http://localhost:3000 ;\
	go test -count=1 -v ./test/e2e/... 

run-dev:
	export PORT=3000
	gow run ./cmd/go-mini-kv-store/

run-ci:
	docker build -t go-mini-kv-store . ;\
	docker run -p 3000:3000 -d go-mini-kv-store