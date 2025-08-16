local random = require("std:random")

print("\n=== TEST RANDOM ===")

local items = { "manzana", "banana", "pera", "naranja" }
local weights = { 50, 2, 50, 2 }

print("random.choice con pesos:")
print("Elegido:", random.choice(items, weights))

print("\nrandom.int:")
print(random.int(1, 10, 2)) -- Número entero entre 1 y 10, con paso 2

print("\nrandom.float:")
print(random.float(1, 5)) -- Float entre 1 y 5 (default precisión)
print(random.float(0, 1, 4)) -- Float entre 0 y 1 con 4 decimales
print(random.float(10.5, 99.9, 0)) -- Float entre 10.5 y 99.9 sin decimales (entero)

print("\nrandom.take:")
local fruits = { "manzana", "banana", "pera", "naranja", "kiwi" }
local taken = random.take(fruits, 3)
for i, v in ipairs(taken) do
	print(i, v)
end
