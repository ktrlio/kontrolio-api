.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o dist/create_user functions/create_user/create_user.go
	env GOOS=linux go build -ldflags="-s -w" -o dist/create_record functions/create_record/create_record.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
