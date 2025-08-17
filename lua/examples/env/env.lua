local env = require("std:env")

-- Cargar variables de archivo
env.load("lua/examples/env/env")

-- Obtener variable
local world = env.get("HELLO")
print("HELLO", world)

-- Definir o actualizar variable
env.set("HELLO", "Tralalero")
print("HELLO:", env.get("HELLO"))
