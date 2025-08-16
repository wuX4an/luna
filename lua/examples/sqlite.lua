local sqlite = require("std:sqlite")

print("=== TEST SQLITE EN MEMORIA ===")

-- Abrir base de datos en memoria
local db, err = sqlite.open(":memory:")
assert(db, err)
local ok, err = assert(db:exec([[
    CREATE TABLE users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL
    );
]]))
assert(ok, err)

-- Crear tabla e insertar datos dentro de una transacción
local ok, err = db:transaction(function(tx)
	assert(tx:exec("INSERT INTO users (name) VALUES ('Ana')"))
	assert(tx:exec("INSERT INTO users (name) VALUES ('Luis')"))
	assert(tx:exec("INSERT INTO users (name) VALUES ('Carlos')"))
end)
assert(ok, err)

-- Consultar y mostrar usuarios
local rows, err = db:query("SELECT id, name FROM users ORDER BY id")
assert(rows, err)

print("Usuarios en la tabla:")
while true do
	local row = rows:next()
	if not row then
		break
	end
	print(string.format("ID=%d, Nombre=%s", row.id, row.name))
end
rows:close()
local countRows, err = db:query("SELECT COUNT(*) AS total FROM users")
assert(countRows, err)
local countRow = countRows:next()
countRows:close()

print("\nTotal de usuarios en la tabla:", countRow.total)
db:close()
