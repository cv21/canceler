test:
	go test .

bench:
	go test . -bench=.

dep:
	dep ensure -v