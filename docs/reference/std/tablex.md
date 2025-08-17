# std\:tablex

The `std:tablex` module provides utility functions to manipulate Lua tables. It allows mapping, filtering, and other functional-style operations over tables.

---

## Functions:

### `tablex.raw(tbl: table): table`

Returns a shallow copy of the table suitable for printing or inspection.

- **Parameters**:
  - `tbl` _(table)_: the table to inspect.

- **Returns**:
  - _(table)_ a raw copy of the table.

**Example:**

```lua
local tablex = require("std:tablex")
local items = { "manzana", "banana", "pera", "naranja" }
print(tablex.raw(items))
-- { "manzana", "banana", "pera", "naranja" }
```

### `tablex.map(tbl: table, fn: function): table`

Applies a function to each element of the table and returns a new table with the results.

- **Parameters**:
  - `tbl` _(table)_: the input table.
  - `fn` _(function)_: the function to apply to each element.

- **Returns**:
  - _(table)_ a new table with the transformed values.

**Example:**

```lua
local tablex = require("std:tablex")
local map = tablex.map({ 1, 2, 3 }, function(v) return v * 2 end)
print(tablex.raw(map))
-- { 2, 4, 6 }
```

### `tablex.filter(tbl: table, fn: function): table`

Filters a table by applying a predicate function, returning only elements that satisfy the condition.

- **Parameters**:
  - `tbl` _(table)_: the input table.
  - `fn` _(function)_: the predicate function; should return `true` to keep an element.

- **Returns**:
  - _(table)_ a new table containing only elements for which `fn` returned `true`.

**Example:**

```lua
local tablex = require("std:tablex")
local filter = tablex.filter({ 1, 2, 3, 4 }, function(v) return v % 2 == 0 end)
print(tablex.raw(filter))
-- { 2, 4 }
```
