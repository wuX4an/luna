## `std/web/fetch/README.md`

# Fetch

**Descripción:** Realiza peticiones HTTP desde Lua (AJAX/Fetch), compatible con `await`.

### Funciones

```lua
-- fetch.get(url)
-- fetch.post(url, opts)
-- Ambas devuelven una función que puede usarse con await:
-- local res = await(fetch.get, url)
-- local res = await(fetch.post, url, opts)

-- opts opcional para POST:
-- {
--   headers = { ["Content-Type"] = "application/json" },
--   body = '{"name":"WuXan"}',
-- }
```

### Ejemplo de uso con `await`

```lua
local fetch = require("std:web").fetch
local await = require("std:web").await
local console = require("std:web").console

print("inicio")

-- GET
local data = await(fetch.get("https://httpbin.org/get"))
console.log("GET Status:", data.status)
console.log("GET Body:", data.text)

-- POST
local post = await(fetch.post(
    "https://httpbin.org/post",
    { headers = { ["Content-Type"] = "application/json" }, body = '{"name":"WuXan"}' }
))
console.log("POST Status:", post.status)
console.log("POST Body:", post.text)

print("final")
```

### Notas

- `fetch.get` y `fetch.post` **ya no requieren callback**.
- `await(fn, ...)` suspende la coroutine hasta que la petición termine.
- Las respuestas devuelven una tabla Lua con:

  ```lua
  {
      status = 200,
      ok     = true,
      body   = "<respuesta del servidor>"
  }
  ```

### Devnote

Añadir una función callback opcional y que pueda retornar la brujeria que hace esa función,
dandole ese valor a la variable asignada. Pero enrealidad... ¿Quién usa la función callback de fetch?
Ni modo, problema del futuro hacer el modulo async, si lo ignoro lo suficiente no me molestará en mi vida :)
