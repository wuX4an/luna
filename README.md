# 🌙 Luna – Modern Lua Runtime

> **Lightweight ⚡ Fast 🚀 Productive 💡**
> Run Lua scripts, build self-contained binaries, and ship projects with zero hassle.

---

## 🛠️ Quick Start

```bash
# Run a Lua script
luna run main.lua

# Initialize a project
luna init

# Build a self-contained binary
luna build

# Start interactive REPL
luna repl
```

---

## ✨ Why Choose Luna?

Luna is **not just Lua**, it’s a **modern developer experience**:

- 🌍 **Cross-platform** – Linux, macOS, Windows binaries
- 📦 **Embedded assets** – package scripts & resources in one binary
- 🚀 **Fast CLI** – intuitive commands and flags
- 🔒 **Secure sandboxing** – control filesystem & network access
- 📚 **Rich STD Library** – async tasks, HTTP, SQLite, JSON, TOML, YAML, and more

---

## 🔹 CLI Overview

```
$ luna --help
Luna: A modern lua runtime

 Usage: luna [command] [flags]

 Commands:
   build       Compile the script into a self-contained executable
   clean       Remove the build directory
   completion  Generate the autocompletion script for the specified shell
   docs        View Luna's documentation
   eval        Evaluate a Lua snippet
   init        Initialize a new project
   repl        Start an interactive Lua Read-Eval-Print Loop (REPL)
   run         Run a Lua script or project
   task        Run or list tasks
   test        Run tests

 Flags:
  -h, --help      help for luna
  -v, --version   Print version
```

---

## 📦 Standard Library (STD)

Luna comes with a **powerful and modern STD**:

| Module   | Purpose                                      |
| -------- | -------------------------------------------- |
| `base64` | Encode/decode Base64 strings                 |
| `crypto` | Hash, HMAC, UUID, randomness                 |
| `env`    | Load & manage environment variables (`.env`) |
| `http`   | Client & server                              |
| `ipc`    | Inter-process communication utilities        |
| `math`   | Extended math functions                      |
| `random` | Random int/float, shuffle, choice            |
| `sqlite` | Embedded SQLite DB with transactions         |
| `tablex` | Extended table utilities                     |
| `test`   | Unit testing helpers                         |
| `time`   | Time & date helpers                          |

## 🌙 Project Layout

```
.
├── src/          # Your Lua scripts
├── dist/         # Generated binaries
├── Luna.toml     # Project config
```

---

## ⚡ Key Features

- 🏗 **Cross-platform build** – Linux, macOS, Windows, ARM64
- 🧩 **Embedded resources** – package images, configs, and scripts together
- 🔄 **Hot REPL** – edit scripts live, preserve context, watch mode
- 📖 **Integrated documentation** – `luna docs`
- 🧪 **Unit testing** – simple, readable test suite (`luna test`)

---

## 🤝 Contributing

I ❤️ contributors!

```bash
git clone https://github.com/wuX4an/luna.git
cd luna
luna run build.lua
```

Check [CONTRIBUTING](CONTRIBUTING.md) for details.

---

## 💡 Motto

> **“Lightweight as Lua, productive as JS, portable as Go, modern as Rust.”**
