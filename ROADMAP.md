---
# Roadmap para Runtime Lua Moderno ("Luna" provisional)
---

## Fase 1: Fundamentos y Prototipo Básico (0-3 meses)

- **Diseño detallado de la arquitectura interna**
  - Núcleo runtime (interpreter o binding Lua + Go/Rust/C)
  - Módulo de carga y sandboxing básico

- **Implementación inicial de STD básica**
  - `fs`, `os`, `path`, `cli` (parser simple), `log`
  - `random`, `math`, `stringx`, `tablex`
  - Tests unitarios y docs iniciales

- **Herramienta CLI básica**
  - `luna run <file.lua>`
  - Argumentos y flags simples

- **Publicación del repositorio con README + roadmap público**

---

## Fase 2: STD avanzada y soporte multiplataforma (3-6 meses)

- **Completar STD clave**
  - `http` cliente y servidor básico
  - `async` con corutinas y promesas
  - `json`, `yaml`, `toml` (parsing/serialización)
  - `sqlite` embebido con consultas parametrizadas
  - `embedded` recursos embebidos en binario

- **Mejoras CLI**
  - Subcomandos, auto-generación de help, validación de flags
  - Interactividad (input, confirm)

- **Cross-compiling básico**
  - Empaquetar binarios para Linux/macOS/Windows

- **Tests, benchmarks y optimización de rendimiento**

---

## Fase 3: Ecosistema, documentación y comunidad (6-9 meses)

- **Documentación completa y amigable**
  - Tutoriales, ejemplos reales, API reference online
  - Guías para contribución y empaquetado de módulos

- **Sistema de paquetes y módulos**
  - Registro oficial o cache descentralizada
  - Importación por URL, versionado, bloqueo de dependencias

- **Extensión de STD**
  - `socket` TCP/UDP con soporte TLS
  - `tar`, `zip` para empaquetado y distribución
  - `validator`, `events`, `template`

- **Integraciones**
  - Plugins para Neovim, LOVE2D
  - Soporte WASM y bindings a C/C++

---

## Fase 4: Madurez y adopción (9-12+ meses)

- **Mejoras en UX y DX**
  - Hot reload, watch mode, debugging integrado
  - CLI extensible y herramientas complementarias

- **Colaboración con empresas y proyectos open source**
  - Promoción, charlas, workshops

- **Soporte extendido y estabilidad**
  - Releases semánticos, testing en CI/CD multiplataforma
  - Seguridad y sandboxing avanzado

- **Evolución de API y runtime**
  - Basado en feedback y tendencias

---

## Tips adicionales

- Mantén comunicación transparente con la comunidad desde el día uno.
- Prioriza calidad de documentación y experiencia de desarrollador.
- Abre early releases para obtener feedback rápido y pivotar.
- Busca colaboraciones y contribuciones externas para crecer.

---

# STD

