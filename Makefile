.PHONY: build run clean

build:
	go build -o expense main.go expense.go

run: build
	./expense

clean:
	rm -f expense

test:
	go test -v