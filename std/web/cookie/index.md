## `std/web/cookie/README.md`

# Cookie

**Descripción:** Manejo de cookies desde Lua.

### Funciones

```lua
cookie.set(name, value, options)   -- options: {expires, path, domain, secure}
cookie.get(name)                   -- Devuelve valor o nil
cookie.delete(name)                -- Borra cookie
cookie.list()                      -- Retorna tabla con todas las cookies
```

Ejemplo:

```lua
cookie.set("username", "wuXan", { path = "/", expires = 3600 })
print(cookie.get("username"))
cookie.delete("username")
```
