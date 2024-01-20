build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: build
	./OpenNomad_WASM

clean:
	go clean
	@rm web/app.wasm