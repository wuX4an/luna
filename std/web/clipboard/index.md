## `std/web/clipboard/README.md`

# Clipboard

**Descripción:** Acceso al portapapeles del navegador desde Lua usando `await`.

### Funciones

```lua
clipboard.write(text)       -- Copia texto al portapapeles (await-compatible)
clipboard.read()            -- Devuelve texto actualmente en el portapapeles (await-compatible)
clipboard.clear()           -- Borra el portapapeles (no todos los navegadores lo soportan)
```

```lua
local clipboard = require("std:web").clipboard
local await = require("std:web").await

-- Leer texto
local text = await(clipboard.read) -- pasa la función Read, que acepta un callback
print(text)
-- Copiar texto
clipboard.write("Hola") -- Write también debe aceptar callback

-- Limpiar portapapeles
clipboard.clear()

```

### Notas

- `write` y `read` dependen de permisos del navegador; algunos requieren que la acción sea disparada por un evento de usuario (click, keypress).
- `clear` no está soportado en todos los navegadores.
- Ahora todas las funciones son **await-compatible**, lo que simplifica su uso en corutinas y evita callbacks explícitos.
