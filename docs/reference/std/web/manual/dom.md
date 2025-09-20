# DOM

**Description:**
Manipulates HTML elements and the DOM structure from Lua. Allows creating elements, modifying attributes, styles, content, and handling events using WASM (`syscall/js`) to interact with the browser DOM.

---

### Overview

```lua
local DOM = require("std:web").dom

-- Get root and clear content
local root = DOM.root()
root:clear()

-- Create container
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

-- Create message
local msg = DOM.div({ id = "msg" }, "Hello, world!")
container:append(msg)
msg:setStyle({ padding = "10px", backgroundColor = "#fff", border = "1px solid #ccc" })

-- Create button
local btn = DOM.button({ id = "btn" }, "Change Message")
container:append(btn)
btn:setStyle({ padding = "10px", cursor = "pointer" })

-- Click event: change message
btn:on("click", function(evt)
    msg:setText("Message updated!")
end)
```

---

### Creating Elements

```lua
local div = DOM.div({ id = "container", class = "main" }, "Initial text")
local span = DOM.span({}, "Hello")
local button = DOM.button({ id = "btn1" }, "Click me")
local input = DOM.input({ type = "text", placeholder = "Type something..." })
```

- **Props:** `id`, `class`, `style`, or other HTML attributes.
- **Children:** optional, can be `string` or DOM element (`userdata`).
- **Reference:** `DOM.div`, `DOM.span`, `DOM.button`, `DOM.input`, or `DOM.create(tag, props, children)` for any HTML tag.

---

### Element Methods

#### Attributes & Properties

| Method                         | Description                       |
| ------------------------------ | --------------------------------- |
| `elem:setAttr(key, value)`     | Sets an HTML attribute.           |
| `elem:getAttr(key)`            | Gets an HTML attribute.           |
| `elem:hasAttr(key)`            | Checks if the attribute exists.   |
| `elem:removeAttr(key)`         | Removes an attribute.             |
| `elem:setProps({ key=value })` | Sets multiple attributes at once. |

#### Style & Content

| Method                         | Description                        |
| ------------------------------ | ---------------------------------- |
| `elem:setStyle({ key=value })` | Modifies inline CSS.               |
| `elem:setText(value)`          | Updates inner text.                |
| `elem:append(child)`           | Adds a child (element or text).    |
| `elem:appendText(value)`       | Appends text to existing content.  |
| `elem:clear()`                 | Removes all children.              |
| `elem:remove()`                | Removes the element from the DOM.  |
| `elem:replaceWith(newElem)`    | Replaces the element with another. |

#### Visibility

| Method          | Description                                 |
| --------------- | ------------------------------------------- |
| `elem:hide()`   | Hides the element (`display: none`).        |
| `elem:show()`   | Shows the element (`display: block`).       |
| `elem:toggle()` | Toggles visibility, keeps original display. |

#### CSS Classes

| Method                        | Description                     |
| ----------------------------- | ------------------------------- |
| `elem:addClass(className)`    | Adds a CSS class.               |
| `elem:removeClass(className)` | Removes a CSS class.            |
| `elem:hasClass(className)`    | Returns `true` if class exists. |
| `elem:toggleClass(className)` | Toggles the CSS class.          |

---

### Root

```lua
local root = DOM.root()  -- Returns main container (<body>)
```

---

### Events

> WARNING: Event support is still partial.

- Register with `elem:on(event, callback)`.
- Callback receives `evt` object:

```lua
evt.target       -- DOM element triggering the event
evt.value        -- Current value (for <input>, <textarea>, etc.)
evt.checked      -- Boolean for <input type="checkbox|radio">
evt.preventDefault()  -- Prevent default browser action
evt.stopPropagation() -- Stop event propagation
```

**Example:**

```lua
button:on("click", function(evt)
    print("Button clicked:", evt.target:getAttr("id"))
end)

input:on("input", function(evt)
    print("Current value:", evt.value)
end)
```

---

### Creating Arbitrary Elements

```lua
local myDiv = DOM.create("div", { id="custom", class="blue" }, "Content")
```

- `DOM.create(tag, props, children)` allows creating any HTML tag.
- `children` can be a `string`, a DOM element, or a table with multiple elements.

---

### Devnote

- Main supported tags: `div`, `span`, `button`, `input`.
- Other HTML elements like `p`, `ul`, `li`, `textarea`, `select`, `option` are not yet documented.
- `evt` object is incomplete.

---
