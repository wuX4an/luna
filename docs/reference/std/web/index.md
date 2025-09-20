# std:web

**Description:**
Central entry point for all web-related functionality in Lua via WebAssembly (`syscall/js`). Provides access to modules for DOM manipulation, events, fetch requests, storage, notifications, audio, keyboard, clipboard, animations, timers, and console logging.

---

### Modules

```lua
local web = require("std:web")
```

- **`web.animate`** → Animate DOM elements with CSS transforms and opacity.
- **`web.audio`** → Play and control audio elements.
- **`web.await`** → Suspend Lua coroutines until asynchronous JS callbacks complete.
- **`web.clipboard`** → Read from and write to the clipboard, await-compatible.
- **`web.console`** → Log messages to the browser console (`log`, `warn`, `error`, `info`).
- **`web.cookie`** → Set, get, delete, and list cookies.
- **`web.dom`** → Create, modify, and manage HTML elements.
- **`web.fetch`** → Make HTTP requests with `GET` and `POST`, compatible with `await`.
- **`web.keyboard`** → Listen to keyboard events, combinations, and disable default key actions.
- **`web.notify`** → Display browser notifications.
- **`web.storage`** → Access `localStorage` and `sessionStorage`.
- **`web.timer`** → `setTimeout`, `setInterval`, `clearTimeout`, `clearInterval`.

---

### Example Usage

```lua
local web = require("std:web")
local DOM = web.dom
local console = web.console
local await = web.await
local fetch = web.fetch

-- Create a div
local div = DOM.div({ id = "app" }, "Hello World")
DOM.root():append(div)
div:setStyle({ padding = "20px", backgroundColor = "#09f", color = "#fff" })

-- Fetch example
local data = await(fetch.get("https://httpbin.org/get"))
console.log("Fetched status:", data.status)

-- Notification
local notify = web.notify
if notify.permission() == "granted" then
    notify.show("Hello!", { body = "Notification from Lua" })
end
```

---

### Notes

- Acts as a **single unified interface** to all browser APIs accessible from Lua.
- All modules are designed to work with WebAssembly and the Lua coroutine model (`await`).
- This is the **main entry point**, so you only need to `require("std:web")` to access any web module.

---
