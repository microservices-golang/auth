LOCAL_BIN := C:/Users/Denis/Desktop/microservices-golang/auth/bin

install-golangci-lint:
	@if not exist "$(LOCAL_BIN)/golangci-lint.exe" (
		echo Installing golangci-lint...
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	) else (
		echo golangci-lint.exe already installed.
	)

lint:
	$(LOCAL_BIN)/golangci-lint.exe run ./... --config .golangci.pipeline.yaml