# Clipboard

**Description:**
Provides access to the browser clipboard from Lua using `await`.

---

### Functions

```lua
clipboard.write(text)    -- Copies text to the clipboard (await-compatible)
clipboard.read()         -- Returns the current clipboard text (await-compatible)
clipboard.clear()        -- Clears the clipboard (not supported in all browsers)
```

```lua
local clipboard = require("std:web").clipboard
local await = require("std:web").await

-- Read text
local text = await(clipboard.read) -- passes the read function, which accepts a callback
print(text)

-- Copy text
clipboard.write("Hello") -- Write must also accept a callback

-- Clear clipboard
clipboard.clear()
```

---

### Notes

- `write` and `read` depend on browser permissions; some require a user-triggered event (click, keypress).
- `clear` is not supported in all browsers.
- All functions are now **await-compatible**, simplifying use in coroutines and avoiding explicit callbacks.

---
