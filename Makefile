gen:
	go generate ./canceler

tests:
	go test ./canceler

bench:
	go test ./canceler -bench=.