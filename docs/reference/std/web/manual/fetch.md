# Fetch

**Description:**
Performs HTTP requests (AJAX/Fetch) from Lua, compatible with `await`.

---

### Functions

```lua
-- GET request
fetch.get(url)

-- POST request with optional options
fetch.post(url, opts)

-- Both return a function usable with await:
-- local res = await(fetch.get, url)
-- local res = await(fetch.post, url, opts)

-- Example POST opts:
-- {
--   headers = { ["Content-Type"] = "application/json" },
--   body    = '{"name":"WuXan"}',
-- }
```

---

### Example with `await`

```lua
local fetch = require("std:web").fetch
local await = require("std:web").await
local console = require("std:web").console

print("Start")

-- GET request
local data = await(fetch.get("https://httpbin.org/get"))
console.log("GET Status:", data.status)
console.log("GET Body:", data.text)

-- POST request
local post = await(fetch.post(
    "https://httpbin.org/post",
    { headers = { ["Content-Type"] = "application/json" }, body = '{"name":"WuXan"}' }
))
console.log("POST Status:", post.status)
console.log("POST Body:", post.text)

print("End")
```

---

### Notes

- `fetch.get` and `fetch.post` **no longer require callbacks**.
- `await(fn, ...)` suspends the coroutine until the request completes.
- Responses return a Lua table:

```lua
{
    status = 200,
    ok     = true,
    body   = "<server response>"
}
```

---

### Devnote

> Future: could add an optional callback to `fetch`, returning the same value to the assigned variable. In practice, the callback is rarely used, so asynchronous handling via `await` is sufficient.

---
