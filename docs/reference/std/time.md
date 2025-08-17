# std:time

The `std:time` module provides functions for working with timestamps, formatting dates and times, and controlling delays.
It is useful for logging, scheduling, time calculations, and formatting dates for display or storage.

---

## Functions

### `time.now(): integer`

Returns the current Unix timestamp (seconds since January 1, 1970).

- **Parameters**: None
- **Returns**:
  - _(integer)_ the current timestamp in seconds.

**Example:**

```lua
local time = require("std:time")

local now = time.now()
print("Current timestamp:", now)
```

---

### `time.format(format_string: string, timestamp: integer): string`

Formats a given timestamp into a human-readable string according to the specified format.
Format strings follow a `strftime`-style syntax.

- **Parameters**:
  - `format_string` _(string)_: formatting pattern (like `%Y-%m-%d %H:%M:%S`).
  - `timestamp` _(integer)_: Unix timestamp to format.

- **Returns**:
  - _(string)_ formatted date and time.

**Example:**

```lua
local time = require("std:time")

local now = time.now()
local formatted = time.format("%A, %B %d, %Y", now)
print(formatted)
-- "Saturday, August 10, 2025"
```

---

### `time.sleep(milliseconds: integer)`

Pauses execution for a specified amount of time.

- **Parameters**:
  - `milliseconds` _(integer)_: duration to sleep in milliseconds.

- **Returns**: None

**Example:**

```lua
local time = require("std:time")

print("Sleeping for 1 second...")
time.sleep(1000)
print("Awake!")
```

---

## Notes

- All timestamps are in **seconds** since Unix epoch.
- The `format` function supports standard `strftime` tokens.
- `sleep` may not be perfectly precise for very short durations due to system scheduling.

## Supported Format Tokens

| Token | Example Value | Meaning                            |
| ----- | ------------- | ---------------------------------- |
| `%Y`  | 2006          | Full year (4 digits)               |
| `%y`  | 06            | Year (2 digits)                    |
| `%m`  | 01            | Month (01-12)                      |
| `%d`  | 02            | Day of month (01-31)               |
| `%H`  | 15            | Hour (24-hour)                     |
| `%I`  | 03            | Hour (12-hour)                     |
| `%M`  | 04            | Minute                             |
| `%S`  | 05            | Second                             |
| `%p`  | PM            | AM/PM                              |
| `%z`  | **-0500**     | Timezone offset                    |
| `%Z`  | **Local**     | Timezone abbreviation              |
| `%A`  | **Sunday**    | Full weekday name (manual)         |
| `%a`  | **Sun**       | Abbreviated weekday name           |
| `%B`  | **August**    | Full month name                    |
| `%b`  | **Aug**       | Abbreviated month name             |
| `%F`  | 2006-01-02    | Equivalent to `%Y-%m-%d`           |
| `%T`  | 15:04:05      | Equivalent to `%H:%M:%S`           |
| `%r`  | 03:04:05 PM   | 12-hour clock time                 |
| `%R`  | 15:04         | 24-hour clock time without seconds |

---
