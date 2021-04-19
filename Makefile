init:
	go mod init test && go mod tidy

build:
	docker build -t test .

dev:
	docker run -p 8888:80 -it --rm test



