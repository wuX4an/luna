# DOM

**Descripción:** Manipulación de elementos HTML y estructura del DOM desde Lua.
Este módulo permite crear elementos HTML, modificar sus atributos, estilos y contenido, y manejar eventos desde Lua, usando WebAssembly (`syscall/js`) para interactuar con el DOM del navegador.

---

### Overview

```lua
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

-- Crear mensaje
local msg = DOM.div({ id = "msg" }, "Hola, mundo!")
container:append(msg)
msg:setStyle({ padding = "10px", backgroundColor = "#fff", border = "1px solid #ccc" })

-- Crear botón
local btn = DOM.button({ id = "btn" }, "Cambiar mensaje")
container:append(btn)
btn:setStyle({ padding = "10px", cursor = "pointer" })

-- Evento click: cambia el mensaje
btn:on("click", function(evt)
    msg:setText("¡Mensaje actualizado!")
end)
```

### Creación de elementos

```lua
local div = DOM.div({ id = "container", class = "main" }, "Texto inicial")
local span = DOM.span({}, "Hola")
local button = DOM.button({ id = "btn1" }, "Click me")
local input = DOM.input({ type = "text", placeholder = "Escribe algo..." })
```

- **Props:** `id`, `class`, `style` u otros atributos HTML.
- **children:** opcional, puede ser texto (`string`) o elementos hijos (`userdata`).
- **Referencia:** `DOM.div`, `DOM.span`, `DOM.button`, `DOM.input`, o `DOM.create(tag, props, children)` para cualquier tag HTML.

---

### Métodos de elementos

#### Atributos y propiedades

| Método                           | Descripción                          |
| -------------------------------- | ------------------------------------ |
| `elem:setAttr(key, value)`       | Modifica un atributo HTML.           |
| `elem:getAttr(key)`              | Obtiene un atributo HTML.            |
| `elem:hasAttr(key)`              | Verifica si existe el atributo.      |
| `elem:removeAttr(key)`           | Elimina un atributo.                 |
| `elem:setProps({ key = value })` | Establece varios atributos a la vez. |

#### Estilo y contenido

| Método                           | Descripción                                   |
| -------------------------------- | --------------------------------------------- |
| `elem:setStyle({ key = value })` | Modifica CSS en línea.                        |
| `elem:setText(value)`            | Actualiza el texto interno.                   |
| `elem:append(child)`             | Añade un hijo (elemento o texto).             |
| `elem:appendText(value)`         | Añade texto adicional al contenido existente. |
| `elem:clear()`                   | Elimina todos los hijos del elemento.         |
| `elem:remove()`                  | Elimina el elemento del DOM.                  |
| `elem:replaceWith(newElem)`      | Reemplaza el elemento por otro.               |

#### Visibilidad

| Método          | Descripción                                         |
| --------------- | --------------------------------------------------- |
| `elem:hide()`   | Oculta el elemento (`display: none`).               |
| `elem:show()`   | Muestra el elemento (`display: block`).             |
| `elem:toggle()` | Alterna visibilidad y mantiene el display original. |

#### Clases CSS

| Método                        | Descripción                                    |
| ----------------------------- | ---------------------------------------------- |
| `elem:addClass(className)`    | Añade una clase CSS.                           |
| `elem:removeClass(className)` | Remueve una clase CSS.                         |
| `elem:hasClass(className)`    | Devuelve `true` si el elemento tiene la clase. |
| `elem:toggleClass(className)` | Alterna la clase CSS.                          |

---

### Root

```lua
local root = DOM.root()  -- Devuelve el contenedor principal (<body>)
```

---

### Eventos

> [!WARNING]
> Los eventos estan hechos v.

- Se registran con `elem:on(event, callback)`.
- El callback recibe un objeto `evt` con la siguiente información **completa**:

```lua
evt.target     -- userdata del elemento que disparó el evento
evt.value      -- valor actual (para <input>, <textarea>, etc.)
evt.checked    -- booleano para <input type="checkbox|radio">
evt.preventDefault()   -- función para evitar acción por defecto del navegador
evt.stopPropagation()  -- función para detener propagación del evento
```

- Ejemplo de uso:

```lua
button:on("click", function(evt)
    print("Botón clickeado:", evt.target:getAttr("id"))
end)

input:on("input", function(evt)
    print("Valor actual:", evt.value)
end)
```

---

### Crear elementos arbitrarios

```lua
local myDiv = DOM.create("div", { id = "custom", class = "blue" }, "Contenido")
```

- `DOM.create(tag, props, children)` permite crear cualquier tag HTML.
- `children` puede ser `string`, un `userdata` DOM, o una tabla con múltiples elementos.

---

### Devnote

- Actualmente se documentan y soportan los tags principales: `div`, `span`, `button`, `input`.
- Faltan otros elementos HTML como `p`, `ul`, `li`, `textarea`, `select`, `option`, etc.
- `evt` está incompleto
