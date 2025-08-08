# Variables
src-dir := "./src"
cli-dir := "./cli"
runtime-dir := "./runtime"
runtime-bin-dir := "./build/runtimes"
build-dir := "build"
bin-name := "luna"
go-flags := "-ldflags='-s -w'"
go-cgo := "CGO_ENABLED=0"

bin-linux-amd64 := "GOOS=linux GOARCH=amd64"
bin-linux-arm64 := "GOOS=linux GOARCH=arm64"
bin-darwin-amd64 := "GOOS=darwin GOARCH=amd64"
bin-darwin-arm64 := "GOOS=darwin GOARCH=arm64"
bin-windows-amd64 := "GOOS=windows GOARCH=amd64"

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
	env {{bin-linux-amd64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_linux_amd64 {{runtime-dir}}
	env {{bin-linux-arm64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_linux_arm64 {{runtime-dir}}
	env {{bin-darwin-amd64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_darwin_amd64 {{runtime-dir}}
	env {{bin-darwin-arm64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_darwin_amd64 {{runtime-dir}}
	env {{bin-windows-amd64}} {{go-cgo}} go build {{go-flags}} -o {{runtime-bin-dir}}/runtime_windows_amd64 {{runtime-dir}}


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
