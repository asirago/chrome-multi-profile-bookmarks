
build:
	go build -o cmpbs

release:
	CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o cmpbs-amd64 ./src
	CGO_ENABLED=0 GOARCH=arm64 go build -ldflags="-s -w" -o cmpbs-arm64 ./src

	# Combine into universal binary
	lipo -create -output cmpbs cmpbs-amd64 cmpbs-arm64

	# Clean up
	rm cmpbs-amd64 cmpbs-arm64