```
std/
├── async.lua           -- Corutinas modernas y tareas asincrónicas
│   ├── async.sleep(ms)
│   ├── async.spawn(fn)
│   ├── async.await(promise)
│   ├── async.all({ ... })

├── cli.lua             -- CLI argument parsing e interacción
│   ├── cli.args() → { "arg1", "arg2" }
│   ├── cli.flags() → { help=true }
│   ├── cli.input(prompt)
│   ├── cli.confirm(prompt)

├── console.lua         -- Colores y estilos para salida en terminal
│   ├── console.log(color, text)
│   ├── console.colorize("red", "error")
│   ├── console.style("bold", "header")

├── crypto.lua          -- Hash, UUID, aleatoriedad
│   ├── crypto.hash("sha256", data)
│   ├── crypto.hmac("sha1", key, data)
│   ├── crypto.random_bytes(n)
│   ├── crypto.uuid()

├── env.lua             -- Variables de entorno y archivos .env
│   ├── env.load(".env")
│   ├── env.get("API_KEY")
│   ├── env.set("DEBUG", "true")

├── fs.lua              -- Sistema de archivos (seguro y multiplataforma)
│   ├── fs.read(path)
│   ├── fs.write(path, data)
│   ├── fs.exists(path)
│   ├── fs.listdir(path)
│   ├── fs.stat(path)

├── http.lua            -- Cliente y servidor HTTP moderno
│   ├── http.get(url [, opts])
│   ├── http.post(url, body)
│   ├── http.request(method, url, opts)
│   ├── http.listen({ port=8080, handler=req_fn })

├── json.lua            -- JSON encode/decode
│   ├── json.encode(tbl)
│   ├── json.decode(str)

├── log.lua             -- Logging con niveles
│   ├── log.info("msg")
│   ├── log.debug("msg")
│   ├── log.warn("msg")
│   ├── log.error("msg")
│   ├── log.level("debug")

├── net.lua             -- Utilidades de red
│   ├── net.resolve(hostname)
│   ├── net.local_ip()
│   ├── net.ping(host)

├── os.lua              -- Funciones del sistema operativo
│   ├── os.exec("cmd")
│   ├── os.exit(code)
│   ├── os.getenv("VAR")
│   ├── os.platform()

├── path.lua            -- Manipulación de rutas de forma portable
│   ├── path.join(...)
│   ├── path.basename(path)
│   ├── path.normalize(path)

├── process.lua         -- Control de procesos y señales
│   ├── process.spawn(cmd [, opts])
│   ├── process.kill(pid)
│   ├── process.on("exit", fn)
│
│
├── sqlite.lua           -- Base de datos SQLite embebida con soporte avanzado
│   ├── sqlite.open(filename [, options]) → db       -- Abre o crea base de datos, opciones opcionales
│   ├── sqlite.in_memory() → db                          -- Base de datos SQLite en memoria volátil
│   ├── db:exec(sql) → (ok, err)                      -- Ejecuta sentencia SQL sin resultado
│   ├── db:query(sql, params) → iterator              -- Consulta con parámetros, retorna iterator de filas
│   ├── db:prepare(sql) → stmt                          -- Prepara sentencia para ejecución repetida
│   ├── stmt:bind(params)                               -- Asocia parámetros a la sentencia preparada
│   ├── stmt:step() → (row_or_nil, done)                -- Ejecuta paso y retorna fila o nil si terminó
│   ├── stmt:reset()                                    -- Reinicia la sentencia preparada para re-ejecución
│   ├── db:close()                                      -- Cierra la base de datos y libera recursos
│   ├── db:transaction(fn) → (ok, err)                  -- Ejecuta función dentro de una transacción atómica
│
├── stringx.lua         -- Extensiones útiles de string
│   ├── stringx.starts_with(str, prefix)
│   ├── stringx.trim(str)
│   ├── stringx.split(str, sep)

├── tablex.lua          -- Extensiones de tabla (como lodash)
│   ├── tablex.map(tbl, fn)
│   ├── tablex.filter(tbl, fn)
│   ├── tablex.clone(tbl)

├── template.lua        -- Mini sistema de templates
│   ├── template.render("Hello, {{name}}", { name="Lua" })

├── test.lua            -- Testing unitario sencillo
│   ├── test("description", function() ... end)
│   ├── expect(value):to_equal(expected)

├── time.lua [X]           -- Tiempo y fecha
│   ├── time.now()
│   ├── time.sleep(ms)
│   ├── time.format(fmt, timestamp)

├── yaml.lua            -- YAML decode (si se incluye un parser mínimo)
│   ├── yaml.decode(str)
│   ├── yaml.encode(table)

├── toml.lua            -- YAML decode (si se incluye un parser mínimo)
│   ├── toml.decode(str)
│   ├── toml.encode(table)

├── embedded.lua         -- Recursos embebidos dentro del binario (como Vite/Bun/Deno)
│   ├── embedded.read("path/inside/package.txt") → string
│   ├── embedded.list("assets/") → { "a.png", "b.css" }
│   ├── embedded.exists("file.ext")

├── socket.lua           -- TCP/UDP/raw sockets (nivel bajo)
│   ├── socket.tcp()
│   ├── socket.udp()
│   ├── sock:connect(host, port)
│   ├── sock:send(data)
│   ├── sock:recv(bytes)
│   ├── sock:close()
│
├── tar.lua              -- Empaquetado TAR simple
│   ├── tar.create(files: { "a.txt", "b.lua" }, "out.tar")
│   ├── tar.extract("archive.tar", "output_dir/")
│   ├── tar.list("archive.tar") → { "a.txt", "b.lua" }
│
├── zip.lua              -- ZIP comprimido (sin dependencias externas)
│   ├── zip.create(files, "out.zip")
│   ├── zip.extract("archive.zip", "dest/")
│   ├── zip.list("archive.zip")
│
├── io.lua               -- Flujo de entrada/salida
│   ├── io.read_stdin()
│   ├── io.write_stdout(text)
│   ├── io.write_stderr(text)
│   ├── io.is_tty(stream)
│
├── random.lua           -- Generador de números y selección (imitar la nueva API de js)
│   ├── random.int(min, max) -> int
│   ├── random.float() -> float
│   ├── random.choice({ "a", "b", "c" }) -> string
│   ├── random.shuffle(tbl) -> ???
│
├── math.lua             -- Funciones matemáticas extendidas
│   ├── math.clamp(x, min, max)
│   ├── math.lerp(a, b, t)
│   ├── math.round(n [, decimals])
│   ├── math.sign(n)
│
├── base64.lua           -- Codificación y decodificación
│   ├── base64.encode(str) → "string"
│   ├── base64.decode(b64str) → string
```

---

# CLI

```
$ luna --help
Luna: A modern lua runtime

Usage: luna [OPTIONS] [COMMAND]

Environment variables:
  DEBUG_RUNTIME_ENABLE  | Show runtime logs in stdout
  DEBUG_BUILD_ENABLE    | Show build runtime logs in stdout

Options:
  -h, --help      | Print help
  -v, --version   | Print version

Commands:
  run       | Run a lua program, or project
            |  luna run main.lua | luna run .
  build     | Compile the script into a self contained executable
            |  luna build main.ts | luna build .
            |  luna build . --target=linux/darwin/windows --arch=arm64/amd64/
  test      | Run tests
            |  luna test main.lua | luna test .
  clean     | Remove the *build* directory
            |  luna clean .
  doc       | Generate and show documentation for a script or a module
            |  luna doc main.lua | luna doc . --dir=doc/
  init      | Initialize a new project
            |  luna init . | luna init hello-world
  repl      | Start an interactive Read-Eval-Print Loop (REPL)
```

