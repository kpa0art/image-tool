.Phony: build

build_windows:
	rm -rf build_windows && mkdir build_windows && env GOOS=windows GOARCH=amd64 go build -o build_windows/grid_generator.exe -v ./cmd/grid-generator

build_linux:
	rm -rf build_linux && mkdir build_linux && env GOOS=linux GOARCH=amd64 go build -o build_linux/grid_generator -v ./cmd/grid-generator


