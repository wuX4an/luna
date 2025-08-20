# Reference

The Luna reference provides detailed documentation for the command-line interface, the Lua runtime, and the standard libraries.

---

## CLI

Luna ships with a built-in command line tool that helps you run, build, and manage projects.  
Here you will find a complete list of commands, flags, and usage examples.

- [CLI Overview](/reference/cli)
- Commands: `run`, `eval`, `repl`, `build`, `init`, `test`, `docs`, `clean`

---

## Lua

Luna extends the Lua runtime with additional features and integrations.
This section provides a reference for the runtime behavior, global functions, module loading, and a minimal Lua (bare-bones) reference.

- [Lua Overview](/reference/lua)
- Topics: runtime integration, module system, interoperability

---

## Standard Library

Luna includes a set of standard modules under the `std:` namespace.
These modules provide utilities for cryptography, base64, time/date, and more.

- [Std Overview](/reference/std)
- Modules: `std:crypto`, `std:base64`, `std:time`, ...

---

## Config

Luna projects use a configuration file called `Luna.toml` to define project metadata, build instructions, and reusable tasks.
This file controls how Luna runs, builds, and manages your project.

- [Config Overview](/reference/config) — full documentation of `Luna.toml` fields and usage examples
- **File:** `Luna.toml` (TOML format)
- **Sections:** `[package]`, `[build]`, `[tasks]`
- **Purpose:** specify project info, build source, entry point, output, and custom tasks
