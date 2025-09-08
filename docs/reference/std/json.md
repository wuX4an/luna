# std:json

The `std:json` module provides functions to **encode Lua tables into JSON strings** and **decode JSON strings into Lua tables**.
It is useful for data interchange, API communication, and configuration parsing.

---

## Functions

### `json.encode(tbl: table): string`

Encodes a Lua table into a JSON string. Nested tables and arrays are supported.

- **Parameters**:
  - `tbl` _(table)_: the Lua table to encode.

- **Returns**:
  - _(string)_: the JSON representation of the table.

**Example:**

```lua
local json = require("std:json")

local tbl = { name = "Alice", age = 30, active = true }
local str = json.encode(tbl)
print(str)
-- {"name":"Alice","age":30,"active":true}
```

---

### `json.decode(str: string): table`

Decodes a JSON string into a Lua table. Nested objects and arrays are converted recursively.

- **Parameters**:
  - `str` _(string)_: the JSON string to decode.

- **Returns**:
  - _(table)_: the Lua table representation of the JSON string.

**Example:**

```lua
local json = require("std:json")

local str = '{"users":[{"name":"Alice"},{"name":"Bob"}],"active":true}'
local tbl = json.decode(str)

print(tbl.active)          -- true
print(tbl.users[1].name)   -- Alice
print(tbl.users[2].name)   -- Bob
```

---

### Notes

- Arrays in JSON are converted to Lua tables with **1-based integer keys**.
- Nested objects and arrays are **recursively converted**.
- Invalid JSON strings will raise an error.

**Example of error handling:**

```lua
local json = require("std:json")

local ok, err = pcall(function()
    local tbl = json.decode('{"invalid_json": }')
end)

print(ok, err)
-- false    failed to decode JSON string: ...
```

---
