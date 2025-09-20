# Storage

**Description:**
Provides access to the browser's **`localStorage`** and **`sessionStorage`** directly from Lua.
Allows storing, retrieving, and deleting key–value data persistently (local) or temporarily (session).

---

## API

### `localStorage`

- `setItem(key, value)` → Stores a persistent value (until manually removed or cleared by the browser).
- `getItem(key)` → Returns the value associated with `key`, or `nil` if it does not exist.
- `removeItem(key)` → Deletes the value associated with `key`.
- `clear()` → Deletes all keys and values.

### `sessionStorage`

- `setItem(key, value)` → Stores a temporary value (cleared when the tab/window closes).
- `getItem(key)` → Returns the value associated with `key`, or `nil` if it does not exist.
- `removeItem(key)` → Deletes the value associated with `key`.
- `clear()` → Deletes all keys and values.

---

## Examples

### Persistent `localStorage`

```lua
local storage = require("std:web").storage
local console = require("std:web").console

-- Store a value
storage.localStorage.setItem("theme", "dark")

-- Retrieve a value
local theme = storage.localStorage.getItem("theme")
if theme then
    console.log("Saved theme:", theme)
else
    console.log("No theme set")
end

-- Modify value
storage.localStorage.setItem("theme", "light")
console.log("Theme changed to:", storage.localStorage.getItem("theme"))

-- Remove value
storage.localStorage.removeItem("theme")
console.log("After removal:", storage.localStorage.getItem("theme")) -- nil

-- Clear all persistent storage
storage.localStorage.clear()
console.log("All localStorage cleared")
```

### Temporary `sessionStorage`

```lua
local storage = require("std:web").storage
local console = require("std:web").console

-- Store a session value
storage.sessionStorage.setItem("username", "wuXan")

-- Retrieve the value
local user = storage.sessionStorage.getItem("username")
console.log("Session user:", user)

-- Remove value
storage.sessionStorage.removeItem("username")

-- Clear all session storage
storage.sessionStorage.clear()
```

---

## Notes

- Both `localStorage` and `sessionStorage` **only accept strings**. To store objects or tables, encode them as JSON:

```lua
local json = require("std:json")
storage.localStorage.setItem("profile", json.encode({ name = "wuXan", age = 23 }))
```

- `localStorage` persists across browser sessions.
- `sessionStorage` is automatically cleared when the tab is closed.

---
