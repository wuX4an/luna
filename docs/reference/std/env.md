# std:env

The `std:env` module provides functions to **load, read, and modify environment variables** from Lua.
It supports loading variables from `.env` files, getting values, and setting new variables in memory and in the system environment.

---

## Functions

### `env.load(filename: string): boolean`

Loads variables from a `.env` file into memory and the system environment.

- **Parameters**:
  - `filename` _(string)_: path to the `.env` file.

- **Returns**:
  - `true` if the file was loaded successfully.

**Example:**

```lua
local env = require("std:env")

-- Load variables from .env file
local ok = env.load(".env")
if ok then
    print("Variables loaded successfully")
end
```

---

### `env.get(key: string): string?`

Retrieves the value of a loaded environment variable.

- **Parameters**:
  - `key` _(string)_: the name of the variable.

- **Returns**:
  - _(string?)_ the variable value if it exists, or `nil` otherwise.

**Example:**

```lua
local env = require("std:env")

-- Get API key
local api_key = env.get("API_KEY")
print("API_KEY:", api_key)
-- API_KEY: 1234567890abcdef
```

---

### `env.set(key: string, value: string): boolean`

Sets or updates an environment variable **in memory and in the system environment**.

- **Parameters**:
  - `key` _(string)_: the name of the variable.
  - `value` _(string)_: the value to assign.

- **Returns**:
  - `true` on success.

> This does **not** modify the `.env` file on disk; the change exists only during program execution.

**Example:**

```lua
local env = require("std:env")

-- Define or update a variable
env.set("DEBUG", "true")
print("DEBUG:", env.get("DEBUG"))
-- DEBUG: true
```

---

## Notes

- Use `.env` files to manage configuration variables for your application.
- `load()` automatically updates the system environment (`os.setenv`).
- `set()` is JIT: changes last only during the program runtime. For persistent changes, you need to modify the `.env` file manually.

---
