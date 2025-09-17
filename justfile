# Variables
src-dir := "./src"
cli-dir := "./cli"
runtime-dir := "./runtime"
runtime-bin-dir := "./build/runtimes"
build-dir := "build"
metadata-dir := "build/metadata"
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


default: clean runtime build
release: clean runtime build-all checksums

run:
	go run {{src-dir}}

GO_DATE := `date +"%Y:%m:%-d:%H"`
GO_VERSION := `git describe --tags --abbrev=0 2>/dev/null || git rev-parse --short HEAD`

build:
  # Build
  mkdir -p {{build-dir}}
  go build -ldflags "-X 'luna/cli.lunaVersion={{GO_VERSION}}' -X 'luna/cli.lunaBuildDate={{GO_DATE}}'" -o {{build-dir}}/{{bin-name}} {{src-dir}}/main.go

build-all:
  # Build all
  mkdir -p {{build-dir}}
  mkdir -p {{build-dir}}/bin
  # Linux amd64
  env {{bin-linux-amd64}} go build -ldflags "-X 'luna/cli.lunaVersion={{GO_VERSION}}' -X 'luna/cli.lunaBuildDate={{GO_DATE}}'" -o {{build-dir}}/bin/{{bin-name}}_linux_amd64 {{src-dir}}/main.go
  # Linux arm64
  env {{bin-linux-arm64}} go build -ldflags "-X 'luna/cli.lunaVersion={{GO_VERSION}}' -X 'luna/cli.lunaBuildDate={{GO_DATE}}'" -o {{build-dir}}/bin/{{bin-name}}_linux_arm64 {{src-dir}}/main.go
  # Darwin amd64
  env {{bin-darwin-amd64}} go build -ldflags "-X 'luna/cli.lunaVersion={{GO_VERSION}}' -X 'luna/cli.lunaBuildDate={{GO_DATE}}'" -o {{build-dir}}/bin/{{bin-name}}_darwin_amd64 {{src-dir}}/main.go
  # Darwin arm64
  env {{bin-darwin-arm64}} go build -ldflags "-X 'luna/cli.lunaVersion={{GO_VERSION}}' -X 'luna/cli.lunaBuildDate={{GO_DATE}}'" -o {{build-dir}}/bin/{{bin-name}}_darwin_arm64 {{src-dir}}/main.go
  # Windows amd64
  env {{bin-windows-amd64}} go build -ldflags "-X 'luna/cli.lunaVersion={{GO_VERSION}}' -X 'luna/cli.lunaBuildDate={{GO_DATE}}'" -o {{build-dir}}/bin/{{bin-name}}_windows_amd64 {{src-dir}}/main.go

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

checksums:
  # Make build/metadata dir
  mkdir -p {{metadata-dir}}
  # Checksums
  cd {{build-dir}}/bin; find . -type f ! -name "checksums.txt" -exec sha256sum {} \; > ../../{{metadata-dir}}/checksums.txt

check:
  # Check files integrity
  cd {{build-dir}}/bin; sha256sum -c ../../{{metadata-dir}}/checksums.txt

clean:
  # Clean
  rm -rf {{build-dir}}

test:
	go test ./...

# Ayuda general
help:
	@echo "Comandos disponibles:"
	@echo "  just run              # Ejecuta el runtime (modo desarrollo)"
	@echo "  just build            # Compila el binario CLI en ./build/luna"
	@echo "  just clean            # Elimina binarios y archivos temporales"
	@echo "  just test             # Ejecuta los tests de Go"
