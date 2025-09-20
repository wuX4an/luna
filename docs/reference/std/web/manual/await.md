# Await

**Description:**
Allows suspending a Lua coroutine until an asynchronous (callback-based) function returns a result. Works like JavaScript’s `await` in Lua within WASM environments.

---

### Main Function

```lua
await(promiseFn) -> value
```

- `promiseFn`: Lua function that accepts a single callback argument. The function must invoke the callback with the result when the asynchronous operation completes.
- Returns: the value passed to the callback, directly in Lua.

---

### Example Usage with `fetch`

```lua
local fetch = require("std:web").fetch
local await = require("std:web").await
local console = require("std:web").console

print("start")

-- GET with await
local data = await(fetch.get("https://httpbin.org/get"))
console.log("GET Status:", data.status)
console.log("GET Body:", data.text)

-- POST with await
local post = await(fetch.post(
    "https://httpbin.org/post",
    { headers = { ["Content-Type"] = "application/json" }, body = '{"name":"WuXan"}' }
))
console.log("POST Status:", post.status)
console.log("POST Body:", post.text)

print("end")
```

---

### How It Works Internally

1. `await` receives a `promiseFn` function.
2. Creates a Go **channel** (`resultChan`) to receive the result.
3. Creates a Lua **callback** that, when executed, sends the result to the channel.
4. Calls `promiseFn` passing the callback.
5. Suspends the coroutine until the channel receives a value.
6. Returns the value to Lua.

---

### Notes

- Ideal for WASM asynchronous APIs like `fetch`, `setTimeout`, `setInterval`.
- Allows writing synchronous-looking Lua code that executes asynchronously in JS.
- Eliminates the need for explicit callbacks in Lua code.

---
