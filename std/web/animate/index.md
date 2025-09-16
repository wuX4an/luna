# Animate

**Descripción:** Permite crear animaciones CSS sobre elementos del DOM desde Lua, usando WebAssembly (`syscall/js`). Soporta transformaciones (`scale`, `rotate`, `x`, `y`) y `opacity`, con control de duración, retraso, bucles, modo `keep` y eventos de finalización.

---

### Overview

```lua
local DOM = require("std:web").dom
local Animate = require("std:web").animate

-- Crear elementos
local box = DOM.div({}, "Hover aquí")
DOM.root():append(box)
DOM.root():append(box2)

-- Estilos básicos
box:setStyle({
    width = "120px",
    height = "120px",
    backgroundColor = "#09f",
    color = "#fff",
    display = "inline-block",
    textAlign = "center",
    lineHeight = "120px",
    cursor = "pointer",
    marginRight = "1em",
    marginTop = "1em",
})

-- Animación directa al entrar (auto trigger)
box:on("mouseenter", function()
    Animate.clear(box)
    Animate.animate(box, {
        duration = 800,
        scale = 2,
        rotate = 360,
        opacity = 0.5,
        keep = true,  -- conserva clase y keyframe hasta limpiar manualmente
    })
end)
```

---

### Funciones principales

| Función                       | Descripción                                                                                                                                                                            |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Animate.animate(elem, opts)` | Ejecuta una animación sobre el elemento. `opts` puede incluir: `x`, `y`, `scale`, `rotate`, `opacity`, `duration`, `delay`, `easing`, `loop`, `fill`, `keep`, `trigger`, `onComplete`. |
| `Animate.clear(elem)`         | Limpia todas las animaciones activas del elemento, eliminando clases y keyframes asociados.                                                                                            |

---

### Opciones (`opts`)

| Opción       | Tipo     | Default      | Descripción                                      |
| ------------ | -------- | ------------ | ------------------------------------------------ |
| `x`, `y`     | number   | `0`          | Desplazamiento en px                             |
| `scale`      | number   | `1`          | Escala del elemento                              |
| `rotate`     | number   | `0`          | Rotación en grados                               |
| `opacity`    | number   | `1`          | Opacidad final                                   |
| `duration`   | number   | `1000`       | Duración en ms                                   |
| `delay`      | number   | `0`          | Retraso en ms                                    |
| `easing`     | string   | `"ease"`     | Función de timing CSS                            |
| `loop`       | bool     | `false`      | Repetir infinitamente si true                    |
| `fill`       | string   | `"forwards"` | Cómo se aplica la animación después de terminar  |
| `keep`       | bool     | `false`      | Si true, conserva clase y keyframe tras terminar |
| `trigger`    | string   | `"auto"`     | `"auto"` o `"hover"`                             |
| `onComplete` | function | `nil`        | Callback cuando la animación finaliza            |

---

### Ejemplo avanzado con template (simulación)

Actualmente **Animation.template** **no está disponible**, pero se puede simular con `animate` y `clear`:

```lua
local anim = Animate.template({
    from = { scale = 1, rotate = 0, opacity = 1 },
    to   = { scale = 2, rotate = 360, opacity = 0.5 },
    duration = 800,
    keep = true,
})

box:on("mouseenter", function()
    anim:run(box)
end)
```

- `from` define los valores iniciales.
- `to` define los valores finales.
- `duration` controla la duración.
- `keep` mantiene la clase y keyframe hasta limpiar.

---

### Devnote

> **Plantillas de animación a futuro:**
> La idea es implementar un sistema tipo `Animation.template()` que permita definir animaciones reutilizables con `from {}` y `to {}`. Por ahora, se puede simular usando `animate` manualmente, pero la futura API permitirá:
>
> - Reutilizar animaciones sin repetir las mismas opciones.
> - Llamar `:run(elem)` en cualquier elemento para aplicar la animación.
> - Posibilidad de combinaciones de transformaciones y `opacity`.
>   Esto hará más limpio y mantenible el código de animaciones en Lua.
