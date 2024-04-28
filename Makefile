.PHONY: check clean build

# Confirm formatted lints
check:
	golangci-lint run ./...

# Remove provisioned files
clean:
	rm -rf ./dist

# Create a development build
build: check
	go build -o ./dist/ibn
