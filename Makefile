cli:
	go build -mod vendor -o bin/hello-world cmd/hello-world/main.go
	go build -mod vendor -o bin/read-file cmd/read-file/main.go

lambda:
	if test -f main; then rm -f main; fi
	if test -f hello-world.zip; then rm -f hello-world.zip; fi
	GOOS=linux go build -mod vendor -o main cmd/hello-world/main.go
	zip hello-world.zip main
	rm -f main

docker:
	docker build -f Dockerfile -t hello-world .

docker-readfile:
	docker build -f Dockerfile.readfile -t read-file .

local:
	docker run -e SFOMUSEUM_MODE=lambda -p 9000:8080 hello-world:latest /main

local-readfile:
	docker run -e SFOMUSEUM_MODE=lambda -e SFOMUSEUM_BUCKET_URI=file:///usr/local/example -p 9000:8080 read-file:latest /main
