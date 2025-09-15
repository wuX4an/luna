## `std/web/audio/README.md`

# Audio

**Descripción:**
Permite reproducir y controlar audio en el navegador desde Lua usando WASM. Expone un objeto `Audio` que envuelve el objeto `Audio` de JavaScript.

### Overview

```lua

local Audio = require("std:web").audio

-- Crear un nuevo audio
local music = Audio.new("bg_music.mp3")

-- Reproducir y controlar el audio
music:play()              -- Reproduce
music:setVolume(0.5)      -- Volumen al 50%
music:setLoop(true)        -- Activar loop
music:setRate(1.2)        -- Reproducción 1.2x
music:setPosition(10)      -- Saltar a 10 segundos

```

### Funciones principales

- `Audio.new(src)` → Crea un nuevo objeto `Audio` con la ruta `src` (string).

  ```lua
  local music = Audio.new("bg_music.mp3")
  ```

### Métodos disponibles en Lua

| Método             | Descripción                                                                                                                                  |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `play()`           | Reproduce el audio.                                                                                                                          |
| `pause()`          | Pausa la reproducción.                                                                                                                       |
| `stop()`           | Detiene y reinicia el audio (`currentTime = 0`).                                                                                             |
| `setVolume(vol)`   | Ajusta el volumen entre 0.0 y 1.0. **Problema:** requiere número, no table.                                                                  |
| `getVolume()`      | Retorna el volumen actual (número).                                                                                                          |
| `setLoop(bool)`    | Activa/desactiva loop.                                                                                                                       |
| `setRate(rate)`    | Ajusta la velocidad de reproducción (`playbackRate`).                                                                                        |
| `setPosition(pos)` | Salta a `pos` segundos en el audio.                                                                                                          |
| `getPosition()`    | Devuelve la posición actual (segundos).                                                                                                      |
| `getMaxPosition()` | Devuelve la duración del audio (`duration`). ⚠ **Problema:** puede devolver `NaN` si los metadatos aún no están cargados.                   |
| `getMinPosition()` | Devuelve 0 (mínimo).                                                                                                                         |
| `getDuration()`    | Igual que `getMaxPosition()`. ⚠ **Problema:** solo funciona después de `loadedmetadata`. Puede devolver `NaN` si se llama demasiado pronto. |
