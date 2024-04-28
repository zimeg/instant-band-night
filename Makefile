.PHONY: clean build

# Remove provisioned files
clean:
	rm -rf ./dist

# Create a development build
build:
	go build -o ./dist/ibn
