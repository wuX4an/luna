# std:base64

The `std:base64` module provides functions to encode and decode data in Base64 format.
It is useful for binary manipulation, transmitting data in plain text, and ensuring compatibility with APIs.

---

## Functions:

### `base64.encode`

`(input: string): string`

Encodes a string into its Base64 representation.

- **Parameters**:
  - `input` _(string)_: the string to encode.
- **Returns**:
  - _(string)_ the Base64-encoded string.

**Example:**

```lua
local base64 = require("std:base64")

print(base64.encode("Hello Luna"))
-- "SGVsbG8gTHVuYQ=="
```

---

### `base64.decode`

`(input: string): string`

Decodes a Base64 string back into its original form.

- **Parameters**:
  - `input` _(string)_: the string to decode.
- **Returns**:
  - _(string)_ the Base64-decoded string.

**Example:**

```lua
local base64 = require("std:base64")

print(base64.decode("Tm90Y2g="))
-- "Notch"
```

---
