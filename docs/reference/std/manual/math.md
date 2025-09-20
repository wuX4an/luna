# std:math

The `std:math` module provides common mathematical functions for Lua.
It includes **trigonometry with unit specification**, exponentiation, logarithms, square roots, and factorials.

---

## Functions

### `math.cos`

`(x: number, unit: 'r'|'g'): number`

Returns the cosine of `x`.

- **Parameters**:
  - `x` _(number)_: the angle.
  - `unit` _(string)_: `'r'` for radians, `'g'` for degrees.

- **Returns**:
  - _(number)_ the cosine of `x`.

**Example:**

```lua
local math = require("std:math")
print(math.cos(math.pi / 3, 'r'))  -- 0.5
print(math.cos(60, 'g'))           -- 0.5
```

---

### `math.sin`

`(x: number, unit: 'r'|'g'): number`

Returns the sine of `x`.

- **Parameters**:
  - `x` _(number)_: the angle.
  - `unit` _(string)_: `'r'` for radians, `'g'` for degrees.

- **Returns**:
  - _(number)_ the sine of `x`.

**Example:**

```lua
local math = require("std:math")
print(math.sin(math.pi / 2, 'r'))  -- 1
print(math.sin(90, 'g'))           -- 1
```

---

### `math.tan`

`(x: number, unit: 'r'|'g'): number`

Returns the tangent of `x`.

- **Parameters**:
  - `x` _(number)_: the angle.
  - `unit` _(string)_: `'r'` for radians, `'g'` for degrees.

- **Returns**:
  - _(number)_ the tangent of `x`.

**Example:**

```lua
local math = require("std:math")
print(math.tan(math.pi / 4, 'r'))  -- 1
print(math.tan(45, 'g'))           -- 1
```

---

### `math.sqrt`

`(x: number): number`

Returns the square root of `x`.

- **Parameters**:
  - `x` _(number)_: value to compute the square root of.

- **Returns**:
  - _(number)_ the square root of `x`.

**Example:**

```lua
local math = require("std:math")
print(math.sqrt(16))  -- 4
```

---

### `math.pow`

`(x: number, y: number): number`

Returns `x` raised to the power of `y`.

- **Parameters**:
  - `x` _(number)_: base.
  - `y` _(number)_: exponent.

- **Returns**:
  - _(number)_ `x`^`y`.

**Example:**

```lua
local math = require("std:math")
print(math.pow(2, 3))  -- 8
```

---

### `math.fact`

`(n: integer): integer`

Returns the factorial of `n`.

- **Parameters**:
  - `n` _(integer)_: non-negative integer.

- **Returns**:
  - _(integer)_ `n!`.

**Example:**

```lua
local math = require("std:math")
print(math.fact(5))  -- 120
```

---

### `math.log`

`(x: number): number`

Returns the natural logarithm (base e) of `x`.

- **Parameters**:
  - `x` _(number)_: value to compute the logarithm of.

- **Returns**:
  - _(number)_ the natural logarithm of `x`.

**Example:**

```lua
local math = require("std:math")
print(math.log(1))  -- 0
```

---

## Constants

| Constant | Value                  | Description                                |
| -------- | ---------------------- | ------------------------------------------ |
| `pi`     | 3.14159265358979323846 | Ratio of circle circumference to diameter  |
| `e`      | 2.71828182845904523536 | Euler's number, base of natural logarithms |
| `tau`    | 6.28318530717958647692 | 2π, full circle in radians                 |
| `phi`    | 1.61803398874989484820 | Golden ratio                               |
| `G`      | 6.67430×10⁻¹¹          | Gravitational constant, m³·kg⁻¹·s⁻²        |

---
