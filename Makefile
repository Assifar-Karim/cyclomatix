BINARY_NAME=cyclo

build:
	@echo "[INFO] Building linux amd64 binary"
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/cyclomatix/main.go
	
	@echo "[INFO] Building linux arm64 binary"
	GOARCH=arm64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-arm64 cmd/cyclomatix/main.go
	
	@echo "[INFO] Building darwin amd64 binary"
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-amd64 cmd/cyclomatix/main.go
	
	@echo "[INFO] Building darwin arm64 binary"
	GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/cyclomatix/main.go
	
	@echo "[INFO] Building windows amd64 binary"
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows-amd64 cmd/cyclomatix/main.go
package: build
	@echo ["INFO"] Compressing linux amd64 binaries
	@mv bin/${BINARY_NAME}-linux-amd64 bin/${BINARY_NAME}
	tar -czvf bin/cyclomatix-linux-amd64.tar.gz bin/${BINARY_NAME}
	@rm bin/${BINARY_NAME}
	
	@echo ["INFO"] Compressing linux arm64 binaries
	@mv bin/${BINARY_NAME}-linux-arm64 bin/${BINARY_NAME}
	tar -czvf bin/cyclomatix-linux-arm64.tar.gz bin/${BINARY_NAME}
	@rm bin/${BINARY_NAME}
	
	@echo ["INFO"] Compressing darwin amd64 binaries
	@mv bin/${BINARY_NAME}-darwin-amd64 bin/${BINARY_NAME}
	tar -czvf bin/cyclomatix-darwin-amd64.tar.gz bin/${BINARY_NAME}
	@rm bin/${BINARY_NAME}
	
	@echo ["INFO"] Compressing darwin arm64 binaries
	@mv bin/${BINARY_NAME}-darwin-arm64 bin/${BINARY_NAME}
	tar -czvf bin/cyclomatix-darwin-arm64.tar.gz bin/${BINARY_NAME}
	@rm bin/${BINARY_NAME}
	
	@echo ["INFO"] Compressing windows amd64 binaries
	@mv bin/${BINARY_NAME}-windows-amd64 bin/${BINARY_NAME}.exe
	zip bin/cyclomatix-windows-amd64.zip bin/${BINARY_NAME}.exe
	@rm bin/${BINARY_NAME}.exe