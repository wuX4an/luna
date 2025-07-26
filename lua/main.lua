-- TIME
local time = require("time")
-- RANDOM
local random = require("random")
local items = { "manzana", "banana", "pera", "naranja" }
local weights = {50, 2, 50, 2}

print(random.choice(items, weights))  -- Esto debe funcionar si `random.choice` está bien definido en Go
print(random.int(1,10,2))
print(random.float(1, 5))          --> 3.14 (por ejemplo)
print(random.float(0, 1, 4))       --> 0.1234
print(random.float(10.5, 99.9, 0)) --> 42 (número entero, sin decimales)
local fruits = {"manzana", "banana", "pera", "naranja", "kiwi"}
local items = { "manzana", "banana", "pera", "naranja" }
local weights = {50, 2, 50, 2}

print(random.choice(items, weights))  -- Esto debe funcionar si `random.choice` está bien definido en Go
print(random.int(1,10,2))
local fruits = {"manzana", "banana", "pera", "naranja", "kiwi"}
local taken = random.take(fruits, 3)
local taken = random.take(fruits, 3)

-- TABLEX 
local tablex = require("tablex")
print(tablex.raw(taken))
local map = tablex.map({1, 2, 3}, function(v) return v * 2 end)
print(tablex.raw(map))
local filter = tablex.filter({1, 2, 3, 4}, function(v) return v % 2 == 0 end)
print(tablex.raw(filter))
