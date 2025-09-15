## `std/web/timer/README.md`

# Timer

**Descripción:** Temporizadores y repetición de tareas.

### Funciones

```lua
setTimeout(fn, ms)       -- Ejecuta fn una vez después de ms ms
setInterval(fn, ms)      -- Ejecuta fn repetidamente cada ms ms
clearTimeout(id)         -- Cancela timeout
clearInterval(id)        -- Cancela intervalo
```

Ejemplo:

```lua
local id = setTimeout(function()
    print("Hola después de 1s")
end, 1000)
```

> [!NOTE]
> Node anda de necio con puros errores pendejos, no lo voy a hacer
