APP=fazztrack-vehicle
APP_EXE="./build/$(APP)"

runs:
	go run main.go server

run:
	./build/$(APP) server
build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=darwin go build -o ${APP_EXE}