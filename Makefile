APP_EXECUTABLE="bin/ewallet"

build:
	mkdir -p bin/
	go build -o $(APP_EXECUTABLE) cmd/ewallet/main.go

run: build
	./bin/ewallet start

migrate-create:
	go run cmd/ewallet/main.go migrate create --filename ${FILENAME}

migrate-up:
	go run cmd/ewallet/main.go migrate up

migrate-down:
	go run cmd/ewallet/main.go migrate down

test:
	go test -short -count=1 -race ./...

static-check:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

copy-config:
	cp ./configs/config.yaml.example ./configs/config.yaml

