# Variables
src-dir := "./src"
cli-dir := "./cli"
runtime-dir := "./runtime"
runtime-bin-dir := "./build/runtimes"
build-dir := "build"
bin-name := "luna"
go-flags := "-ldflags='-s -w'"
go-cgo := "CGO_ENABLED=0"
go-root := "$(go env GOROOT)"

bin-linux-amd64 := "GOOS=linux GOARCH=amd64"
bin-linux-arm64 := "GOOS=linux GOARCH=arm64"
bin-darwin-amd64 := "GOOS=darwin GOARCH=amd64"
bin-darwin-arm64 := "GOOS=darwin GOARCH=arm64"
bin-windows-amd64 := "GOOS=windows GOARCH=amd64"
bin-js-wasm := "GOOS=js GOARCH=wasm"


default: runtime build

# Ejecutar el runtime directamente (modo desarrollo)
run:
	go run {{src-dir}}

# Compilar el binario CLI de Luna (con todos los comandos)
build:
	mkdir -p {{build-dir}}
	CGO_ENABLED=0 go build {{go-flags}} -o {{build-dir}}/{{bin-name}} {{src-dir}}/main.go

runtime:
  mkdir -p {{runtime-bin-dir}}
  # Make linux/amd64
  env {{bin-linux-amd64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_linux_amd64 {{runtime-dir}}
  # Make linux/arm64
  env {{bin-linux-arm64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_linux_arm64 {{runtime-dir}}
  # Make darwin/amd64
  env {{bin-darwin-amd64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_darwin_amd64 {{runtime-dir}}
  # Make darwin/arm64
  env {{bin-darwin-arm64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_darwin_amd64 {{runtime-dir}}
  # Make windows/amd64
  env {{bin-windows-amd64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_windows_amd64 {{runtime-dir}}
  # Make js/wasm
  env {{bin-js-wasm}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_js_wasm {{runtime-dir}}
  mkdir -p {{build-dir}}/wasm
  cp {{go-root}}/lib/wasm/wasm_exec.js {{build-dir}}/wasm/wasm.js
  chmod 644 {{build-dir}}/wasm/wasm.js
  cp std/web/index.html build/wasm
  cp std/web/sw.js build/wasm

# Limpiar binarios y temporales
clean:
	rm -rf {{build-dir}} *.tar.gz *.exe *.out

# Ejecutar tests de Go
test:
	go test ./...

# Empaquetar un script Lua como ejecutable (ej: just package src)
package script:
	{{build-dir}}/{{bin-name}} build {{script}}

# Ayuda general
help:
	@echo "Comandos disponibles:"
	@echo "  just run              # Ejecuta el runtime (modo desarrollo)"
	@echo "  just build            # Compila el binario CLI en ./build/luna"
	@echo "  just clean            # Elimina binarios y archivos temporales"
	@echo "  just test             # Ejecuta los tests de Go"
	@echo "  just package <file>   # Empaqueta un script Lua como ejecutable"
