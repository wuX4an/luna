## `std/web/notify/README.md`

# Notify

**Descripción:** API para mostrar notificaciones nativas del navegador desde Lua.
Es un wrapper sobre la [Web Notifications API](https://developer.mozilla.org/en-US/docs/Web/API/notification), integrado con `await`.

---

### Funciones

```lua
notify.requestPermission()      -- Pide permiso al usuario (await-compatible, devuelve "granted"/"denied")
notify.permission()             -- Muestra el estado del permiso
notify.show(title, options)     -- Muestra una notificación
```

---

### Opciones de `notify.show`

`options` es una tabla con campos opcionales:

- `body` _(string)_ → Texto del cuerpo.
- `icon` _(string)_ → URL de un ícono a mostrar.
- `image` _(string)_ → Imagen grande.
- `tag` _(string)_ → Identificador único (reemplaza notificación previa con el mismo tag).
- `song` (string) -> Sónido de notificación
- `silent` _(boolean)_ → Desactiva sonido.

---

### Ejemplo de uso

```lua
local notify = require("std:web").notify
local await = require("std:web").await

local perm = notify.permission()
print("Permiso:", perm)

if perm == "granted" then
	notify.show("Hola!", {
		body = "Esto es una notificación desde Lua",
		icon = "icon.png",
	})
else
	await(notify.requestPermission)
end

```

---

### Notas

- `notify.requestPermission()` **debe ejecutarse dentro de una acción del usuario** (click, keypress) en algunos navegadores.
- `notify.show()` **no suspende**: devuelve inmediatamente (fire-and-forget).
- Si no se pidió permiso previamente, algunos navegadores muestran un prompt automático al llamar `show`.
