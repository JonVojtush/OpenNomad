build:
	GOOS=wasip1 GOARCH=wasm go build -o web/app.wasm main.go
	go build -o OpenNomad.bin
	
run: build
	./OpenNomad.bin

clean:
	go clean
	@rm web/app.wasm
	@rm OpenNomad.bin