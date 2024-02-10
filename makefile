build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o OpenNomad.bin

run: build
	./OpenNomad.bin

clean:
	go clean
	@rm web/app.wasm
	@rm web/styles/style.css
	@rm OpenNomad.bin