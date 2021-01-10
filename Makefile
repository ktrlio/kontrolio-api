.PHONY: build clean deploy

build:
	env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o dist/authorizer functions/authorizer/authorizer.go
	env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o dist/login functions/login/login.go
	env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o dist/create_user functions/create_user/create_user.go
	env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o dist/create_record functions/create_record/create_record.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
