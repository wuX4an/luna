# Contributing to Luna

First of all, **thank you** for considering contributing to Luna! 🙌  
Whether it's a bug fix, a new feature, or improving docs, your help makes Luna better for everyone.

---

## Table of Contents

1. [Getting Started](#getting-started)
2. [Development Workflow](#development-workflow)
3. [Writing Tests](#writing-tests)
4. [Style Guide](#style-guide)
5. [Documentation](#documentation)
6. [Reporting Issues](#reporting-issues)
7. [Community](#community)

---

## Getting Started

1. **Fork** the repository.
2. **Clone** your fork:

   ```bash
   git clone https://github.com/<your-username>/luna.git
   cd luna
   ```

3. **Download the dependencies**:

   ```bash
   go mod download
   ```

---

## Development Workflow

- Create a new branch for each feature or fix:

  ```bash
  git checkout -b feature/my-new-feature
  ```

- Make your changes.

- Commit with clear messages:

  ```bash
  git add .
  git commit -m "Add awesome feature X"
  ```

- Push your branch and open a Pull Request.

---

## Writing Tests

Luna includes a built-in testing library (`std/test.lua`).

Example:

```lua
local test = require("test")

test("basic math works", function()
  local sum = 1 + 1
  expect(sum):to_equal(2)
end)
```

- Add tests for **all new features**.
- Ensure **no tests are failing** before submitting a PR.

---

## Documentation

- Update `docs/` for any public API changes.
- Include examples when possible.
- Use Markdown tables, code blocks, and links for clarity.

---

## Reporting Issues

Found a bug or have a suggestion? Open an issue! Include:

- Steps to reproduce
- Expected vs actual behavior
- OS / Luna version
- Minimal reproducible code if applicable

---

## Community

- Join discussions on [Github Discussions](https://github.com/wuX4an/luna/discussions)
- Be friendly and respectful. Help others whenever you can!

---

Thanks for contributing to **Luna**! 🚀
