# std:sqlite

The `std:sqlite` module provides an easy-to-use interface to SQLite databases from Lua.  
It supports creating databases, executing SQL statements, performing transactions, and iterating over query results.

---

## Functions

### `sqlite.open(path: string): (db, err)`

Opens a SQLite database at the given path. Use `":memory:"` for an in-memory database.

**Example:**

```lua
local sqlite = require("std:sqlite")
local db, err = sqlite.open(":memory:")
assert(db, err)
```

---

### `db:exec(sql: string): (ok, err)`

Executes an SQL statement without returning rows (e.g., `CREATE`, `INSERT`, `UPDATE`).

**Example:**

```lua
db:exec([[
  CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
  );
]])
```

---

### `db:transaction(fn: function): (ok, err)`

Executes multiple SQL statements atomically inside a transaction.

- `fn` receives a transaction object `tx` to run multiple `tx:exec()` calls.

**Example:**

```lua
db:transaction(function(tx)
    tx:exec("INSERT INTO users (name) VALUES ('Ana')")
    tx:exec("INSERT INTO users (name) VALUES ('Luis')")
end)
```

---

### `db:query(sql: string): (rows, err)`

Executes a query returning rows (e.g., `SELECT`).

- `rows` is an iterator object with methods:
  - `rows:next() -> table?` — fetch the next row or `nil`.
  - `rows:close()` — close the iterator.
  - `rows:iter()` — optional iterator function compatible with `for r in rows:iter() do ... end`.

**Example using `next()`:**

```lua
local rows, err = db:query("SELECT id, name FROM users ORDER BY id")
while true do
    local row = rows:next()
    if not row then break end
    print(string.format("ID=%d, Name=%s", row.id, row.name))
end
rows:close()
```

**Example using `iter()`:**

```lua
local rows = db:query("SELECT id, name FROM users ORDER BY id")
for r in rows:iter() do
    print(r.id, r.name)
end
rows:close()
```

---

### `db:close()`

Closes the database connection.

**Example:**

```lua
db:close()
```

---

## Complete Example

```lua
local sqlite = require("std:sqlite")

-- Open an in-memory database
local db, err = sqlite.open(":memory:")
assert(db, err)

-- Create a table
db:exec([[
CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);
]])

-- Insert data inside a transaction
db:transaction(function(tx)
    tx:exec("INSERT INTO users (name) VALUES ('Ana')")
    tx:exec("INSERT INTO users (name) VALUES ('Luis')")
    tx:exec("INSERT INTO users (name) VALUES ('Carlos')")
end)

-- Query and display users
local rows = db:query("SELECT id, name FROM users ORDER BY id")
print("=== Users in the table ===")
while true do
    local row = rows:next()
    if not row then break end
    print(string.format("ID=%d, NAME=%s", row.id, row.name))
end
rows:close()

-- Count total users
local countRows = db:query("SELECT COUNT(*) AS total FROM users")
local countRow = countRows:next()
countRows:close()
print("\n=== Total users ===")
print("TOTAL USERS:", countRow.total)

-- Close the database
db:close()
```

---
