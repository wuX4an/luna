# Cookie

**Description:**
Manages cookies from Lua.

---

### Functions

```lua
cookie.set(name, value, options)   -- options: {expires, path, domain, secure}
cookie.get(name)                   -- Returns value or nil
cookie.delete(name)                -- Deletes cookie
cookie.list()                      -- Returns a table with all cookies
```

---

### Example

```lua
cookie.set("username", "dog", { path = "/", expires = 3600 })
print(cookie.get("username"))
cookie.delete("username")
```

---
