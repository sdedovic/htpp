# htpp

A thin wrapper around Go's `html/template` that adds an `extends` directive for template inheritance. Built to power a ParcelJS plugin so Go templates can be used with the Parcel development server.

## How it works

A child template declares its parent with `extends` on the first line:

```
extends ./base.htpp

{{ define "body" }}
<div>
  <p>Hello from home!</p>
</div>
{{ end }}
```

`htpp` resolves the full extends chain, compiles the templates in order (base first, then each child), and renders the result.

See [`examples/simple/`](./examples/simple) for a working example.

## Installation

**Nix:**
```bash
nix run github:sdedovic/htpp -- <filename>
```

**Go:**
```bash
go install github.com/sdedovic/htpp/cmd/htpp@latest
```

## CLI

```
htpp [options] <filename>

Options:
  --stdin               Read JSON from stdin and apply it as template data
  --print-dependencies  Print the extends chain (one file per line). Does not render.
```

**Render a template:**
```bash
htpp examples/simple/home.htpp
```

**Render with data from stdin:**
```bash
echo '{"name": "world"}' | htpp --stdin examples/simple/home.htpp
```

**Print dependency chain (for cache invalidation):**
```bash
htpp --print-dependencies examples/simple/home.htpp
```

## Syntax

Only relative paths are supported in `extends`:

```
extends ./base.htpp      # sibling file
extends ../layouts/base.htpp  # parent directory
```

The rest of the file is a standard Go `html/template`. Use `{{ define "block" }}` to override blocks defined in the parent.
