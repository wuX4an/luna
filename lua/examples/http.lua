local http = require("std:http")

print("\n=== HTTP ===")

local res = http.request("GET", "https://ifconfig.me", {
	headers = { Accept = "application/json" },
	-- query = { q = "lua" },
	-- timeout = 5000,
})

if res then
	print(res.body)
else
	print("request failed")
end

-- Server

http.server()
