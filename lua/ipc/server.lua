local ipc = require("std:ipc")

-- servidor
local server = ipc.server({ path = "/tmp/misocket.sock" })
print("Iniciando servidor IPC...")
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
end)
