--[[
-- init.lua
local server   = require("std:http/server")
local request  = require("std:http/request")

return {
  Server = server,
  Request = request,
}
--]]
local http = require("std:http/server")

local routes = function()
	local r = http.Engine()

	-- Middleware
	r:Use(function(c, next)
		print("Request:", c.Method, c.Path)
		next()
	end)

	r:GET("/api/hello", function(c)
		c.status(http.Status.OK)
		c.header("X-Message", "Hi!")
		local html = "<h1>{{name}}</h1>"
		return c.HTML(html, {
			name = "hola",
		})
	end)

	-- Rutas simples
	r:GET("/ping", function(c)
		c.status(http.Status.OK)
		return c.String("pong")
	end)

	r:GET("/api/hello", function(c)
		c.status(http.Status.OK)
		c.header("X-Message", "Hi!")
		return c.JSON({ ok = true, message = "Hello Lua!" })
	end)

	-- Grupos de rutas
	local api = r:Group("/api")
	api:GET("/users", function(c)
		c.status(http.Status.OK)
		return c.JSON({ users = { "Alice", "Bob" } })
	end)

	local page = r:Group("/page")
	page.Static("/", "./dist" --[[Path o embed path]]) -- r.Static o c.File según tu engine

	-- Param y Query
	r:GET("/users/:id", function(c)
		local id = c.Param("id")
		local search = c.Query("q")
		return c.JSON({ id = id, q = search })
	end)

	-- Download
	r.GET("/file", function(c)
		c.Download("hono.lua") -- path or embedPath
		return c.String("hola")
	end)
	-- File
	r.GET("/file", function(c)
		return c.File("hono.lua") -- path or embedPath
	end)

	-- Get body json
	r:POST("/login", function(c)
		local data = c.BodyJSON()
		return c.JSON({ user = data.username })
	end)

	-- Cookie
	r:GET("/setcookie", function(c)
		c.SetCookie("session_id", "abc123", {
			path = "/",
			httpOnly = true,
			secure = true,
			domain = "example.com",
			maxAge = 3600,
			expire = Date(2000, 11, 24, 10, 30, 59, 900),
		})
		return c.String("Cookie set!")
	end)

	r:GET("/getcookie", function(c)
		local sid = c.Cookie("session_id")
		if not sid then
			return c.String("No cookie found")
		end
		return c.JSON({ session = sid })
	end)

	r:GET("/logout", function(c)
		c.ClearCookie("session_id")
		if c.Cookie("session_id") == nil then
			-- Redirect
			c.Redirect("/login", "Why are you here?")
		end
		return c.String("Cookie cleared!")
	end)
	-- Error Handler Global
	r:OnError(function(c, err)
		c.status(500)
		return c.JSON({ error = tostring(err) })
	end)
	-- Retornas el engine ya configurado
	return r
end

-- Luego arrancas el servidor:
local r = routes()
r:Listen(8080)
