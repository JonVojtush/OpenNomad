build:
	tinygo build -o web/app.wasm -target wasm
	go build -o OpenNomad.bin
	
run: build
	./OpenNomad.bin

clean:
	go clean
	@rm web/app.wasm
	@rm web/styles/style.css
	@rm OpenNomad.bin