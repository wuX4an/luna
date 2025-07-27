package http

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// ServerProposal imprime la propuesta de configuración del servidor HTTP
func ServerProposal(L *lua.LState) int {
	fmt.Print(`[http.server proposal]

local http = require("http")

-- Load sublibs
local server  = http.server
local Config  = server.Config
local Route   = server.Route
local Server  = server.Server

-- Config
local config = Config.new({
  host = "0.0.0.0",
  port = 8080,
  tls = {
    cert = "./cert.pem",
    key  = "./cert.key"
  },
  log = true,
  static = {
    dir  = "static",
    path = "/static"
  }
})

-- Routes
local routes = {
  Route.new("GET", "/", function(req, res)
    res:render("index.html", { name = "Juan" })
  end),

  Route.new("GET", "/hello/:name", function(req, res)
    local name = req:param("name")
    res:send("Hello " .. name)
  end),

  Route.new("GET", "/bye", function(req, res)
    local name = req:query("name")
    res:send(name and "Bye " .. name or "Bye world")
  end),

  Route.new("GET", "/api", function(req, res)
    res:json({ received = "hola" })
  end),

  Route.new("*", "*", function(req, res)
    res:status(404):send("Not found")
  end)
}

-- Crear e iniciar servidor
local srv = Server.new(config, routes)
srv:listen()
`)
	return 0
}
