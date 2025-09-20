# Audio

**Description:**
Enables audio playback and control in the browser from Lua using WASM. Exposes an `Audio` object that wraps JavaScript’s native `Audio`.

---

### Overview

```lua
local Audio = require("std:web").audio

-- Create a new audio
local music = Audio.new("music.mp3")

-- Play and control audio
music:play()              -- Play
music:setVolume(0.5)      -- Set volume to 50%
music:setLoop(true)       -- Enable looping
music:setRate(1.2)        -- Playback at 1.2x speed
music:setPosition(10)     -- Jump to 10 seconds
```

---

### Main Functions

- `Audio.new(src)` → Creates a new `Audio` object from the given `src` (string).

  ```lua
  local music = Audio.new("music.mp3")
  ```

---

### Methods Available in Lua

| Method             | Description                                                                                             |
| ------------------ | ------------------------------------------------------------------------------------------------------- |
| `play()`           | Plays the audio.                                                                                        |
| `pause()`          | Pauses playback.                                                                                        |
| `stop()`           | Stops and resets the audio (`currentTime = 0`).                                                         |
| `setVolume(vol)`   | Sets volume between `0.0` and `1.0`. ⚠ Requires a number, not a table.                                 |
| `getVolume()`      | Returns the current volume (number).                                                                    |
| `setLoop(bool)`    | Enables/disables looping.                                                                               |
| `setRate(rate)`    | Sets playback speed (`playbackRate`).                                                                   |
| `setPosition(pos)` | Jumps to `pos` seconds in the audio.                                                                    |
| `getPosition()`    | Returns the current position (in seconds).                                                              |
| `getMaxPosition()` | Returns the audio duration (`duration`). ⚠ May return `NaN` if metadata is not yet loaded.             |
| `getMinPosition()` | Returns `0` (minimum).                                                                                  |
| `getDuration()`    | Same as `getMaxPosition()`. ⚠ Only works after `loadedmetadata`. May return `NaN` if called too early. |

---
