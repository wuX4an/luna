local ipc = require("std:ipc")

-- cliente
local client = ipc.client({ path = "/tmp/misocket.sock" })
client:connect()
client:send("Hola servidor IPC")
local reply = client:recv(1024)
print("Respuesta del servidor:", reply)
client:close()
