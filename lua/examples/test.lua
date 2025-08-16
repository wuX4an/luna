local test = require("std:test")

print("\n=== TEST ===")

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

-- test 5
test.describe("Multiplicación parametrizada", function()
	test.each({
		{ a = 2, b = 3, esperado = 7, name = "2 * 3 = 6" },
		{ a = 5, b = 5, esperado = 25, name = "5 * 5 = 25" },
	}, function(caso)
		local res = caso.a * caso.b
		return test.expect(res):to_equal(caso.esperado)
	end)
end)
