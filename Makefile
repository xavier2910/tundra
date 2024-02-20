binaries := thetundra

build: $(binaries)

$(binaries):
	@go mod tidy
	@mkdir build
	@go build -o ./build/$@ cmd/$@/main.go

install:
	@go install github.com/xavier2910/tundra

clean:
	@rm -rf build