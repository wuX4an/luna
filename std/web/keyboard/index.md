# Keyboard

**Descripción:** Captura y maneja eventos del teclado desde Lua en el navegador usando WebAssembly (`syscall/js`).
Este módulo permite escuchar teclas individuales, combinaciones de teclas, consultar el estado actual de las teclas y opcionalmente desactivar comportamientos por defecto del navegador.

---

### Overview

```lua
local keyboard = require("std:web").keyboard

-- Escuchar cualquier tecla presionada
keyboard.on("keydown", function(evt)
    print("Tecla presionada:", evt.key)
end)

-- Escuchar solo la primera vez que se libera una tecla
keyboard.once("keyup", function(evt)
    print("Tecla liberada:", evt.key)
end)

-- Escuchar combinación Ctrl+S
keyboard.onCombo({"Control", "s"}, function(evt)
    print("Guardar archivo")
    evt:preventDefault()  -- evita acción por defecto del navegador
end)

-- Consultar estado de una tecla
if keyboard.isPressed("Shift") then
    print("Shift está presionada")
end

-- Desactivar keybindings por defecto
keyboard.disableDefaults()  -- desactiva todos
keyboard.disableDefaults({"Tab", "F1", "ArrowUp"})  -- solo algunas teclas
```

---

### Funciones principales

| Función                            | Descripción                                                                            |
| ---------------------------------- | -------------------------------------------------------------------------------------- |
| `keyboard.on(event, callback)`     | Escucha un evento `"keydown"` o `"keyup"`.                                             |
| `keyboard.once(event, callback)`   | Escucha un evento solo una vez.                                                        |
| `keyboard.onCombo(keys, callback)` | Escucha combinaciones de teclas simultáneas.                                           |
| `keyboard.isPressed(key)`          | Devuelve `true` si la tecla indicada está presionada.                                  |
| `keyboard.disableDefaults(keys?)`  | Desactiva el comportamiento por defecto de todas las teclas o de una tabla específica. |

---

### Eventos

- Los eventos se registran con `keyboard.on`, `keyboard.once` o `keyboard.onCombo`.
- Cada callback recibe un objeto `evt` con la siguiente información:

```lua
evt.key             -- la tecla presionada (ej: "a", "Shift", "Control")
evt.code            -- código físico de la tecla (ej: "KeyA", "ShiftLeft")
evt.ctrlKey         -- booleano, true si Ctrl está presionada
evt.shiftKey        -- booleano, true si Shift está presionada
evt.altKey          -- booleano, true si Alt está presionada
evt.metaKey         -- booleano, true si Meta/Command está presionada
evt.preventDefault() -- función para evitar acción por defecto del navegador
evt.stopPropagation()-- función para detener propagación del evento
```

- Ejemplo:

```lua
keyboard.on("keydown", function(evt)
    print("Tecla:", evt.key, "Ctrl presionada:", evt.ctrlKey)
end)
```

---

### Combinaciones de teclas

- `keyboard.onCombo(keys, callback)` permite detectar combinaciones simultáneas.
- `keys` es una tabla de strings con los nombres de las teclas.

```lua
keyboard.onCombo({"Shift", "Alt", "P"}, function(evt)
    print("Combinación activada")
end)
```

- Las combinaciones solo disparan el callback cuando **todas las teclas indicadas están presionadas al mismo tiempo**.

---

### Desactivar keybindings por defecto

- `keyboard.disableDefaults()` → bloquea todas las teclas para que no actúen según el navegador.
- `keyboard.disableDefaults({"Tab", "F1"})` → bloquea solo las teclas indicadas.

```lua
keyboard.disableDefaults()  -- bloqueo global
keyboard.disableDefaults({"ArrowUp", "ArrowDown"})  -- bloqueo selectivo
```

---

### Devnote

- Internamente mantiene un mapa de teclas presionadas para combinaciones y `isPressed`.
- Funciona en navegadores modernos con soporte de WebAssembly y `syscall/js`.
- Actualmente soporta eventos `keydown` y `keyup`; no hay soporte para `keypress` legacy.
- La desactivación de keybindings afecta **solo eventos interceptados por el módulo**, no extensiones o scripts externos.
