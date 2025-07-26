-- TIME

print("\n=== TEST TIME ===")
-- Aquí podrías añadir funciones de time para probar, si tienes disponibles


-- RANDOM
local random = require("std:random")

print("\n=== TEST RANDOM ===")

local items = { "manzana", "banana", "pera", "naranja" }
local weights = {50, 2, 50, 2}

print("random.choice con pesos:")
print("Elegido:", random.choice(items, weights))

print("\nrandom.int:")
print(random.int(1, 10, 2))      -- Número entero entre 1 y 10, con paso 2

print("\nrandom.float:")
print(random.float(1, 5))        -- Float entre 1 y 5 (default precisión)
print(random.float(0, 1, 4))     -- Float entre 0 y 1 con 4 decimales
print(random.float(10.5, 99.9, 0)) -- Float entre 10.5 y 99.9 sin decimales (entero)


print("\nrandom.take:")
local fruits = {"manzana", "banana", "pera", "naranja", "kiwi"}
local taken = random.take(fruits, 3)
for i, v in ipairs(taken) do
    print(i, v)
end


-- TABLEX
local tablex = require("std:tablex")

print("\n=== TEST TABLEX ===")

print("tablex.raw de taken:")
print(tablex.raw(taken))

local map = tablex.map({1, 2, 3}, function(v) return v * 2 end)
print("tablex.map (x2):")
print(tablex.raw(map))

local filter = tablex.filter({1, 2, 3, 4}, function(v) return v % 2 == 0 end)
print("tablex.filter (pares):")
print(tablex.raw(filter))


-- SQLITE
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
  if not row then break end
  print(string.format("ID=%d, Nombre=%s", row.id, row.name))
end
rows:close()
local countRows, err = db:query("SELECT COUNT(*) AS total FROM users")
assert(countRows, err)
local countRow = countRows:next()
countRows:close()

print("\nTotal de usuarios en la tabla:", countRow.total)
db:close()

-- TEST
print("\n=== TEST ===")

local test = require("std:test")

-- test 1
test.run("simple test", function()
  return test.expect(5):to_equal(4)
end)

-- test 2
test("inline test", function()
  return test.expect(42):to_equal(41)
end)

-- test 3
test.expect(2):to_equal(2)

-- test 4
test.describe("basic math", function()
	test.run("addition", function()
	  return test.expect(1 + 2):to_equal(3)
	end)

	test.run("multiplication", function()
		return test.expect(2 * 3):to_equal(5)
	end)
end)
