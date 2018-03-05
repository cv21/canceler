gen:
	go generate .

test:
	go test .

bench:
	go test . -bench=.