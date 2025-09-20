# std:random

The `std:random` module provides functions for generating random numbers, selecting random items, and weighted randomness.  
It is useful for games, simulations, testing, and any scenario that requires non-deterministic values.

---

## Functions

### `random.choice`

`(items: table, weights: table?): any`

Selects a random item from a list. Optionally, you can provide weights to bias the selection.

- **Parameters**:
  - `items` _(table)_: list of items to choose from.
  - `weights` _(table, optional)_: weights corresponding to each item.
- **Returns**:
  - A single randomly selected item.

**Example:**

```lua
local random = require("std:random")

local items = { "apple", "banana", "pear", "orange" }
local weights = { 50, 2, 50, 2 }

print("Chosen item:", random.choice(items, weights))
```

---

### `random.int`

`(min: integer, max: integer, step: integer?): integer`

Generates a random integer within a specified range and optional step.

- **Parameters:**
  - min _(integer)_: minimum value (inclusive).
  - max _(integer)_: maximum value (inclusive).
  - step _(integer, optional)_: step increment. Default is 1.
- Returns:
  - A random integer.

**Example:**

```lua
local random = require("std:random")

print(random.int(1, 10, 2)) -- random integer between 1 and 10 with step 2
```

---

### `random.float`

`(min: number, max: number, precision: integer?): number`

Generates a random float number within a specified range and optional decimal precision.

- **Parameters:**
  - `min` (number): minimum value.
  - `max` (number): maximum value.
  - `precision` (integer, optional): number of decimal places. Default is 2.
- **Returns:**
  - A random float.

**Example:**

```lua
local random = require("std:random")

print(random.float(1, 5))         -- float between 1 and 5
print(random.float(0, 1, 4))      -- float between 0 and 1 with 4 decimals
print(random.float(10.5, 99.9, 0))  -- float between 10.5 and 99.9 as integer
```

---

### `random.take`

`(items: table, n: integer): table`

Returns a table containing n randomly selected items from the input table.

- **Parameters:**
  - `items` _(table)_: list of items to select from.
  - `n` _(integer)_: number of items to take.

- **Returns:**
  - A table with `n` randomly selected items.

**Example:**

```lua
local random = require("std:random")

local fruits = { "apple", "banana", "pear", "orange", "kiwi" }
local taken = random.take(fruits, 3)

for i, v in ipairs(taken) do
  print(i, v)
-- 1   banana
-- 2   kiwi
-- 3   apple
end
```

---

## Notes:

- `choice` respects the `weights` array if provided; otherwise, selection is uniform.
- `int` and `float` ranges are inclusive.
- `take` may return fewer items than requested if `n` > `#items`.

---
