## `std/web/await/README.md`

# Await

**Descripción:** Permite suspender una coroutine de Lua hasta que una función asíncrona (basada en callback) devuelva un resultado.
Funciona como un `await` de JavaScript dentro de Lua en entornos WASM.

---

### Función principal

```lua
await(promiseFn) -> value
```

- `promiseFn`: función Lua que acepta un callback como único argumento.
  La función debe invocar el callback pasando como argumento el resultado cuando la operación asíncrona termine.

- Devuelve: el valor pasado al callback, directamente a Lua.

---

### Ejemplo de uso con `fetch`

```lua
local fetch = require("std:web").fetch
local await = require("std:web").await
local console = require("std:web").console

print("inicio")

-- GET con await
local data = await(fetch.get("https://httpbin.org/get"))
console.log("GET Status:", data.status)
console.log("GET Body:", data.text)

-- POST con await
local post = await(fetch.post(
    "https://httpbin.org/post",
    { headers = { ["Content-Type"] = "application/json" }, body = '{"name":"WuXan"}' }
))
console.log("POST Status:", post.status)
console.log("POST Body:", post.text)

print("final")
```

---

### Cómo funciona internamente

1. `await` recibe una función `promiseFn`.
2. Crea un **canal** Go (`resultChan`) para recibir el resultado.
3. Crea un **callback Lua**, que al ejecutarse envía el resultado al canal.
4. Llama a `promiseFn` pasando el callback.
5. Suspende la coroutine hasta que el canal reciba un valor.
6. Devuelve el valor a Lua.

---

### Notas

- Ideal para usar con APIs asíncronas WASM, como `fetch`, `setTimeout`, `setInterval`.
- Permite escribir código síncrono en Lua que en realidad se ejecuta de manera asíncrona en JS.
- Evita usar callbacks explícitos en el código Lua.
