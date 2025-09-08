# Proposal: Modular HTTP API in Lua (Hono-style)

## 1. Introduction

The goal of this project is to develop a **modular and extensible HTTP API in Lua**, inspired by modern frameworks like Hono or Express.js. This API allows developers to:

- Easily define GET/POST routes.
- Implement **middleware** for logging, authentication, or validation.
- Handle **route groups** and static routes.
- Manage **cookies**, file downloads, static files, and redirects.
- Return responses in **JSON, HTML, or plain text**.
- Handle errors globally using an **Error Handler**.

This project targets **lightweight Lua environments or embedded servers**, offering explicit control over request and response handling.

---

## 2. Code Structure

### 2.1. Server Initialization

```lua
local http = require("std:http/server")

local routes = function()
    local r = http.Engine()
```

- `http.Engine()` creates a **routing and middleware engine**, similar to `express()` in Node.js.
- Returns an object `r` exposing methods: `GET`, `POST`, `Use`, `Group`, `OnError`.

---

### 2.2. Middleware

```lua
r:Use(function(c, next)
    print("Request:", c.Method, c.Path)
    next()
end)
```

- Executes logic before each route.
- Examples: logging, authentication, parameter validation.
- Call `next()` to continue to the route handler.

---

### 2.3. Simple Routes

```lua
r:GET("/ping", function(c)
    c.status(http.Status.OK)
    return c.String("pong")
end)

r:GET("/api/hello", function(c)
    c.status(http.Status.OK)
    c.header("X-Message", "Hi!")
    return c.JSON({ ok = true, message = "Hello Lua!" })
end)
```

- Supports returning **JSON**, **plain text**, or dynamic content.
- Headers and HTTP status codes can be configured per route.

---

### 2.4. Route Groups

```lua
local api = r:Group("/api")
api:GET("/users", function(c)
    c.status(http.Status.OK)
    return c.JSON({ users = { "Alice", "Bob" } })
end)

local page = r:Group("/page")
page.Static("/", "./dist") -- Serve static files
```

- Organize routes under common prefixes (`/api`, `/page`).
- Serve static content with `Static()`.

---

### 2.5. Path Parameters and Query Strings

```lua
r:GET("/users/:id", function(c)
    local id = c.Param("id")
    local search = c.Query("q")
    return c.JSON({ id = id, q = search })
end)
```

- `Param("id")` retrieves route parameters.
- `Query("q")` retrieves query string parameters.

---

### 2.6. File Downloads and Static Files

```lua
r:GET("/file", function(c)
    c.Download("hono.lua")
    return c.String("hola")
end)

r:GET("/file", function(c)
    return c.File("hono.lua")
end)
```

- `Download()` forces file download.
- `File()` returns the file as a normal response.

---

### 2.7. JSON Body Handling

```lua
r:POST("/login", function(c)
    local data = c.BodyJSON()
    return c.JSON({ user = data.username })
end)
```

- Automatically parses JSON body into a Lua table.
- Simplifies handling of POST requests with JSON payloads.

---

### 2.8. Cookies and Sessions

```lua
r:GET("/setcookie", function(c)
    c.SetCookie("session_id", "abc123", { path="/", httpOnly=true, secure=true, maxAge=3600 })
    return c.String("Cookie set!")
end)

r:GET("/getcookie", function(c)
    local sid = c.Cookie("session_id")
    if not sid then return c.String("No cookie found") end
    return c.JSON({ session = sid })
end)

r:GET("/logout", function(c)
    c.ClearCookie("session_id")
    c.Redirect("/login", "Why are you here?")
    return c.String("Cookie cleared!")
end)
```

- Complete cookie management: creation, retrieval, clearing, and redirecting.

---

### 2.9. Global Error Handler

```lua
r:OnError(function(c, err)
    c.status(500)
    return c.JSON({ error = tostring(err) })
end)
```

- Catches uncaught errors globally.
- Centralizes logging and error reporting.

---

### 2.10. Server Start

```lua
return r

-- Start the server
local r = routes()
r:Listen(8080)
```

- Returns the configured engine.
- Server listens on port 8080.

---

## 3. Advantages

1. **Modularity**: Independent middleware, groups, and route handlers.
2. **Clarity**: Declarative and easy-to-read code.
3. **Extensibility**: Add new functionality (WebSockets, Auth) without modifying core.
4. **Lightweight**: Runs in embedded Lua environments.
5. **Developer-friendly**: Inspired by modern frameworks for easier adoption.

---

## 4. Next Steps

- Integrate **dynamic HTML templates**.
- Improve **session management and cookie security**.
- Add **robust support for query and body parameters**.
- Provide **complete developer documentation** with examples.
