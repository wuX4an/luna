# Timer

**Description:**
Provides timers and repeated task execution from Lua in the browser.

---

### Functions

```lua
setTimeout(fn, ms)      -- Executes fn once after ms milliseconds
setInterval(fn, ms)     -- Executes fn repeatedly every ms milliseconds
clearTimeout(id)        -- Cancels a timeout
clearInterval(id)       -- Cancels an interval
```

---

### Example

```lua
local id = setTimeout(function()
    print("Hello after 1 second")
end, 1000)

-- Example of interval
local count = 0
local intervalId = setInterval(function()
    count = count + 1
    print("Interval tick:", count)
    if count >= 5 then
        clearInterval(intervalId)
    end
end, 1000)
```

- `setTimeout` returns an identifier that can be passed to `clearTimeout` to cancel it.
- `setInterval` returns an identifier that can be passed to `clearInterval` to stop the repeated execution.

---
