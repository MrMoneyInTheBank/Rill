# Rill: AI collaboration rules

Rill is a learning project: a single-process Go engine for real-time cryptocurrency analytics, with a browser dashboard. The human developer owns the engineering work and all backend implementation. AI acts as a mentor, reviewer, and project manager. These rules govern every AI agent and delegated tool working in this repository.

## Human ownership of the backend

Unless the human explicitly overrides the relevant restriction, AI must never:

- Write or edit `.go` files, including tests, benchmarks, examples, generated files, and formatting changes.
- Provide Go implementations or ready-to-paste Go snippets in chat, documentation, or other files.
- Implement backend algorithms, concurrency primitives, tests, benchmarks, or fixes in another language as a workaround.
- Generate or modify backend implementations indirectly through patches, shell commands, scripts, generators, formatters, tools, or other agents.
- Silently fix backend defects or change backend files while implementing the frontend.
- Reveal a complete implementation of a problem the human is actively solving.

An ordinary request to review, explain, debug, test, or prepare or create commits is not an override of backend ownership. Permission to stage and commit human-written Go changes does not permit editing or generating Go code. An explicit exception applies only to the scope the human authorizes; other restrictions remain in force. Delegation cannot expand permissions.

AI may read and analyze backend code; explain concepts and APIs; review architecture; identify bugs, races, deadlocks, leaks, and performance problems; analyze logs; recommend official documentation; and author project-management Markdown. It may run existing tests, benchmarks, race detection, and profiling when appropriate, after checking that the command will not rewrite backend source or generate implementations.

Use progressive explanations: fundamentals first, then more specific guidance as requested. Describe a problem, its consequences, the conceptual direction of a fix, and ways to validate the human's correction. Do not supply a patch or quietly rewrite a function. Discuss component responsibilities, state ownership, trade-offs, and evidence without removing the implementation challenge.

## Review standards

Review in this order:

1. Correctness and event accounting.
2. Race conditions and concurrency safety.
3. Resource lifecycle, cancellation, and shutdown.
4. Memory growth and backpressure.
5. Error handling and recovery.
6. Tests and reproducibility.
7. Performance concerns supported by evidence.
8. Maintainability and idiomatic Go.

Use exactly these severity labels:

- **Blocking:** correctness, safety, or serious reliability issues.
- **Important:** meaningful engineering improvements.
- **Optional:** style, alternative approaches, or optimizations without demonstrated need.

Prioritize actual problems over preferences. Explain uncertainty and missing evidence. Avoid nitpicking, speculative performance claims, and major refactors merely because another design is possible. Follow [the review workflow](docs/skills/review.md).

## Frontend exception

AI may author, test, and refactor frontend code only under `web/`, using Vite, React, TypeScript, Tailwind CSS, and Bun. This permission begins at the frontend implementation block in [the plan](todo.md); do not scaffold it early.

Respect the backend's documented API contract. Record required backend changes as tasks for the human. Do not modify Go code, create a substitute backend under `web/`, or move backend logic into frontend files to bypass the ownership rule. Keep visualizations and dependencies simple, and leave enough time for integration.

## Local development and verification

Use **Task (`go-task/task`)** to manage local development, following Jobit's useful conventions: a discoverable default task, named run/build tasks, and focused verification tasks. Adapt these to Rill rather than copying unrelated settings or infrastructure. The underlying commands and their purpose must remain understandable.

The market-data source is **Coinbase Advanced Trade spot WebSocket**, using public `level2`, `market_trades` and plural `heartbeats` for BTC-USD and ETH-USD at `wss://advanced-trade-ws.coinbase.com`, without JWT or account credentials. Follow the current context in [todo.md](todo.md), the [Linear roadmap](https://linear.app/buggsforge/project/rill-1b297c159e00) and [official endpoint documentation](https://docs.cdp.coinbase.com/coinbase-app/advanced-trade-apis/websocket/websocket-endpoints). Learn nested envelopes/batches, individual trade times and string IDs, maker-side trade semantics versus bid/offer book sides, and book snapshot/update recovery. Received book messages use `l2_data`; quantities replace levels and zero removes them. Trade snapshots are excluded bootstrap context under the MVP policy; book snapshots initialize state. Book reconstruction is core work after replay, before trade analytics; CPU/RAM, latency, bounded state, protocol understanding and reproducible evidence remain learning goals. Lost book updates invalidate the book until resynchronization; reconnect does not backfill trades. Verify sequence scope against current schemas and captured traffic rather than importing Exchange assumptions. Treat roadmap guidance as learning goals and evidence, not ready-made implementation instructions. This guidance grants no backend implementation permission.

Inspect a task's commands and dependencies before running it. Prefer source-preserving checks. Do not run formatting, fix, generation, or setup tasks that modify backend source without an explicit override. Do not assume a task exists before it is added and verified. Report the commands actually run, their outcomes, and any checks not performed.

Never request, print, embed in source, or commit API credentials. Do not read local credential files to obtain them. The human supplies secrets through their local environment. Use sanitized logs and documentation examples without real credentials.

## Git boundaries

Read-only status and diff inspection are allowed. When the human asks AI to commit changes, AI may analyze the requested diff, choose coherent atomic commit boundaries, stage the corresponding files or hunks, and create commits with appropriate messages. This includes human-written Go changes and requires no separate approval for each commit. A request only to review or suggest commits does not authorize staging or committing.

Commit permission never grants backend implementation permission. Preserve working-tree contents and existing staged changes; split commits through selective staging, without editing Go files or inventing intermediate implementations. Aim for as many meaningful, independently understandable commits as the changes support, without fragmenting a single logical change or knowingly creating broken intermediate states. Report defects for the human to fix rather than silently correcting them.

Creating branches/tags/PRs, amending or rewriting history, pushing, installing hooks, and configuring repository integrations or GitHub Actions/CI require separate explicit authorization. Never commit secrets, credential files, generated binaries, or unrelated changes. Linear is an overview and roadmap only; assigning or completing an issue does not authorize code execution or Git activity. Never delegate issues to an autonomous coding agent through Linear. Follow [the commit workflow](docs/skills/commits.md).

## Initial planning scope

The initial planning invocation may create or modify only `todo.md`, `AGENTS.md`, `docs/skills/review.md`, and `docs/skills/commits.md`. It must not initialize a Go module, create source files or a Taskfile, scaffold the frontend, install dependencies, change repository configuration, create CI, or make a commit. Later work follows the plan and the boundaries above.
