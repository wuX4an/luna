# Notify

**Description:**
API to show native browser notifications from Lua. Wraps the [Web Notifications API](https://developer.mozilla.org/en-US/docs/Web/API/notification) and is integrated with `await`.

---

### Functions

```lua
notify.requestPermission()      -- Requests user permission (await-compatible, returns "granted"/"denied")
notify.permission()             -- Returns the current permission status
notify.show(title, options)     -- Displays a notification
```

---

### `notify.show` Options

`options` is a table with optional fields:

- `body` _(string)_ → Notification body text.
- `icon` _(string)_ → URL of an icon to display.
- `image` _(string)_ → Large image.
- `tag` _(string)_ → Unique identifier (replaces previous notifications with the same tag).
- `song` _(string)_ → Notification sound.
- `silent` _(boolean)_ → Mutes the notification.

---

### Example Usage

```lua
local notify = require("std:web").notify
local await = require("std:web").await

local perm = notify.permission()
print("Permission:", perm)

if perm == "granted" then
    notify.show("Hello!", {
        body = "This is a notification from Lua",
        icon = "icon.png",
    })
else
    await(notify.requestPermission)
end
```

---

### Notes

- `notify.requestPermission()` **must be called inside a user action** (click, keypress) in some browsers.
- `notify.show()` **does not suspend**: it returns immediately (fire-and-forget).
- If permission has not been requested before, some browsers automatically show a prompt when calling `show`.

---
