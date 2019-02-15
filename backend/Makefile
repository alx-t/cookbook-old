TARGET=cookbook
PROD=beta/cookbook
PROD_PATH=

all: clean format build

beta: clean format build_beta

clean:
	rm -rf $(TARGET)

format:
	gofmt -l -s -w .

build:
	go build -o $(TARGET) main.go

clean_beta:
	rm -rf $(PROD)

build_beta:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(PROD) main.go

deploy_beta:

run:
	docker run -d --rm \
	-p 3007:3007 \
	--name=go_cookbook \
	cookbook
