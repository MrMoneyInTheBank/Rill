# Rill — prepared Linear roadmap

**Status: prepared in repository, NOT created in Linear.**

Team: `buggsforge` · Project: `rill` · Core checkpoint: Friday October 23, 2026 23:59 IST · Optional deadline: Sunday October 25, 2026 23:59 IST.

**Budget:** 16 core two-hour blocks = **32 focused hours**. Up to six additional two-hour **contingency** blocks (12 hours) on Saturday/Sunday only if mandatory acceptance checks remain unfinished. Maximum 44 focused hours; contingency is not an invitation to add features.

**Inventory:** 16 parent blocks, 64 scheduled learning/evidence tasks, 8 unscheduled stretch tasks. The 64 child task budgets total 32 hours; parents are rollups, not additional time. Local `RILL-*` identifiers are planning keys, **not** Linear issue IDs.

**Provider:** Public **Coinbase Exchange WebSocket**, `wss://ws-feed.exchange.coinbase.com`, channels `matches` and **singular** `heartbeat`. First `last_match` is not a newly arrived live trade. Match `side` is maker side; per-product `sequence` and heartbeat `last_trade_id` help flag possible gaps but are not exact missing-trade counts. Learn/verify this from [official channel docs](https://docs.cdp.coinbase.com/exchange/websocket-feed/channels) and [best practices](https://docs.cdp.coinbase.com/exchange/websocket-feed/best-practices).

**Human ownership:** All Go backend, tests and benchmarks human-authored. Codex may mentor/review and implement frontend only under `web/` during Block 12. No agent-driven CI, branches, PRs, automated repository changes, or Linear-to-Git actions. The explicitly authorized planning-document commits do not expand these permissions.

## Milestones

- **M1 — Foundation: public feed and deterministic replay** · 2026-10-20T21:00:00+05:30
  - Gate: Exchange public `matches`/`heartbeat` samples and `last_match` policy, normalized adapter, cancellable bounded ingress, deterministic JSONL replay.
- **M2 — Engine: correct bounded multi-product analytics** · 2026-10-21T21:00:00+05:30
  - Gate: Fixed/rolling event-time analytics, bounded dedup and accounting, BTC/ETH isolation, retry/backoff and Exchange continuity warnings.
- **M3 — Dashboard: verified engine, metrics and SSE integration surface** · 2026-10-22T21:00:00+05:30
  - Gate: Regression/race gate passes, metrics/API contract, slow-client-safe SSE and frontend fixture/type/build checks.
- **M4 — Release readiness: integration, performance and final acceptance** · 2026-10-23T23:59:00+05:30
  - Gate: Live/replay UI, tested fault handling and backpressure, reproducible match-by-match benchmarks, profiles, soak and documented acceptance.

**Friday checkpoint:** Decide whether the 16 core blocks meet all acceptance criteria. If not, allocate the smallest number of optional weekend blocks to correctness, recovery, replay, backpressure, integration, profiling and docs. **Sunday is the contingency deadline, not the default target.**

## Creation and sizing

Assign to the current human user in the `buggsforge` team, not to bots. Use native T-shirt estimates if configured; otherwise size labels. Do not modify shared Linear workspace estimation settings. All timestamps IST; when only date due fields are available, preserve exact planned time in issue descriptions. Dependencies below reference local keys until created. Linear remains entirely untouched until explicit creation is requested.

- XS: ≤20 min
- S: 21–45 min
- M: 46–90 min
- L: 91–120 min, parent rollup only
- XL: split or stretch only

## Planned issue inventory

### RILL-B01 — Block 01 — Explore Coinbase Exchange wire protocol

**Type:** parent · **Size:** L · **Milestone:** M1 · **Due:** 2026-10-20T12:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, devex, ingestion
**Parent:** None · **Depends on:** None

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Exchange feed overview, subscriptions, `matches`, first `last_match`, per-product `heartbeat`, message sequence, Go modules and Task.

**Goal and evidence:** Capture the actual subscribe acknowledgment, sample matches and heartbeat; verify BTC-USD rate without credentials; explain what is and isn't a fresh execution.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B01.1 — 01.1 — Explore Coinbase Exchange subscriptions

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T10:20:00+05:30
**Budget:** 20 minutes · **Labels:** devex, execution:human, ingestion, scope:mvp, size:XS, study/research
**Parent:** RILL-B01 · **Depends on:** None

**Research / exploration:** Explore Coinbase Exchange subscriptions.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Explore Coinbase Exchange wire protocol in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B01.2 — 01.2 — Inspect `matches`, `last_match` and `heartbeat` message types

**Type:** task · **Size:** S · **Milestone:** M1 · **Due:** 2026-10-20T11:00:00+05:30
**Budget:** 40 minutes · **Labels:** devex, execution:human, ingestion, scope:mvp, size:S, testing
**Parent:** RILL-B01 · **Depends on:** RILL-B01.1

**Research / exploration:** Inspect `matches`, `last_match` and `heartbeat` message types.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Explore Coinbase Exchange wire protocol in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B01.3 — 01.3 — Verify BTC-USD public connection and capture evidence

**Type:** task · **Size:** S · **Milestone:** M1 · **Due:** 2026-10-20T11:40:00+05:30
**Budget:** 40 minutes · **Labels:** devex, execution:human, ingestion, scope:mvp, size:S
**Parent:** RILL-B01 · **Depends on:** RILL-B01.2

**Research / exploration:** Verify BTC-USD public connection and capture evidence.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Explore Coinbase Exchange wire protocol in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B01.4 — 01.4 — Establish local Go/Task workflow

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T12:00:00+05:30
**Budget:** 20 minutes · **Labels:** devex, documentation, execution:human, ingestion, scope:mvp, size:XS, testing
**Parent:** RILL-B01 · **Depends on:** RILL-B01.3

**Research / exploration:** Establish local Go/Task workflow.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Explore Coinbase Exchange wire protocol in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B02 — Block 02 — Define the source adapter and normalize matches

**Type:** parent · **Size:** L · **Milestone:** M1 · **Due:** 2026-10-20T15:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, ingestion, testing
**Parent:** None · **Depends on:** RILL-B01

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Source-agnostic domain models, JSON parsing, exact trade IDs, maker/taker side, RFC3339 time, validation and error classification.

**Goal and evidence:** Show malformed-input fixtures and normalize an Exchange `match` without leaking wire-specific fields into analytics; exclude `last_match` explicitly.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B02.1 — 02.1 — Study domain models and maker-side semantics

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T13:20:00+05:30
**Budget:** 20 minutes · **Labels:** backend, execution:human, ingestion, scope:mvp, size:XS, study/research, testing
**Parent:** RILL-B02 · **Depends on:** RILL-B01

**Research / exploration:** Study domain models and maker-side semantics.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Define the source adapter and normalize matches in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B02.2 — 02.2 — Investigate integer trade IDs, decimal prices and timestamps

**Type:** task · **Size:** M · **Milestone:** M1 · **Due:** 2026-10-20T14:20:00+05:30
**Budget:** 60 minutes · **Labels:** backend, execution:human, ingestion, scope:mvp, size:M, testing
**Parent:** RILL-B02 · **Depends on:** RILL-B02.1

**Research / exploration:** Investigate integer trade IDs, decimal prices and timestamps.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Define the source adapter and normalize matches in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B02.3 — 02.3 — Define adapter output and failure classifications

**Type:** task · **Size:** S · **Milestone:** M1 · **Due:** 2026-10-20T14:50:00+05:30
**Budget:** 30 minutes · **Labels:** backend, execution:human, ingestion, scope:mvp, size:S, testing
**Parent:** RILL-B02 · **Depends on:** RILL-B02.2

**Research / exploration:** Define adapter output and failure classifications.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Define the source adapter and normalize matches in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B02.4 — 02.4 — Test valid and malformed Exchange messages

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T15:00:00+05:30
**Budget:** 10 minutes · **Labels:** backend, execution:human, ingestion, scope:mvp, size:XS, testing
**Parent:** RILL-B02 · **Depends on:** RILL-B02.3

**Research / exploration:** Test valid and malformed Exchange messages.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Define the source adapter and normalize matches in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B03 — Block 03 — Build a bounded, cancellable live pipeline

**Type:** parent · **Size:** L · **Milestone:** M1 · **Due:** 2026-10-20T18:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, ingestion, concurrency
**Parent:** None · **Depends on:** RILL-B02

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Context cancellation, channels, goroutine lifecycle, socket receive buffers versus trade queues, backpressure, drop/block trade-offs.

**Goal and evidence:** Demonstrate shutdown, bounded capacity and measured overload; justify a starting queue size rather than accepting 1,024 on faith.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B03.1 — 03.1 — Study goroutine lifecycle and cancellation

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T16:20:00+05:30
**Budget:** 20 minutes · **Labels:** backend, concurrency, execution:human, ingestion, scope:mvp, size:XS, study/research
**Parent:** RILL-B03 · **Depends on:** RILL-B02

**Research / exploration:** Study goroutine lifecycle and cancellation.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build a bounded, cancellable live pipeline in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B03.2 — 03.2 — Investigate socket buffers versus bounded trade queues

**Type:** task · **Size:** M · **Milestone:** M1 · **Due:** 2026-10-20T17:20:00+05:30
**Budget:** 60 minutes · **Labels:** backend, concurrency, execution:human, ingestion, scope:mvp, size:M
**Parent:** RILL-B03 · **Depends on:** RILL-B03.1

**Research / exploration:** Investigate socket buffers versus bounded trade queues.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build a bounded, cancellable live pipeline in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B03.3 — 03.3 — Compare blocking and dropping under synthetic bursts

**Type:** task · **Size:** S · **Milestone:** M1 · **Due:** 2026-10-20T17:50:00+05:30
**Budget:** 30 minutes · **Labels:** backend, concurrency, execution:human, ingestion, scope:mvp, size:S
**Parent:** RILL-B03 · **Depends on:** RILL-B03.2

**Research / exploration:** Compare blocking and dropping under synthetic bursts.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build a bounded, cancellable live pipeline in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B03.4 — 03.4 — Verify graceful shutdown and capacity behavior

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T18:00:00+05:30
**Budget:** 10 minutes · **Labels:** backend, concurrency, execution:human, ingestion, scope:mvp, size:XS, testing
**Parent:** RILL-B03 · **Depends on:** RILL-B03.3

**Research / exploration:** Verify graceful shutdown and capacity behavior.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build a bounded, cancellable live pipeline in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B04 — Block 04 — Reproduce inputs with JSONL replay

**Type:** parent · **Size:** L · **Milestone:** M1 · **Due:** 2026-10-20T21:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, replay
**Parent:** None · **Depends on:** RILL-B03

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** JSONL framing, schema versioning, captured arrival time, determinism, replay clock semantics and adapter reuse.

**Goal and evidence:** Show that recorded `subscriptions`/`last_match`/`match`/`heartbeat` feed the same parser and repeated replay yields equal outcomes.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B04.1 — 04.1 — Study JSONL framing and replay clocks

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T19:20:00+05:30
**Budget:** 20 minutes · **Labels:** backend, documentation, execution:human, replay, scope:mvp, size:XS, study/research
**Parent:** RILL-B04 · **Depends on:** RILL-B03

**Research / exploration:** Study JSONL framing and replay clocks.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Reproduce inputs with JSONL replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B04.2 — 04.2 — Capture versioned Exchange raw messages

**Type:** task · **Size:** M · **Milestone:** M1 · **Due:** 2026-10-20T20:20:00+05:30
**Budget:** 60 minutes · **Labels:** backend, execution:human, replay, scope:mvp, size:M
**Parent:** RILL-B04 · **Depends on:** RILL-B04.1

**Research / exploration:** Capture versioned Exchange raw messages.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Reproduce inputs with JSONL replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B04.3 — 04.3 — Test same adapter in replay and live modes

**Type:** task · **Size:** S · **Milestone:** M1 · **Due:** 2026-10-20T20:50:00+05:30
**Budget:** 30 minutes · **Labels:** backend, execution:human, replay, scope:mvp, size:S, testing
**Parent:** RILL-B04 · **Depends on:** RILL-B04.2

**Research / exploration:** Test same adapter in replay and live modes.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Reproduce inputs with JSONL replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B04.4 — 04.4 — Verify deterministic repeated replay

**Type:** task · **Size:** XS · **Milestone:** M1 · **Due:** 2026-10-20T21:00:00+05:30
**Budget:** 10 minutes · **Labels:** backend, documentation, execution:human, replay, scope:mvp, size:XS
**Parent:** RILL-B04 · **Depends on:** RILL-B04.3

**Research / exploration:** Verify deterministic repeated replay.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Reproduce inputs with JSONL replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B05 — Block 05 — Implement event-time fixed windows

**Type:** parent · **Size:** L · **Milestone:** M2 · **Due:** 2026-10-21T12:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, analytics
**Parent:** None · **Depends on:** RILL-B04

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** UTC 10-second half-open windows, event-time watermark, two-second lateness heuristic, out-of-order data, immutable finalization.

**Goal and evidence:** Test boundaries, idle gaps and late-arrival exclusion; explain why heartbeats cannot advance event-time windows.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B05.1 — 05.1 — Study tumbling windows and half-open boundaries

**Type:** task · **Size:** XS · **Milestone:** M2 · **Due:** 2026-10-21T10:20:00+05:30
**Budget:** 20 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:XS, study/research, testing
**Parent:** RILL-B05 · **Depends on:** RILL-B04

**Research / exploration:** Study tumbling windows and half-open boundaries.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Implement event-time fixed windows in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B05.2 — 05.2 — Define per-symbol watermark and lateness policy

**Type:** task · **Size:** M · **Milestone:** M2 · **Due:** 2026-10-21T11:20:00+05:30
**Budget:** 60 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:M
**Parent:** RILL-B05 · **Depends on:** RILL-B05.1

**Research / exploration:** Define per-symbol watermark and lateness policy.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Implement event-time fixed windows in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B05.3 — 05.3 — Test out-of-order and future-timestamp events

**Type:** task · **Size:** S · **Milestone:** M2 · **Due:** 2026-10-21T11:50:00+05:30
**Budget:** 30 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:S
**Parent:** RILL-B05 · **Depends on:** RILL-B05.2

**Research / exploration:** Test out-of-order and future-timestamp events.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Implement event-time fixed windows in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B05.4 — 05.4 — Verify immutability of finalized windows

**Type:** task · **Size:** XS · **Milestone:** M2 · **Due:** 2026-10-21T12:00:00+05:30
**Budget:** 10 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:XS, testing
**Parent:** RILL-B05 · **Depends on:** RILL-B05.3

**Research / exploration:** Verify immutability of finalized windows.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Implement event-time fixed windows in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B06 — Block 06 — Add incremental rolling analytics

**Type:** parent · **Size:** L · **Milestone:** M2 · **Due:** 2026-10-21T15:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, analytics
**Parent:** None · **Depends on:** RILL-B05

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Ring buffers, one-second buckets, rolling 60-second VWAP and trades/sec, warm-up, staleness, complexity bounds.

**Goal and evidence:** Compare incremental results to a simple reference for out-of-order and long idle input.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B06.1 — 06.1 — Study incremental one-second bucket algorithms

**Type:** task · **Size:** XS · **Milestone:** M2 · **Due:** 2026-10-21T13:20:00+05:30
**Budget:** 20 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:XS, study/research
**Parent:** RILL-B06 · **Depends on:** RILL-B05

**Research / exploration:** Study incremental one-second bucket algorithms.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Add incremental rolling analytics in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B06.2 — 06.2 — Define rolling 60-second window and warm-up

**Type:** task · **Size:** M · **Milestone:** M2 · **Due:** 2026-10-21T14:20:00+05:30
**Budget:** 60 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:M
**Parent:** RILL-B06 · **Depends on:** RILL-B06.1

**Research / exploration:** Define rolling 60-second window and warm-up.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Add incremental rolling analytics in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B06.3 — 06.3 — Compare VWAP and intensity with reference results

**Type:** task · **Size:** S · **Milestone:** M2 · **Due:** 2026-10-21T14:50:00+05:30
**Budget:** 30 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:S
**Parent:** RILL-B06 · **Depends on:** RILL-B06.2

**Research / exploration:** Compare VWAP and intensity with reference results.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Add incremental rolling analytics in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B06.4 — 06.4 — Test long idle gaps and bounded retention

**Type:** task · **Size:** XS · **Milestone:** M2 · **Due:** 2026-10-21T15:00:00+05:30
**Budget:** 10 minutes · **Labels:** analytics, backend, execution:human, scope:mvp, size:XS
**Parent:** RILL-B06 · **Depends on:** RILL-B06.3

**Research / exploration:** Test long idle gaps and bounded retention.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Add incremental rolling analytics in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B07 — Block 07 — Account for every match and bound state

**Type:** parent · **Size:** L · **Milestone:** M2 · **Due:** 2026-10-21T18:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, reliability, testing
**Parent:** None · **Depends on:** RILL-B06

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Terminal-outcome invariants, finite numeric checks, bounded dedup, IDs and precision, queue overflow, state limits.

**Goal and evidence:** Reconcile accepted/incorporated/rejected/pending counts after drain; distinguish local drops from unknowable upstream misses.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B07.1 — 07.1 — Define accepted-event conservation invariant

**Type:** task · **Size:** XS · **Milestone:** M2 · **Due:** 2026-10-21T16:15:00+05:30
**Budget:** 15 minutes · **Labels:** backend, execution:human, reliability, scope:mvp, size:XS, study/research, testing
**Parent:** RILL-B07 · **Depends on:** RILL-B06

**Research / exploration:** Define accepted-event conservation invariant.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Account for every match and bound state in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B07.2 — 07.2 — Study ID deduplication and bounded eviction

**Type:** task · **Size:** M · **Milestone:** M2 · **Due:** 2026-10-21T17:15:00+05:30
**Budget:** 60 minutes · **Labels:** backend, execution:human, reliability, scope:mvp, size:M, testing
**Parent:** RILL-B07 · **Depends on:** RILL-B07.1

**Research / exploration:** Study ID deduplication and bounded eviction.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Account for every match and bound state in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B07.3 — 07.3 — Test validation, drops and finite aggregates

**Type:** task · **Size:** S · **Milestone:** M2 · **Due:** 2026-10-21T17:50:00+05:30
**Budget:** 35 minutes · **Labels:** backend, execution:human, reliability, scope:mvp, size:S, testing
**Parent:** RILL-B07 · **Depends on:** RILL-B07.2

**Research / exploration:** Test validation, drops and finite aggregates.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Account for every match and bound state in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B07.4 — 07.4 — Reconcile counters and enforce capacity limits

**Type:** task · **Size:** XS · **Milestone:** M2 · **Due:** 2026-10-21T18:00:00+05:30
**Budget:** 10 minutes · **Labels:** backend, documentation, execution:human, reliability, scope:mvp, size:XS, testing
**Parent:** RILL-B07 · **Depends on:** RILL-B07.3

**Research / exploration:** Reconcile counters and enforce capacity limits.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Account for every match and bound state in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B08 — Block 08 — Enable two products and fault-tolerant recovery

**Type:** parent · **Size:** L · **Milestone:** M2 · **Due:** 2026-10-21T21:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, ingestion, reliability
**Parent:** None · **Depends on:** RILL-B07

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Per-product sequence diagnostics, `heartbeat.last_trade_id`, exponential backoff and jitter, retry classification, resubscription, cancellation.

**Goal and evidence:** Test BTC/ETH independence; disconnects, stale heartbeat, subscription rejection and cancellation during backoff; no overlapping reader.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B08.1 — 08.1 — Study per-product `sequence` and heartbeat `last_trade_id`

**Type:** task · **Size:** S · **Milestone:** M2 · **Due:** 2026-10-21T19:30:00+05:30
**Budget:** 30 minutes · **Labels:** backend, execution:human, ingestion, reliability, scope:mvp, size:S, study/research
**Parent:** RILL-B08 · **Depends on:** RILL-B07

**Research / exploration:** Study per-product `sequence` and heartbeat `last_trade_id`.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Enable two products and fault-tolerant recovery in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B08.2 — 08.2 — Classify recoverable and fatal feed failures

**Type:** task · **Size:** M · **Milestone:** M2 · **Due:** 2026-10-21T20:20:00+05:30
**Budget:** 50 minutes · **Labels:** backend, execution:human, ingestion, reliability, scope:mvp, size:M
**Parent:** RILL-B08 · **Depends on:** RILL-B08.1

**Research / exploration:** Classify recoverable and fatal feed failures.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Enable two products and fault-tolerant recovery in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B08.3 — 08.3 — Explore backoff, jitter, resubscribe and cancellation

**Type:** task · **Size:** S · **Milestone:** M2 · **Due:** 2026-10-21T20:50:00+05:30
**Budget:** 30 minutes · **Labels:** backend, execution:human, ingestion, reliability, scope:mvp, size:S, testing
**Parent:** RILL-B08 · **Depends on:** RILL-B08.2

**Research / exploration:** Explore backoff, jitter, resubscribe and cancellation.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Enable two products and fault-tolerant recovery in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B08.4 — 08.4 — Test ETH-USD independence and reconnect gaps

**Type:** task · **Size:** XS · **Milestone:** M2 · **Due:** 2026-10-21T21:00:00+05:30
**Budget:** 10 minutes · **Labels:** backend, execution:human, ingestion, reliability, scope:mvp, size:XS, testing
**Parent:** RILL-B08 · **Depends on:** RILL-B08.3

**Research / exploration:** Test ETH-USD independence and reconnect gaps.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Enable two products and fault-tolerant recovery in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B09 — Block 09 — Use correctness and debugging reserve

**Type:** parent · **Size:** L · **Milestone:** M3 · **Due:** 2026-10-22T12:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, testing, reliability
**Parent:** None · **Depends on:** RILL-B08

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Table-driven adversarial fixtures, race detector, message ordering, duplicate and continuity warnings, deterministic fault injection.

**Goal and evidence:** Resolve or precisely record failures; avoid new features.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B09.1 — 09.1 — Assemble adversarial Exchange fixtures

**Type:** task · **Size:** S · **Milestone:** M3 · **Due:** 2026-10-22T10:30:00+05:30
**Budget:** 30 minutes · **Labels:** execution:human, reliability, scope:mvp, size:S, study/research, testing
**Parent:** RILL-B09 · **Depends on:** RILL-B08

**Research / exploration:** Assemble adversarial Exchange fixtures.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Use correctness and debugging reserve in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B09.2 — 09.2 — Study deterministic fault injection

**Type:** task · **Size:** S · **Milestone:** M3 · **Due:** 2026-10-22T10:55:00+05:30
**Budget:** 25 minutes · **Labels:** execution:human, reliability, scope:mvp, size:S, testing
**Parent:** RILL-B09 · **Depends on:** RILL-B09.1

**Research / exploration:** Study deterministic fault injection.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Use correctness and debugging reserve in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B09.3 — 09.3 — Run tests and Go race detector

**Type:** task · **Size:** M · **Milestone:** M3 · **Due:** 2026-10-22T11:45:00+05:30
**Budget:** 50 minutes · **Labels:** execution:human, reliability, scope:mvp, size:M, testing
**Parent:** RILL-B09 · **Depends on:** RILL-B09.2

**Research / exploration:** Run tests and Go race detector.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Use correctness and debugging reserve in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B09.4 — 09.4 — Debug and document unresolved correctness issues

**Type:** task · **Size:** XS · **Milestone:** M3 · **Due:** 2026-10-22T12:00:00+05:30
**Budget:** 15 minutes · **Labels:** documentation, execution:human, reliability, scope:mvp, size:XS, testing
**Parent:** RILL-B09 · **Depends on:** RILL-B09.3

**Research / exploration:** Debug and document unresolved correctness issues.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Use correctness and debugging reserve in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B10 — Block 10 — Measure engine behavior and define browser contract

**Type:** parent · **Size:** L · **Milestone:** M3 · **Due:** 2026-10-22T15:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, performance
**Parent:** None · **Depends on:** RILL-B09

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Message/sec versus trades/sec, observed event age, monotonic local processing latency, bounded histograms and immutable snapshots.

**Goal and evidence:** Provide an example snapshot showing transport health, per-product freshness, sequence uncertainty, queue saturation and local drops.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B10.1 — 10.1 — Define message/trade and processing rates

**Type:** task · **Size:** XS · **Milestone:** M3 · **Due:** 2026-10-22T13:20:00+05:30
**Budget:** 20 minutes · **Labels:** backend, execution:human, performance, scope:mvp, size:XS, study/research
**Parent:** RILL-B10 · **Depends on:** RILL-B09

**Research / exploration:** Define message/trade and processing rates.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Measure engine behavior and define browser contract in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B10.2 — 10.2 — Study latency histograms and staleness semantics

**Type:** task · **Size:** M · **Milestone:** M3 · **Due:** 2026-10-22T14:15:00+05:30
**Budget:** 55 minutes · **Labels:** backend, execution:human, performance, scope:mvp, size:M
**Parent:** RILL-B10 · **Depends on:** RILL-B10.1

**Research / exploration:** Study latency histograms and staleness semantics.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Measure engine behavior and define browser contract in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B10.3 — 10.3 — Design immutable snapshots and health diagnostics

**Type:** task · **Size:** S · **Milestone:** M3 · **Due:** 2026-10-22T14:45:00+05:30
**Budget:** 30 minutes · **Labels:** backend, documentation, execution:human, performance, scope:mvp, size:S, testing
**Parent:** RILL-B10 · **Depends on:** RILL-B10.2

**Research / exploration:** Design immutable snapshots and health diagnostics.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Measure engine behavior and define browser contract in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B10.4 — 10.4 — Validate snapshot completeness against UI needs

**Type:** task · **Size:** XS · **Milestone:** M3 · **Due:** 2026-10-22T15:00:00+05:30
**Budget:** 15 minutes · **Labels:** backend, execution:human, performance, scope:mvp, size:XS
**Parent:** RILL-B10 · **Depends on:** RILL-B10.3

**Research / exploration:** Validate snapshot completeness against UI needs.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Measure engine behavior and define browser contract in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B11 — Block 11 — Publish HTTP/SSE safely

**Type:** parent · **Size:** L · **Milestone:** M3 · **Due:** 2026-10-22T18:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, backend, integration
**Parent:** None · **Depends on:** RILL-B10

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** SSE semantics, slow-consumer isolation, coalescing, cancellation and bounded fan-out.

**Goal and evidence:** Show that slow browsers cannot cause unbounded state or block trade ingestion.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B11.1 — 11.1 — Study HTTP SSE and browser EventSource

**Type:** task · **Size:** XS · **Milestone:** M3 · **Due:** 2026-10-22T16:20:00+05:30
**Budget:** 20 minutes · **Labels:** backend, execution:human, integration, scope:mvp, size:XS, study/research
**Parent:** RILL-B11 · **Depends on:** RILL-B10

**Research / exploration:** Study HTTP SSE and browser EventSource.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Publish HTTP/SSE safely in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B11.2 — 11.2 — Explore cancellation and slow-client isolation

**Type:** task · **Size:** M · **Milestone:** M3 · **Due:** 2026-10-22T17:20:00+05:30
**Budget:** 60 minutes · **Labels:** backend, execution:human, integration, scope:mvp, size:M
**Parent:** RILL-B11 · **Depends on:** RILL-B11.1

**Research / exploration:** Explore cancellation and slow-client isolation.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Publish HTTP/SSE safely in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B11.3 — 11.3 — Test bounded fan-out and coalescing

**Type:** task · **Size:** S · **Milestone:** M3 · **Due:** 2026-10-22T17:50:00+05:30
**Budget:** 30 minutes · **Labels:** backend, execution:human, integration, scope:mvp, size:S, testing
**Parent:** RILL-B11 · **Depends on:** RILL-B11.2

**Research / exploration:** Test bounded fan-out and coalescing.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Publish HTTP/SSE safely in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B11.4 — 11.4 — Verify health/snapshot reconnect behavior

**Type:** task · **Size:** XS · **Milestone:** M3 · **Due:** 2026-10-22T18:00:00+05:30
**Budget:** 10 minutes · **Labels:** backend, execution:human, integration, scope:mvp, size:XS
**Parent:** RILL-B11 · **Depends on:** RILL-B11.3

**Research / exploration:** Verify health/snapshot reconnect behavior.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Publish HTTP/SSE safely in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B12 — Block 12 — Build the dashboard within frontend-only permission

**Type:** parent · **Size:** L · **Milestone:** M3 · **Due:** 2026-10-22T21:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, frontend, integration
**Parent:** None · **Depends on:** RILL-B11

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** React/TypeScript, EventSource, bounded visualization histories, accessible stale/error states.

**Goal and evidence:** Codex builds only under `web/`; label Coinbase Exchange; human reviews contract and owns every Go change.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B12.1 — 12.1 — Define typed frontend contract and view states

**Type:** task · **Size:** XS · **Milestone:** M3 · **Due:** 2026-10-22T19:15:00+05:30
**Budget:** 15 minutes · **Labels:** execution:human, frontend, integration, scope:mvp, size:XS, study/research
**Parent:** RILL-B12 · **Depends on:** RILL-B11

**Research / exploration:** Define typed frontend contract and view states.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build the dashboard within frontend-only permission in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B12.2 — 12.2 — Codex builds allowed frontend under `web/`

**Type:** task · **Size:** M · **Milestone:** M3 · **Due:** 2026-10-22T20:20:00+05:30
**Budget:** 65 minutes · **Labels:** execution:codex-frontend, frontend, integration, scope:mvp, size:M
**Parent:** RILL-B12 · **Depends on:** RILL-B12.1

**Research / exploration:** Codex builds allowed frontend under `web/`.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build the dashboard within frontend-only permission in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B12.3 — 12.3 — Verify BTC/ETH, feed warnings and metric displays

**Type:** task · **Size:** S · **Milestone:** M3 · **Due:** 2026-10-22T20:45:00+05:30
**Budget:** 25 minutes · **Labels:** execution:human, frontend, integration, scope:mvp, size:S, testing
**Parent:** RILL-B12 · **Depends on:** RILL-B12.2

**Research / exploration:** Verify BTC/ETH, feed warnings and metric displays.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build the dashboard within frontend-only permission in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B12.4 — 12.4 — Check responsive behavior and type/build results

**Type:** task · **Size:** XS · **Milestone:** M3 · **Due:** 2026-10-22T21:00:00+05:30
**Budget:** 15 minutes · **Labels:** execution:human, frontend, integration, scope:mvp, size:XS
**Parent:** RILL-B12 · **Depends on:** RILL-B12.3

**Research / exploration:** Check responsive behavior and type/build results.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Build the dashboard within frontend-only permission in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B13 — Block 13 — Integrate and debug live plus replay

**Type:** parent · **Size:** L · **Milestone:** M4 · **Due:** 2026-10-23T12:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, integration, testing
**Parent:** None · **Depends on:** RILL-B12

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Contract tests, fault isolation, reproducible run commands, health under reconnect/shutdown.

**Goal and evidence:** Compare fixture → analytics → SSE → dashboard and demonstrate live BTC/ETH.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B13.1 — 13.1 — Compare replay results through API and dashboard

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T10:25:00+05:30
**Budget:** 25 minutes · **Labels:** execution:human, integration, scope:mvp, size:S, study/research, testing
**Parent:** RILL-B13 · **Depends on:** RILL-B12

**Research / exploration:** Compare replay results through API and dashboard.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Integrate and debug live plus replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B13.2 — 13.2 — Test live feed and two-symbol views

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T10:50:00+05:30
**Budget:** 25 minutes · **Labels:** execution:human, integration, scope:mvp, size:S, testing
**Parent:** RILL-B13 · **Depends on:** RILL-B13.1

**Research / exploration:** Test live feed and two-symbol views.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Integrate and debug live plus replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B13.3 — 13.3 — Exercise reconnect, stale heartbeat and slow clients

**Type:** task · **Size:** M · **Milestone:** M4 · **Due:** 2026-10-23T11:45:00+05:30
**Budget:** 55 minutes · **Labels:** execution:human, integration, scope:mvp, size:M, testing
**Parent:** RILL-B13 · **Depends on:** RILL-B13.2

**Research / exploration:** Exercise reconnect, stale heartbeat and slow clients.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Integrate and debug live plus replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B13.4 — 13.4 — Debug integration and record run commands

**Type:** task · **Size:** XS · **Milestone:** M4 · **Due:** 2026-10-23T12:00:00+05:30
**Budget:** 15 minutes · **Labels:** documentation, execution:human, integration, scope:mvp, size:XS, testing
**Parent:** RILL-B13 · **Depends on:** RILL-B13.3

**Research / exploration:** Debug integration and record run commands.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Integrate and debug live plus replay in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B14 — Block 14 — Benchmark individual-match throughput

**Type:** parent · **Size:** L · **Milestone:** M4 · **Due:** 2026-10-23T15:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, performance, testing
**Parent:** None · **Depends on:** RILL-B13

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Go benchmarks, latency distributions, allocations, Little's law, synthetic burst/saturation testing.

**Goal and evidence:** Report match decoding separately from analytics and queue; do not invent Exchange (different API) batch workloads or equate synthetic with live throughput.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B14.1 — 14.1 — Study Go benchmarks and offered-load testing

**Type:** task · **Size:** XS · **Milestone:** M4 · **Due:** 2026-10-23T13:20:00+05:30
**Budget:** 20 minutes · **Labels:** documentation, execution:human, performance, scope:mvp, size:XS, study/research, testing
**Parent:** RILL-B14 · **Depends on:** RILL-B13

**Research / exploration:** Study Go benchmarks and offered-load testing.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Benchmark individual-match throughput in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B14.2 — 14.2 — Benchmark individual match decoding and analytics

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T14:05:00+05:30
**Budget:** 45 minutes · **Labels:** execution:human, performance, scope:mvp, size:S, testing
**Parent:** RILL-B14 · **Depends on:** RILL-B14.1

**Research / exploration:** Benchmark individual match decoding and analytics.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Benchmark individual-match throughput in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B14.3 — 14.3 — Compare queue policies under bursty load

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T14:45:00+05:30
**Budget:** 40 minutes · **Labels:** documentation, execution:human, performance, scope:mvp, size:S, testing
**Parent:** RILL-B14 · **Depends on:** RILL-B14.2

**Research / exploration:** Compare queue policies under bursty load.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Benchmark individual-match throughput in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B14.4 — 14.4 — Record throughput, allocations and latency evidence

**Type:** task · **Size:** XS · **Milestone:** M4 · **Due:** 2026-10-23T15:00:00+05:30
**Budget:** 15 minutes · **Labels:** documentation, execution:human, performance, scope:mvp, size:XS, testing
**Parent:** RILL-B14 · **Depends on:** RILL-B14.3

**Research / exploration:** Record throughput, allocations and latency evidence.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Benchmark individual-match throughput in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B15 — Block 15 — Profile and stress fault tolerance

**Type:** parent · **Size:** L · **Milestone:** M4 · **Due:** 2026-10-23T18:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, performance, reliability
**Parent:** None · **Depends on:** RILL-B14

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** CPU and heap profiling, goroutine leaks, slow-consumer faults, reconnect storms, buffered backpressure and soak testing.

**Goal and evidence:** Produce bounded-heap and overload evidence, and record actual bottlenecks.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B15.1 — 15.1 — Study CPU and heap profiling with pprof

**Type:** task · **Size:** XS · **Milestone:** M4 · **Due:** 2026-10-23T16:20:00+05:30
**Budget:** 20 minutes · **Labels:** execution:human, performance, reliability, scope:mvp, size:XS, study/research, testing
**Parent:** RILL-B15 · **Depends on:** RILL-B14

**Research / exploration:** Study CPU and heap profiling with pprof.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Profile and stress fault tolerance in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B15.2 — 15.2 — Inspect allocation and CPU bottleneck hypotheses

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T16:50:00+05:30
**Budget:** 30 minutes · **Labels:** execution:human, performance, reliability, scope:mvp, size:S
**Parent:** RILL-B15 · **Depends on:** RILL-B15.1

**Research / exploration:** Inspect allocation and CPU bottleneck hypotheses.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Profile and stress fault tolerance in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B15.3 — 15.3 — Run bounded soak and failure drills

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T17:35:00+05:30
**Budget:** 45 minutes · **Labels:** execution:human, performance, reliability, scope:mvp, size:S
**Parent:** RILL-B15 · **Depends on:** RILL-B15.2

**Research / exploration:** Run bounded soak and failure drills.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Profile and stress fault tolerance in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B15.4 — 15.4 — Verify goroutine/memory bounds under overload

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T18:00:00+05:30
**Budget:** 25 minutes · **Labels:** execution:human, performance, reliability, scope:mvp, size:S
**Parent:** RILL-B15 · **Depends on:** RILL-B15.3

**Research / exploration:** Verify goroutine/memory bounds under overload.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Profile and stress fault tolerance in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B16 — Block 16 — Document evidence and checkpoint

**Type:** parent · **Size:** L · **Milestone:** M4 · **Due:** 2026-10-23T21:00:00+05:30
**Budget:** Rollup 120 minutes · **Labels:** scope:mvp, type:block, size:L, documentation, testing
**Parent:** None · **Depends on:** RILL-B15

Core two-hour learning block; parent rolls up 120 minutes of child research and evidence.

**Learn/look up:** Reproducibility, README/Task commands, feed delivery limits, connection state, replay assumptions, known uncertainties.

**Goal and evidence:** Mark MVP gates honestly at Friday checkpoint; use up to six optional weekend blocks for remaining core work only.

**Execution rules:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

See [todo.md](../todo.md) for architecture decisions, Coinbase Exchange protocol references, and validation limits.

### RILL-B16.1 — 16.1 — Review all acceptance evidence

**Type:** task · **Size:** M · **Milestone:** M4 · **Due:** 2026-10-23T20:00:00+05:30
**Budget:** 60 minutes · **Labels:** documentation, execution:human, scope:mvp, size:M, study/research, testing
**Parent:** RILL-B16 · **Depends on:** RILL-B15

**Research / exploration:** Review all acceptance evidence.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Document evidence and checkpoint in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B16.2 — 16.2 — Document startup, replay and protocol assumptions

**Type:** task · **Size:** S · **Milestone:** M4 · **Due:** 2026-10-23T20:35:00+05:30
**Budget:** 35 minutes · **Labels:** documentation, execution:human, scope:mvp, size:S, testing
**Parent:** RILL-B16 · **Depends on:** RILL-B16.1

**Research / exploration:** Document startup, replay and protocol assumptions.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Document evidence and checkpoint in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B16.3 — 16.3 — Record measured limitations and unresolved gates

**Type:** task · **Size:** XS · **Milestone:** M4 · **Due:** 2026-10-23T20:55:00+05:30
**Budget:** 20 minutes · **Labels:** documentation, execution:human, scope:mvp, size:XS, testing
**Parent:** RILL-B16 · **Depends on:** RILL-B16.2

**Research / exploration:** Record measured limitations and unresolved gates.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Document evidence and checkpoint in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-B16.4 — 16.4 — Make Friday checkpoint and optional weekend decision

**Type:** task · **Size:** XS · **Milestone:** M4 · **Due:** 2026-10-23T21:00:00+05:30
**Budget:** 5 minutes · **Labels:** documentation, execution:human, scope:mvp, size:XS, testing
**Parent:** RILL-B16 · **Depends on:** RILL-B16.3

**Research / exploration:** Make Friday checkpoint and optional weekend decision.

**Evidence:** Record what you learned, the chosen trade-off, test or observed result and any unresolved question. Consult Document evidence and checkpoint in todo.md. Avoid following an implementation script; author all Go changes yourself.

**Ownership:** Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S01 — Order book reconstruction

**Type:** stretch · **Size:** XL · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:XL, backend, ingestion, study/research
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Order book reconstruction: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S02 — Docker local deployment

**Type:** stretch · **Size:** L · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:L, devex
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Docker local deployment: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S03 — Cloud deployment

**Type:** stretch · **Size:** XL · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:XL, devex, study/research
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Cloud deployment: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S04 — Additional analytics and honest volume imbalance

**Type:** stretch · **Size:** L · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:L, analytics, study/research
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Additional analytics and honest volume imbalance: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S05 — Add SOL-USD

**Type:** stretch · **Size:** S · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:S, ingestion, testing
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Add SOL-USD: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S06 — Arbitrary-instant rolling queries

**Type:** stretch · **Size:** XL · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:XL, analytics, study/research
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Arbitrary-instant rolling queries: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S07 — Continuous bounded recording

**Type:** stretch · **Size:** L · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:L, replay, reliability
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Continuous bounded recording: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

### RILL-S08 — Evidence-backed performance optimization

**Type:** stretch · **Size:** L · **Milestone:** Unscheduled · **Due:** Not scheduled
**Budget:** Not budgeted · **Labels:** scope:stretch, size:L, performance
**Parent:** None · **Depends on:** RILL-B16

Unscheduled stretch item. Not part of 32 core hours or the optional 12-hour contingency. Consider only after all MVP acceptance gates pass.

Evidence-backed performance optimization: Investigate requirements, alternatives and validation evidence before implementation.

Human owns Go backend and all Git operations by default. This user-authorized planning-file edit is an explicit limited exception; no further agent Git writes, CI, hooks, PRs or Linear-triggered repository mutations. Codex may mentor/review and implement frontend only under web/ in Block 12.

## Optional contingency blocks (not created or scheduled)

C1–C6: six possible two-hour blocks across Oct 24–25. Allocate only after the Friday checkpoint, in order of blockers. No automatic expansion of issues or hours; defer stretch goals. Sunday final checkpoint is 23:59 IST.

## Acceptance and authoritative source

[todo.md](../todo.md) is the source of truth for Exchange protocol assumptions, source adapter boundary, failure matrix, retry classification/backoff, socket-versus-Go-buffer design, queue saturation/drop trade-offs, watermark and replay semantics, and test evidence. This inventory summarizes learning targets rather than prescribing backend implementation.

**No Linear project, milestones, labels, or issues have been created yet.**
