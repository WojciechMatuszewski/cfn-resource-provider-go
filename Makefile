.PHONY: build test clean deploy

build:
	TAGS=logging make -f makebuild  # this runs build steps required by the cfn cli

deploy:
	cfn submit --set-default

test:
	cfn generate
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler cmd/main.go

clean:
	rm -rf bin
