local DOM = require("std:web").dom

-- Obtener root y limpiar contenido
local root = DOM.root()
root:clear()

-- Crear contenedor
local container = DOM.div({ id = "app", class = "main-container" }, "")
root:append(container)
container:setStyle({
	display = "flex",
	flexDirection = "column",
	gap = "10px",
	width = "300px",
	margin = "50px auto",
	padding = "20px",
	backgroundColor = "#f9f9f9",
	borderRadius = "10px",
})

-- Make Message
local msg = DOM.div({ id = "msg" }, "Hello World!")
container:append(msg)
msg:setStyle({ padding = "10px", backgroundColor = "#fff", border = "1px solid #ccc" })

-- Make Button
local btn = DOM.button({ id = "btn" }, "Change Message")
container:append(btn)
btn:setStyle({ padding = "10px", cursor = "pointer" })

-- Event on click: Change Message
btn:on("click", function(evt)
	msg:setText("¡Updated Message!")
end)
