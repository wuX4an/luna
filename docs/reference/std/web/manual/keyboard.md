# Keyboard

**Description:**
Captures and handles keyboard events from Lua in the browser using WebAssembly (`syscall/js`). Allows listening to individual keys, key combinations, querying the current key state, and optionally disabling default browser behaviors.

---

### Overview

```lua
local keyboard = require("std:web").keyboard

-- Listen for any key press
keyboard.on("keydown", function(evt)
    print("Key pressed:", evt.key)
end)

-- Listen only once when a key is released
keyboard.once("keyup", function(evt)
    print("Key released:", evt.key)
end)

-- Listen for Ctrl+S combination
keyboard.onCombo({"Control", "s"}, function(evt)
    print("Save file")
    evt:preventDefault()  -- prevents default browser action
end)

-- Check if a key is currently pressed
if keyboard.isPressed("Shift") then
    print("Shift is pressed")
end

-- Disable default keybindings
keyboard.disableDefaults()                  -- disables all
keyboard.disableDefaults({"Tab", "F1"})    -- disables specific keys
```

---

### Main Functions

| Function                           | Description                                                                     |
| ---------------------------------- | ------------------------------------------------------------------------------- |
| `keyboard.on(event, callback)`     | Listen for `"keydown"` or `"keyup"` events.                                     |
| `keyboard.once(event, callback)`   | Listen for an event only once.                                                  |
| `keyboard.onCombo(keys, callback)` | Listen for simultaneous key combinations.                                       |
| `keyboard.isPressed(key)`          | Returns `true` if the specified key is currently pressed.                       |
| `keyboard.disableDefaults(keys?)`  | Disables default browser behavior for all keys or for a specific table of keys. |

---

### Events

Callbacks receive an `evt` object with the following:

```lua
evt.key             -- pressed key (e.g., "a", "Shift", "Control")
evt.code            -- physical key code (e.g., "KeyA", "ShiftLeft")
evt.ctrlKey         -- boolean, true if Ctrl is pressed
evt.shiftKey        -- boolean, true if Shift is pressed
evt.altKey          -- boolean, true if Alt is pressed
evt.metaKey         -- boolean, true if Meta/Command is pressed
evt.preventDefault()-- function to prevent default browser action
evt.stopPropagation()-- function to stop event propagation
```

**Example:**

```lua
keyboard.on("keydown", function(evt)
    print("Key:", evt.key, "Ctrl pressed:", evt.ctrlKey)
end)
```

---

### Key Combinations

- `keyboard.onCombo(keys, callback)` detects simultaneous key presses.
- `keys` is a table of key names:

```lua
keyboard.onCombo({"Shift", "Alt", "P"}, function(evt)
    print("Combination triggered")
end)
```

- Callback fires only when **all specified keys are pressed at the same time**.

---

### Disable Default Keybindings

- `keyboard.disableDefaults()` → blocks all keys from triggering default browser actions.
- `keyboard.disableDefaults({"Tab", "F1"})` → blocks only specified keys.

```lua
keyboard.disableDefaults()                  -- global block
keyboard.disableDefaults({"ArrowUp", "ArrowDown"}) -- selective block
```

---

### Devnote

- Internally maintains a map of pressed keys for `onCombo` and `isPressed`.
- Works in modern browsers with WebAssembly and `syscall/js`.
- Supports only `keydown` and `keyup` events; no legacy `keypress`.
- Default keybinding disabling affects **only events intercepted by this module**, not external scripts or extensions.

---
