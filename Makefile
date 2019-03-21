testconvey:
	go test ./convey/... -v -race -short

testsimple:
	go test ./simple/... -v -race -short

testall:
	go test ./... -v -race -short

testwithexclude:
	go test ./... -v -race -short | grep -v mock