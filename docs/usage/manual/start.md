## Quick start

Run a script:

```bash
luna run hello.lua
```

Evaluate inline code:

```bash
luna eval 'print("Hello, Luna!")'
```

Build into an executable:

```bash
luna build hello.lua -o hello
./hello
```

---

## CLI Overview

- `luna build <file>` — compile the script into a self-contained executable.
- `luna clean` — remove the build directory.
- `luna docs` — open Luna's documentation.
- `luna eval <code>` — evaluate a Lua snippet inline.
- `luna init` — initialize a new project.
- `luna repl` — start an interactive REPL (Read-Eval-Print Loop).
- `luna run <file>` — run a Lua program or project.
- `luna test` — run tests.
- `luna help [command]` — show help about Luna or a specific command.
