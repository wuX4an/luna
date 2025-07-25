# Variables
src-dir := "./src"
bin-name := "luna"
build-dir := "build"
go-flags := "-ldflags='-s -w'"

# Ejecutar el runtime directamente
run:
	go run {{src-dir}}

# Compilar el binario
build:
	go build {{go-flags}} -o {{build-dir}}/{{bin-name}} {{src-dir}}

# Limpiar binarios
clean:
	rm -rf {{build-dir}}

# Test de todo el proyecto
test:
	go test ./...

# Mostrar ayuda
help:
	@echo "Comandos disponibles:"
	@echo "  just run      # Ejecuta el runtime"
	@echo "  just build    # Compila el binario en ./build/luna"
	@echo "  just clean    # Elimina los binarios"
	@echo "  just test     # Ejecuta todos los tests"
