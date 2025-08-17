local crypto = require("std:crypto")

-- Hash
print(crypto.hash("sha256", "hello world"))
-- e.g.: b94d27b9934d3e08a52e52d7da7dabfade0a6b5e...

-- HMAC
print(crypto.hmac("sha1", "secret", "message"))
-- e.g.: 5f4c5f9d0e1a8e7b9a3c0c1f2d3e4f56789abcd0

-- Random bytes
local rb = crypto.random_bytes(16)
print(rb) -- binary string (can convert to hex if needed)

-- UUID
print(crypto.uuid())
-- e.g.: 550e8400-e29b-41d4-a716-446655440000
