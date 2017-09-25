run: run_url

run_url:
	echo 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run main.go -type url $(flags)

run_file:
	echo '/etc/passwd\n/etc/hosts' | go run main.go -type file $(flags)

test:
	go test ./... $(flags) -tags=integration