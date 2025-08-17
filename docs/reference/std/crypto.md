# std:crypto

The `std:crypto` module provides functions for **hashing, HMAC, UUID generation, and cryptographically secure random bytes**. It is useful for security, authentication, and unique identifier generation.

---

## Functions

### `crypto.hash(algorithm: string, data: string): string`

Computes the hash of `data` using the specified `algorithm`.

- **Parameters**:
  - `algorithm` _(string)_: Hash algorithm (`"md5"`, `"sha1"`, `"sha256"`)
  - `data` _(string)_: The input string to hash

- **Returns**:
  - _(string)_ Hexadecimal representation of the hash

**Example:**

```lua
local crypto = require("std:crypto")
print(crypto.hash("sha256", "hello world"))
-- "b94d27b9934d3e08a52e52d7da7dabfac484efe3..."
```

---

### `crypto.hmac(algorithm: string, key: string, data: string): string`

Generates a **HMAC** for `data` using `key` and the specified `algorithm`.

- **Parameters**:
  - `algorithm` _(string)_: HMAC algorithm (`"sha1"`, `"sha256"`)
  - `key` _(string)_: Secret key for HMAC
  - `data` _(string)_: Input string to sign

- **Returns**:
  - _(string)_ Hexadecimal representation of the HMAC

**Example:**

```lua
local crypto = require("std:crypto")
print(crypto.hmac("sha1", "secret", "message"))
-- "0caf649feee4953d87bf903ac1176c45e028df16"
```

---

### `crypto.random_bytes(n: number): string`

Generates `n` cryptographically secure random bytes.

- **Parameters**:
  - `n` _(number)_: Number of bytes to generate

- **Returns**:
  - _(string)_ Random bytes as a binary string

**Example:**

```lua
local crypto = require("std:crypto")
local rb = crypto.random_bytes(16)
print(rb)
-- (hex example) "\x7f\x3a\x9d\x2c\x4b\x12\x88\xef\x90\xab\xcd\x34\x56\x78\x9a\xbc"
```

---

### `crypto.uuid(): string`

Generates a **UUID v4** string.

- **Returns**:
  - _(string)_ UUID in standard string format

**Example:**

```lua
local crypto = require("std:crypto")
print(crypto.uuid())
-- "b0d8fa96-fc6b-43b2-bb02-55cfeaeb81d7"
```

---

## Supported Algorithms

| Algorithm | Description                       | Output Length |
| --------- | --------------------------------- | ------------- |
| `md5`     | Message-Digest Algorithm 5        | 16 bytes      |
| `sha1`    | Secure Hash Algorithm 1           | 20 bytes      |
| `sha256`  | Secure Hash Algorithm 2 (256-bit) | 32 bytes      |

**Notes:**

- All outputs are in **hexadecimal format**.
- `md5` is fast but **not recommended for cryptographic security**.
- `sha1` is more secure than `md5` but considered **weak for new systems**.
- `sha256` is the recommended default for most **secure hashing needs**.
