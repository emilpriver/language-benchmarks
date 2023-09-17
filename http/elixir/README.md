# Elixir Web Server

### Build

```bash
MIX_ENV=prod mix release
```

### Run

Start
```bash
_build/prod/rel/server/bin/server start
```

Stop (you may also send SIGINT/SIGTERM)
```bash
_build/prod/rel/server/bin/server stop
```

## Dev

#### Run with REPL

```bash
iex -S mix
```
