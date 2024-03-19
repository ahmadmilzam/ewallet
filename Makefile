APP_EXECUTABLE="bin/ewallet"

build:
	mkdir -p bin/
	go build -o $(APP_EXECUTABLE) cmd/server/main.go

run: build
	./bin/ewallet start

test:
	go test -short -count=1 -race ./...

static-check:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

copy-config:
	cp ./configs/config.yaml.example ./configs/config.yaml

