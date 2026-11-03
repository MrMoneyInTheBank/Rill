# Rill

Rill is a learning project: a single-process Go engine for real-time cryptocurrency analytics, with a browser dashboard planned.

The current application is a feed exploration tool. It connects to Coinbase Advanced Trade, requests BTC-USD order-book and trade data plus heartbeats, and logs decoded messages. Book reconstruction, trade analytics, replay, recovery, and the dashboard are not implemented yet.

## Prerequisites

- Go 1.27.2 or newer, as required by `go.mod`.
- [Task](https://taskfile.dev/) (`go-task/task`) available as `task`.
- Network access to download Go dependencies and connect to `wss://advanced-trade-ws.coinbase.com` for live exploration.

Run commands from the repository root. Go downloads module dependencies when needed. Public `level2`, `market_trades`, and `heartbeats` subscriptions require no account or credentials; no `.env` file is required. Task loads a local `.env` if present, but the feed no longer reads or sends JWTs. Keep credentials out of source, logs, and Git.

## Run and build

```sh
task
task run
```

`task` lists available commands. `task run` runs vet and tests, builds the application, then starts `build/rill`. Press Ctrl+C to request shutdown. The application releases its connection before reporting a returned error and exiting unsuccessfully.

To build without starting the live feed:

```sh
task build
```

This runs vet and tests before compiling `cmd/feed` into `build/rill`. Building does not connect to Coinbase. The executable targets the operating system and architecture of the Go build environment by default. Build output is ignored by Git.

Run the existing executable directly to skip rebuilding and verification:

```sh
./build/rill
```

The experiment logs sent requests and decoded received messages; it no longer collects payload-size histories. It does not reconnect automatically. A sent-request log records sending a request; inspect received acknowledgments to establish subscription acceptance. A parsing failure or non-text message currently ends the session rather than being skipped.

## Local commands

| Command | Behavior |
| --- | --- |
| `task` / `task --list` | List available tasks. |
| `task lint` | Run `go vet ./...`. |
| `task test` | Run `go test ./...`. |
| `task check` | Run vet and tests without formatting source. |
| `task build` | Run checks and build `build/rill`. |
| `task run` | Build, then run `build/rill`; application arguments are not forwarded. |
| `task fmt` | Rewrite Go source formatting under `cmd/` and `internal/`; invoke explicitly when wanted. |
| `task clean` | Remove the build directory. |

The test task forwards arguments after `--`, for example:

```sh
task test -- -race
task test -- -run TestSendMsg -v
```

Existing tests cover the WebSocket transport helpers using local mock servers. They need permission to bind a loopback port. Passing them does not verify live provider behavior, parser correctness, or the planned analytics. There is currently no formatting-check task; `task fmt` changes files and is separate from `task check`.

## Learning plan

See [todo.md](todo.md) for the current scope and learning sequence, and [AGENTS.md](AGENTS.md) for collaboration boundaries. The current context at the top of the plan supersedes its historical provider and scheduling assumptions. The human owns all Go implementation, tests, and benchmarks; AI supports review, mentoring, and documentation.
