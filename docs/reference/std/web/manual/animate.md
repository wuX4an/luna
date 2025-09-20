# Animate

**Description:** Provides CSS animation capabilities on DOM elements from Lua using WebAssembly (`syscall/js`). Supports transformations (`scale`, `rotate`, `x`, `y`) and `opacity`, with control over duration, delay, loops, `keep` mode, and completion events.

---

### Overview

```lua
local DOM = require("std:web").dom
local Animate = require("std:web").animate

-- Create elements
local box = DOM.div({}, "Hover here")
DOM.root():append(box)
DOM.root():append(box2)

-- Basic styles
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

-- Direct animation on mouse enter (auto trigger)
box:on("mouseenter", function()
    Animate.clear(box)
    Animate.animate(box, {
        duration = 800,
        scale = 2,
        rotate = 360,
        opacity = 0.5,
        keep = true,  -- preserve class and keyframe until cleared manually
    })
end)
```

---

### Main Functions

| Function                      | Description                                                                                                                                                                   |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Animate.animate(elem, opts)` | Runs an animation on the element. `opts` may include: `x`, `y`, `scale`, `rotate`, `opacity`, `duration`, `delay`, `easing`, `loop`, `fill`, `keep`, `trigger`, `onComplete`. |
| `Animate.clear(elem)`         | Clears all active animations from the element, removing associated classes and keyframes.                                                                                     |

---

### Options (`opts`)

| Option       | Type     | Default      | Description                                     |
| ------------ | -------- | ------------ | ----------------------------------------------- |
| `x`, `y`     | number   | `0`          | Translation in pixels                           |
| `scale`      | number   | `1`          | Element scaling                                 |
| `rotate`     | number   | `0`          | Rotation in degrees                             |
| `opacity`    | number   | `1`          | Final opacity                                   |
| `duration`   | number   | `1000`       | Duration in ms                                  |
| `delay`      | number   | `0`          | Delay in ms                                     |
| `easing`     | string   | `"ease"`     | CSS timing function                             |
| `loop`       | bool     | `false`      | Repeat infinitely if true                       |
| `fill`       | string   | `"forwards"` | How the animation applies after finishing       |
| `keep`       | bool     | `false`      | If true, preserves class and keyframe after end |
| `trigger`    | string   | `"auto"`     | `"auto"` or `"hover"`                           |
| `onComplete` | function | `nil`        | Callback when the animation finishes            |

---

### Advanced Example with Template (simulation)

Currently **Animation.template** is **not available**, but it can be simulated with `animate` and `clear`:

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

- `from` defines initial values.
- `to` defines final values.
- `duration` controls the animation length.
- `keep` preserves class and keyframe until cleared.

---

### Devnote

**Future animation templates:**
The goal is to implement a `Animation.template()` system for reusable animations with `from {}` and `to {}`.

This future API will allow:

- Reusing animations without repeating options.
- Calling `:run(elem)` on any element to apply the animation.
- Combining transformations and `opacity` more cleanly.

---
