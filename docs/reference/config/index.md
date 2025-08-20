# Luna.toml Configuration Reference

This page describes the configuration file used by Luna projects (`Luna.toml`), including sections, fields, and examples.

---

## \[package]

Defines basic metadata for the project.

```toml
[package]
name = "app"
version = "0.1.0"
```

**Fields:**

- `name` — the project name (string, required)
- `version` — the project version (string, optional, default `"0.1.0"`)

---

## \[build]

Specifies how to build the project into a self-contained executable.

```toml
[build]
source = "src/"
entry = "main.lua"
target = "linux/amd64"
output = "dist"
```

**Fields:**

- `source` — directory containing Lua source files
- `entry` — main Lua script to execute (string, required)
- `target` — build target platform (`linux/amd64`, `darwin/amd64`, etc.)
- `output` — directory or file name for the compiled binary

---

## \[tasks]

Defines reusable tasks (scripts) for the project, including optional dependencies.

```toml
[tasks]
source = "tasks/"

[tasks.clean]
script = "clean.lua"
desc = "Clean generated files"

[tasks.build]
script = "build.lua"
desc = "Compile the project"
depends = ["clean"]

[tasks.run]
script = "run.lua"
desc = "Run the project"
depends = ["build"]
```

**Fields:**

- `source` — directory containing task scripts
- Each task (e.g., `[tasks.build]`) can define:
  - `script` — the Lua file to run for the task
  - `desc` — a short description (optional)
  - `depends` — list of other tasks to run first (optional)

**Usage:**

```bash
luna task             # list all tasks
luna task build       # run the "build" task and its dependencies
luna task run         # run the "run" task (will execute "build" first)
```
