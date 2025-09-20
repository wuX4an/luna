# std:test

The `std:test` module provides a lightweight testing framework for Lua.
It supports assertions, named tests, test grouping with `describe`, and parameterized tests with `each`.

---

## Recommended way to define tests

You can define tests in a way that they **only execute when requested**, using the global `__TESTS__` table.

```lua
-- Define tests
__TESTS__ = __TESTS__ or {}
table.insert(__TESTS__, function()
    local test = require("std:test")
    test.run("simple test", function()
        return test.expect(5):to_equal(4)
    end)
end)

-- Execute tests conditionally (optional)
if os.getenv("LUNA_RUN_TESTS") then
    for _, t in ipairs(__TESTS__ or {}) do t() end
end
```

**Advantages:**

- Tests are registered but **do not run by default**.
- You can execute all tests in a controlled environment.
- Keeps production code clean and backwards-compatible.
- Cleaned by the builder.

---

## How to run the tests

1. **Set the environment variable and run your script:**

```bash
LUNA_RUN_TESTS=1 luna main.lua
```

2. **Or using the CLI command:**

```bash
luna test main.lua
```

> This will set `LUNA_RUN_TESTS` internally and execute all registered tests.

---

## Functions

### `test.run | test`

`(name: string, fn: function)`

Runs a named test.

- **Parameters**:
  - `name` _(string)_: the name of the test.
  - `fn` _(function)_: a function returning a test expectation.

**Example:**

```lua
local test = require("std:test")

-- Simple test
test.run("simple test", function()
    return test.expect(5):to_equal(4)  -- Expect 5 to equal 4
end)

-- Inline test
test("inline test", function()
    return test.expect(42):to_equal(41)
end)
```

---

### `test.expect`

`(value)`

Creates an expectation object to assert values.

- **Methods:**
  - `:to_equal(expected)` — asserts that the value equals `expected`.

**Example:**

```lua
-- Direct expectation
test.expect(2):to_equal(2)  -- This test passes
```

---

### `test.describe`

`(name: string, fn: function)`

Groups multiple related tests under a description.

- **Parameters:**
  - `name` _(string)_: the description for the group.
  - `fn` _(function)_: contains multiple `test.run` calls or other assertions.

**Example:**

```lua
test.describe("basic math", function()
    test.run("addition", function()
        return test.expect(1 + 2):to_equal(3)
    end)
    test.run("multiplication", function()
        return test.expect(2 * 3):to_equal(5)
    end)
end)
```

---

### `test.each`

`(cases: table, fn: function)`

Runs parameterized tests for each case in a table.

- **Parameters:**
  - `cases` _(table)_: an array of test case tables.
  - `fn` _(function)_: function receiving a case object and returning an expectation.

**Example:**

```lua
test.describe("Parameterized multiplication", function()
    test.each({
        { a = 2, b = 3, expected = 6, name = "2 * 3 = 6" },
        { a = 5, b = 5, expected = 25, name = "5 * 5 = 25" },
    }, function(case)
        local res = case.a * case.b
        return test.expect(res):to_equal(case.expected)
    end)
end)
```

---

## Complete Example

```lua
-- Register tests
__TESTS__ = __TESTS__ or {}
table.insert(__TESTS__, function()
    local test = require("std:test")
    test.run("simple test", function()
        return test.expect(5):to_equal(4)  -- Fail
    end)
    test("inline test", function()
        return test.expect(42):to_equal(41)  -- Fail
    end)
    test.expect(2):to_equal(2)  -- Pass
    test.describe("basic math", function()
        test.run("addition", function() return test.expect(1 + 2):to_equal(3) end)
        test.run("multiplication", function() return test.expect(2 * 3):to_equal(5) end)
    end)
    test.describe("Parameterized multiplication", function()
        test.each({
            { a = 2, b = 3, expected = 6, name = "2 * 3 = 6" },
            { a = 5, b = 5, expected = 25, name = "5 * 5 = 25" },
        }, function(case)
            local res = case.a * case.b
            return test.expect(res):to_equal(case.expected)
        end)
    end)
end)
```