---

# WORKFLOW

```
.
├── build/          -- luna build
│   ├── debug
│   │   ├── darwin
│   │   ├── linux
│   │   └── windows
│   └── release
│       ├── darwin
│       ├── linux
│       └── windows
├── doc/            -- luna doc
├── .git/           -- luna init
├── .gitignore      -- luna init
├── Luna.toml       -- luna init
├── build.lua       -- luna init
└── src/            -- luna init
    └── main.lua
```

---

# Luna.toml

```
[module]
name = "app"
version = "0.1.0"

[build]
optimized = true # require golang installed
```

---

Chatgpt
¡Muy buen avance 🚀!
Tu roadmap, STD y CLI para **Luna** ya están bastante completos, pero si lo que quieres es **diferenciarte de Rust/JS/Python/etc.**, yo me enfocaría en **DX (developer experience)** y en **features modernas que otros ecosistemas aún no hacen bien**.

Aquí te dejo ideas que podrías agregar para destacar:

---

## 🔥 Features diferenciales (más allá de Rust/JS/Python)

### 1. **Compilación Universal y Ultra-Ligera**

- `luna build main.lua --static --tiny` → produce un **binario estático de <1MB**, listo para distribuir sin dependencias (tipo Go, pero más liviano).
- Cross-compiling transparente: desde Linux puedes compilar para macOS/Windows/ARM sin toolchains extra (esto ya te diferenciaría muchísimo de Node y Python).

### 2. **Bundling de Proyecto Estilo Bun/Deno**

- `luna build .` empaqueta todo el proyecto (incluyendo assets, `.lua`, binarios embebidos, recursos estáticos).
- Soporte para **embedded assets** sin necesidad de "extra tooling" (tipo `luna.embed "assets/"`).
- Esto haría que Luna sea ideal para apps portables (CLI, TUI, servers pequeños).

### 3. **STD con APIs Modernas Inspiradas en JS/Python pero Simples**

- `http.fetch` estilo `fetch` de JS.
- `random` estilo moderna `crypto.getRandomValues`.
- `async` inspirado en Promises/Await pero minimalista (no necesitas todo el peso de Node).
- API declarativa para CLI y logs (`cli.option("--port", { default=8080 })`).

### 4. **Reactividad en el REPL**

- REPL interactivo que recuerde contexto, permita importar código en caliente y hasta _watch mode_ (`luna run --watch main.lua`).
- Algo que Rust no tiene, y JS solo lo logra con herramientas externas.

### 5. **Sandboxing y Seguridad por Diseño**

- Poder correr scripts con permisos controlados tipo Deno:

  ```sh
  luna run app.lua --allow-net --deny-fs
  ```

- Diferenciarte de Node (que es inseguro por default) y de Python (que no tiene sandbox real).

---

## 📦 Ecosistema

### 6. **Package Manager minimalista y descentralizado**

- `luna add "github:user/repo@0.1.0"`
- `luna add "https://example.com/module.lua"`
- Como Deno pero **sin lock-in** a un registry central.
- Resolver dependencias vía Git, URL, o caché local.

### 7. **Build Scripts en Lua**

- Similar a Cargo (Rust), pero usando **Lua itself**:

  ```lua
  -- build.lua
  task("default", function()
    print("Building...")
    os.exec("luna build .")
  end)
  ```

---

## 🛠️ Developer Experience

### 8. **Documentación Integrada**

- `luna doc main.lua` → genera docs navegables en HTML o Markdown a partir de anotaciones tipo LDoc/JSDoc.
- Diferencia inmediata vs Node/Python que dependen de tooling externo.

### 9. **Testing como ciudadano de primera clase**

- `luna test` ya lo tienes, pero podrías darle **reportes bonitos en TUI** (como Vitest/Jest).
- `luna coverage` para métricas rápidas sin instalar nada.

### 10. **Integración con Editores**

- LSP nativo desde el runtime (sin instalar servidores externos).
- `luna lsp` → arranca un servidor de lenguaje listo para VSCode/Neovim/Emacs.

---

## 💡 Filosofía

- **Ligero como Lua, DX moderno como JS, portable como Go, compilado como Rust.**
- Si lo pones así, tienes un "punto medio" muy atractivo:
  - más **portable** que Node/Python,
  - más **simple** que Rust/Go,
  - más **seguro** que Lua clásico.

---

👉 Te lanzo una idea diferencial grande:
¿Qué tal si **Luna puede exportar binarios WASM + nativos en un mismo comando**?
Ejemplo:

```sh
luna build . --target=wasm,linux,windows
```

Y te da un `.wasm` junto con ejecutables listos.
Eso pondría a Luna como **runtime universal** (server, CLI y browser).

---
