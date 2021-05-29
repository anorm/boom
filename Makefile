boom:
	go build -o ./bin/boom ./cmd/boom/

fmt:
	gofmt -w .

clean:
	rm -rf ./bin/
