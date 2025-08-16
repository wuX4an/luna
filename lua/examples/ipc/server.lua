local ipc = require("std:ipc")

-- servidor
local server = ipc.server({ path = "/tmp/luna.sock" })
print("Iniciando servidor IPC...")

local function on_close(msg)
	print(msg .. " Shutdown")
end

server:start(function(conn)
	while true do
		local data = conn:recv(1024)
		if not data then
			break
		end
		print("Recibido:", data)
		conn:send("Echo: " .. data)
	end
	conn:close()
end, on_close)
