local time = require("std:time")

-- Obtener timestamp actual (segundos Unix)
local now = time.now()
print("Timestamp actual:", now)

-- Formatear fecha y hora usando formato estilo strftime
local fecha = time.format("%Y-%m-%d %H:%M:%S", now)
print("Fecha formateada:", fecha)

-- Dormir 1 segundo (1000 ms)
print("Durmiendo 1 segundo...")
time.sleep(1000)
print("¡Despierto!")

-- Más ejemplos de formatos compatibles:
print(time.format("%A, %B %d, %Y", now)) -- Ej: "Saturday, August 10, 2025"

local formatted = time.format("%Y-%m-%dT%H:%M:%S %p %A (%a) %B (%b) %d day, %I hour, %M minute, %S second %Z %z", now)
print(formatted)
