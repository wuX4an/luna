local tablex = require("std:tablex")

print("\n=== TEST TABLEX ===")

local items = { "manzana", "banana", "pera", "naranja" }
print("tablex.raw de taken:")
print(tablex.raw(items))

local map = tablex.map({ 1, 2, 3 }, function(v)
	return v * 2
end)
print("tablex.map (x2):")
print(tablex.raw(map))

local filter = tablex.filter({ 1, 2, 3, 4 }, function(v)
	return v % 2 == 0
end)
print("tablex.filter (pares):")
print(tablex.raw(filter))
