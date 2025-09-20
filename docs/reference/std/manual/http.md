# std:http

The `std:http` module provides a simple interface to perform HTTP requests in Lua.
It supports common HTTP methods (`GET`, `POST`, etc.), headers, query parameters, optional timeout, and sending request bodies (JSON or form-urlencoded).

---

## Functions

### `http.request`

`(method: string, url: string, opts: table?): table?`

Performs an HTTP request.

- **Parameters**:
  - `method` _(string)_: HTTP method (`"GET"`, `"POST"`, etc.).
  - `url` _(string)_: the full URL to request.
  - `opts` _(table?)_: optional table with fields:
    - `headers` _(table)_: key-value table of HTTP headers.
    - `query` _(table)_: key-value table of query parameters.
    - `timeout` _(number)_: request timeout in milliseconds.
    - `body` _(string)_: request body (for POST/PUT requests).

- **Returns**:
  - `table?` with the response:
    - `status` — HTTP status code.
    - `body` — response body as a string.

  - Returns `nil` and an error message if the request fails.

---

## Examples

### GET Request

```lua
local http = require("std:http")

local res = http.request("GET", "https://ifconfig.me", {
    headers = { Accept = "application/json" },
    query = { q = "lua" },
    timeout = 5000,
})

if res then
    print("=== Response Body ===")
    print(res.body)      -- Prints the response body
    print("Status:", res.status)
else
    print("Request failed")  -- Prints if request could not be completed
end
```

---

### POST Request (Form URL-encoded)

```lua
local http = require("std:http")

local res = http.request("POST", "https://httpbin.org/post", {
    headers = { ["Content-Type"] = "application/x-www-form-urlencoded" },
    body = "foo=bar&baz=42",  -- Form data
})

if res then
    print("=== Response Body ===")
    print(res.body)      -- Should show "form": {"foo":"bar","baz":"42"} in JSON
    print("Status:", res.status)
else
    print("Request failed")
end
```

---

### POST Request (JSON)

```lua
local http = require("std:http")

local res = http.request("POST", "https://httpbin.org/post", {
    headers = { ["Content-Type"] = "application/json" },
    body = '{"foo":"bar","baz":42}',
})

if res then
    print("=== Response Body ===")
    print(res.body)      -- Should show "json": {"foo":"bar","baz":42} in JSON
    print("Status:", res.status)
else
    print("Request failed")
end
```

---

## Notes

- Use the `query` table to append query parameters to the URL automatically.
- The `timeout` option prevents hanging requests.
- Headers are optional but recommended when requesting JSON or APIs that require authentication.
- For POST/PUT requests:
  - Use `"Content-Type": "application/x-www-form-urlencoded"` for form submissions.
  - Use `"Content-Type": "application/json"` for JSON payloads.

- Sending a `body` is optional; omit it for GET requests.

---
