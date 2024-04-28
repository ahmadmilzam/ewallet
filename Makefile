APP_EXECUTABLE="bin/main"

build:
	mkdir -p bin/
	go build -o $(APP_EXECUTABLE) cmd/main.go

test:
	# go test -short -count=1 -race ./...
	go test -cover ./...

run: build
	./bin/main start

migrate-create:
	go run cmd/main.go migrate create --filename $(FILENAME)

migrate-up:
	go run cmd/main.go migrate up

migrate-down:
	go run cmd/main.go migrate down

static-check:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

copy-config:
	cp ./config/config.yaml.sample ./config/config.yaml