## `std/web/storage/README.md`

# Storage

**Descripción:**
Este módulo provee acceso a **`localStorage`** y **`sessionStorage`** del navegador directamente desde Lua.
Permite guardar, recuperar y eliminar datos clave–valor de forma persistente (local) o temporal (sesión).

---

## API

### `localStorage`

- `setItem(key, value)` → Guarda un valor persistente (hasta que se borre manualmente o el navegador lo limpie).
- `getItem(key)` → Devuelve el valor asociado a `key`, o `nil` si no existe.
- `removeItem(key)` → Elimina el valor asociado a `key`.
- `clear()` → Borra todas las claves y valores.

### `sessionStorage`

- `setItem(key, value)` → Guarda un valor temporal (se borra al cerrar la pestaña/ventana).
- `getItem(key)` → Devuelve el valor asociado a `key`, o `nil` si no existe.
- `removeItem(key)` → Elimina el valor asociado a `key`.
- `clear()` → Borra todas las claves y valores.

---

## Ejemplos

### Uso de `localStorage` (persistente)

```lua
local storage = require("std:web").storage
local console = require("std:web").console

-- Guardar un valor
storage.localStorage.setItem("theme", "dark")

-- Recuperar un valor
local theme = storage.localStorage.getItem("theme")
if theme then
    console.log("Tema guardado:", theme)
else
    console.log("No hay tema configurado")
end

-- Modificarlo
storage.localStorage.setItem("theme", "light")
console.log("Tema cambiado a:", storage.localStorage.getItem("theme"))

-- Eliminarlo
storage.localStorage.removeItem("theme")
console.log("Después de eliminar:", storage.localStorage.getItem("theme")) -- nil

-- Limpiar todo el almacenamiento persistente
storage.localStorage.clear()
console.log("Todo el localStorage borrado")
```

---

### Uso de `sessionStorage` (temporal)

```lua
local storage = require("std:web").storage
local console = require("std:web").console

-- Guardar un valor de sesión
storage.sessionStorage.setItem("username", "wuXan")

-- Recuperar el valor
local user = storage.sessionStorage.getItem("username")
console.log("Usuario de sesión:", user)

-- Eliminarlo
storage.sessionStorage.removeItem("username")

-- Limpiar todo el almacenamiento de sesión
storage.sessionStorage.clear()
```

---

## Notas

- Tanto `localStorage` como `sessionStorage` solo aceptan **strings**.
  Si necesitas guardar objetos o tablas, convierte a JSON antes:

  ```lua
  local json = require("std:json")
  storage.localStorage.setItem("profile", json.encode({ name = "wuXan", age = 23 }))
  ```

- `localStorage` es persistente entre sesiones del navegador.
- `sessionStorage` se elimina automáticamente al cerrar la pestaña.
