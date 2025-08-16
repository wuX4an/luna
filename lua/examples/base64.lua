local b64 = require("std:base64")

local string = "Hello world"
local encode = b64.encode(string)

print(encode)

print(b64.decode(encode))
