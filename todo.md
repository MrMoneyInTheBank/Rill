# Rill v0.1.0 — learning and delivery plan

**Objective:** Build and explain a single-process Go engine that ingests live crypto trades, computes correct bounded-memory analytics, replays recorded inputs deterministically, and exposes measured behavior through a small dashboard. All Go implementation, tests, and benchmarks belong to the human developer; Codex mentors and reviews, and implements only the frontend under `web/` in Block 12. See [AGENTS.md](AGENTS.md).

**Core checkpoint:** Friday, **October 23, 2026 at 11:59 PM IST**. **Optional final deadline:** Sunday, **October 25, 2026 at 11:59 PM IST**, only if core acceptance work needs more time. Roadmap preparation is October 19; core blocks run October 20–23.

**Roadmap-only automation:** Linear tracks scope, ownership, sizes and deadlines. Git operations remain human-owned except for this explicitly authorized planning-file update. No other agent staging, commits, branches, tags, pushes, pull requests, hooks, repository integration, or GitHub Actions/CI setup is authorized. Linear completion must never trigger code execution or repository changes.

**Budget:** **16 core sequential blocks × 2 hours = 32 focused hours**, plus up to **6 optional contingency blocks × 2 hours = 12 hours** on October 24–25 (44 hours maximum), including reading, design, implementation, testing, debugging, AI interaction, frontend integration, and documentation. The overview uses four relative workdays, four blocks each; breaks are additional calendar time. This is a demanding eight focused hours per day, not a promise that every unfamiliar concept will fit. Preserve any completed blocks and account for migration work within remaining blocks; changing providers does not automatically reset the budget. Stop at each timebox and use the fallback policy below.

## MVP acceptance criteria

- Live Coinbase Exchange BTC-USD match ingestion: public subscription, heartbeat health, validated match messages, cancellation, and basic reconnect with visible possible data gaps.
- Immutable finalized 10-second windows: trade count, volume, VWAP; explicit event-time lateness policy and `late_events_dropped`.
- Incremental 60-second rolling VWAP and trades/second, evaluated at one-second event-time boundaries, with visible warm-up and staleness.
- Independent BTC-USD and ETH-USD analytics; bounded state and explicit event accounting throughout.
- JSONL replay through the same normalization and analytics, with repeatable results for documented input/order/clock assumptions.
- Responsive dark dashboard: pair selection, bounded trade tape and recent price plot, both analytics, ingestion/processing throughput, latency, health, and rejection/drop counters.
- Evidence from correctness tests, race detection, benchmarks, CPU/heap profiles, overload and slow-browser checks; measured p50/p95/p99 where samples support them.
- Task-based local commands and a README sufficient to run, test, replay, and interpret results without credentials in the repository.

## Budget overview

Every row is one two-hour block. The detailed blocks below are the work queue; architectural notes are reference material, not additional assignments.

| Block | Relative day | Focus | Observable result |
| --- | --- | --- | --- |
| 1 | Day 1 | Setup, Task, public feed access | Local skeleton and verified trade/heartbeat subscriptions |
| 2 | Day 1 | Normalization and validation | Tested provider-independent trade model |
| 3 | Day 1 | Live lifecycle and concurrency | Bounded, cancellable ingestion |
| 4 | Day 1 | Deterministic replay foundation | Offline input through the same normalizer |
| 5 | Day 2 | Fixed windows and lateness | Immutable, tested 10-second results |
| 6 | Day 2 | Rolling analytics | Bounded incremental 60-second results |
| 7 | Day 2 | Accounting, duplicates, limits | Reconciled counters and hard bounds |
| 8 | Day 2 | Multi-symbol and reconnect | BTC+ETH isolation and recovery |
| 9 | Day 3 | Correctness and debugging reserve | Adversarial fixtures and race checks pass |
| 10 | Day 3 | Metrics and browser contract | Measurable engine and agreed snapshot fields |
| 11 | Day 3 | HTTP/SSE | Slow-client-safe snapshot delivery |
| 12 | Day 3 | Frontend — Codex implementation | Dashboard under `web/` |
| 13 | Day 4 | Integration and debugging reserve | Live/replay dashboard and failure handling |
| 14 | Day 4 | Benchmark baseline | Reproducible throughput/latency/allocation evidence |
| 15 | Day 4 | Profiling and reliability | CPU/heap analysis and bounded-memory soak |
| 16 | Day 4 | Final fixes, docs, acceptance | Honest v0.1.0 handoff |

Blocks 9 and 13 reserve four full hours for verification/debugging. Block 16 reserves another hour for fixes; each implementation block also includes checks. Up to six additional two-hour contingency blocks may be used on October 24–25 only for unfinished MVP gates; do not reset completed work. Optional work cannot consume these reserves.

**Stretch goals, after every MVP gate passes:** order book reconstruction; Docker; cloud deployment; additional analytics such as volume imbalance; SOL-USD; arbitrary-instant rolling queries; continuous recording; measured optimizations. None is needed to complete this plan. GitHub Actions/CI is excluded from this roadmap; any future Git work belongs to the human.

## Working assumptions and architecture reference

### Coinbase Exchange WebSocket — protocol facts and learning questions (October 19, 2026)

**Read official Exchange documentation, not Exchange (different API) docs:** [Overview](https://docs.cdp.coinbase.com/exchange/websocket-feed/overview), [Channels](https://docs.cdp.coinbase.com/exchange/websocket-feed/channels), [Best practices](https://docs.cdp.coinbase.com/exchange/websocket-feed/best-practices), [Errors](https://docs.cdp.coinbase.com/exchange/websocket-feed/errors), [Rate limits](https://docs.cdp.coinbase.com/exchange/websocket-feed/rate-limits). Validate live examples in Block 1 rather than trusting illustrative payloads.

- Public endpoint: `wss://ws-feed.exchange.coinbase.com`. Public market data needs no account/JWT. Subscribe to **`matches`** and **`heartbeat`** (singular) for `BTC-USD`, then `ETH-USD`. Exchange messages use a top-level `type`, not Exchange's `channel/events/trades` message. Acknowledge `subscriptions` and inspect errors.
- Exchange `matches` sends **one match per message**. The first trade-like message after subscribing is `last_match`, the last known execution: treat it as initial context, **not a newly observed live trade**, unless you deliberately justify a different policy. Do not refer to Exchange's snapshot/update batches.
- `match` includes `trade_id` (JSON integer), `sequence`, `product_id`, `price` and `size` (decimal strings), `time`, `side`, `maker_order_id` and `taker_order_id`. `side` is the **maker** side: a maker sell means aggressive buy. Decide a safe representation for IDs and for frontend JSON precision.
- Exchange `heartbeat` is **per product**, approximately every second, with `sequence`, `last_trade_id`, `product_id` and `time`. It provides a liveness/data-quality signal, not proof that all matches arrived. Compare WebSocket Ping/Pong, application heartbeat, and per-product market inactivity.
- `sequence` is per product and applies to feed messages, not just trades. Sequence gaps do **not** equal missing-trade counts. The `matches` channel can drop messages; use `last_trade_id` and sequence observations as **continuity diagnostics**, not automatic backfill or exactly-once guarantees. Read the docs on message loss, out-of-order and duplicate sequences; verify in your own capture.
- `ticker` is not the same as `matches`; `full` adds order lifecycle complexity and is a stretch goal. Do not assume the two Coinbase APIs share limits, error formats, heartbeat semantics or delivery guarantees.

**Investigate:** What does a WebSocket subscription acknowledgment mean? How can an idle product coexist with a healthy connection? Can a REST trade-history query help investigate gaps, and what completeness guarantees would still be missing? What are the real rate/connection limits? How do malformed, unexpected, oversized, and unsupported messages differ? What can and cannot be inferred from a skipped sequence number?

Recommended general-purpose WebSocket client to evaluate: [coder/websocket](https://github.com/coder/websocket). Learn the Exchange application protocol yourself, not raw WebSocket framing.

### Source adapter, fault tolerance and recovery: design questions

A **source adapter** translates provider-specific wire messages into a provider-independent trade model plus separate health/control/diagnostic signals. It is a semantic boundary, not permission to build a generic plugin framework. Learn when an interface helps, and when a concrete adapter is simpler. Replay should exercise the **same decoder and normalizer**. Exchange-specific `sequence`, heartbeat `last_trade_id`, connection session and wire errors belong in diagnostics, not in every analytics formula.

Create a **failure matrix** for dial/DNS/TLS errors, subscribe timeout/rejection, unknown message types, malformed JSON, oversized frames, silent connection, stale per-product heartbeat, WebSocket close/read failure, sequence anomalies, local queue saturation, replay corruption, shutdown during retry and stalled SSE clients. For each, decide: detect, classify, record, retry or stop, expose health, and test. Distinguish transient transport failures, fatal configuration errors, and uncertain market-data continuity.

**Learn/look up:** exponential backoff, jitter, retry caps, reset-after-stable-connection, context-aware cancellation, reconnection/resubscription, goroutine ownership, read deadlines, Ping/Pong, and error wrapping/classification. An initial experiment might use a one-second retry base and a roughly 30-second cap, but these are tunable choices, not Coinbase requirements. A retry must not spawn overlapping readers; reconnect does not recover lost history automatically. Preserve analytic state when appropriate and reset session diagnostics.

### Buffers, backpressure and loss policy

Distinguish OS/socket receive buffers, WebSocket frame/message limits, decoded **trade** queue, bounded per-symbol analytics, and browser/SSE mailboxes. Coinbase's [best practices](https://docs.cdp.coinbase.com/exchange/websocket-feed/best-practices) discuss socket receive buffering and slow consumers; that does **not** imply an unbounded Go channel.

**Learn/look up:** arrival/service rates, burstiness, queue occupancy, Little's law, blocking sends, drop-newest/drop-oldest, nonblocking admission, reader starvation, slow-consumer disconnects and tail latency. **1,024 normalized trades** is a starting hypothesis, not a magic capacity. Compare alternatives with synthetic bursts; pick one with explicit local-loss accounting, bounded memory and heartbeat responsiveness. Distinguish locally dropped trades from an **unknown** number of upstream missing trades. Measure high-water marks, queue wait, drops and recovery.

### Rill's Exchange admission and health policies

- Classify `subscriptions`, `heartbeat`, `last_match`, `match`, errors and unknown types separately. Admit validated **`match`** only. Count initial `last_match` as ignored context; it must not seed live analytics or dedup unless justified and tested.
- Process individual matches in received order. Keep the trade's exchange event time distinct from local receive time; no Exchange (different API) message exchange event timestamp exists here. Preserve wire-level sequence and trade-ID diagnostics separately.
- A heartbeat watchdog (candidate five seconds per subscribed product) is a **local heuristic**; distinguish connection health, per-product heartbeat age and last-trade staleness. A healthy socket does not prove market-data completeness.
- Track bounded per-product sequence and heartbeat-last-trade diagnostics. On reconnect reset session baselines, retain analytic state and bounded dedup cache. An anomaly is a warning to investigate, not an exact lost-trade count or an automatic claim of backfill.

### Responsibilities and state ownership

Start with the following responsibilities; create each package only when its first real implementation needs it.

| Location | Responsibility and dependencies |
| --- | --- |
| `cmd/rill` | Configuration, live/replay selection, wiring, cancellation and shutdown |
| `internal/model` | Provider-independent trade and result meaning; no networking dependencies |
| `internal/ingest` | Coinbase subscriptions, heartbeat/session health, message decoding and match-trade normalization |
| `internal/analytics` | Fixed/rolling computations and event-time policy; no sockets or wall-clock reads |
| `internal/replay` | Stream recorded messages into the existing decoder, normalizer and processor; no second analytics implementation |
| `internal/stream` | Add only when coordinating admission, per-symbol state and snapshots warrants separation from `cmd/rill` |
| `internal/api` | HTTP/SSE transport of immutable snapshots; cannot mutate analytics |
| `web/` | Display the agreed contract; no replacement market-processing engine |

One upstream reader feeds one bounded channel. One processor owns the map of all symbol state. HTTP handlers consume immutable, size-bounded copies. This avoids locks around every window update and makes processing order inspectable. A synchronous pipeline is simpler but couples reads to computation; sharded processors may help measured CPU contention later. A goroutine per symbol adds ordering, lifecycle and snapshot complexity without current evidence of benefit.

Start with an ingress capacity of 1,024 normalized trades, not messages. On full live ingress, reject the newest trade and count `ingress_queue_full`; do not spawn a goroutine or grow storage to make room. Continue classifying the bounded message so every extracted update has an outcome and heartbeat remain readable. This sacrifices known local events to maintain a bounded reader path. Blocking instead can transfer pressure upstream and end the connection. Measure queue occupancy, drops and latency before changing this choice. Replay correctness mode calls the same processor synchronously, so execution speed cannot cause nondeterministic queue drops; a separate load exercise can test the live queue.

On shutdown, stop admissions, stop/close the source, then drain admitted events before stopping the processor. Give shutdown a bounded grace period (initially five seconds); abandoned admissions get `shutdown_dropped`. The source owns channel closure. Cancellation must unblock reads, sends, retry waits and HTTP handlers. Do not mark open windows final merely because the process exits.

### Time and analytic semantics

These are recommended v0.1 defaults to decide and test in Blocks 2, 5 and 6, not implementation instructions.

- Keep **trade event time**, **message server-send time**, **local receive UTC**, and **local monotonic durations** distinct. A each match has its own event time and captured receive timestamp. Reject missing/invalid trade times and malformed/non-positive/non-finite numeric strings. Preserve IDs as strings. `float64` prices/volumes are acceptable for this analytics exercise with documented tolerances, finite aggregate checks, and no financial-ledger precision claim. Choose decimal/fixed-point arithmetic later if evidence requires it.
- Per symbol, let `M` be the greatest timestamp of an incorporated trade and `W = M − 2 seconds`. Before the first such trade there is no watermark. Reject an arriving trade with `t < previous W`, incrementing `late_events_dropped`; equality is admissible. Validate, deduplicate and check numeric admissibility before advancing `M`. This deliberately stricter rule can reject a trade whose 10-second window is still open, so both analytics share one clear lateness boundary.
- Fixed intervals are UTC epoch-aligned, half-open `[start, start + 10 seconds)`. Incorporate eligible data before advancing finalization; an interval becomes final when its end is at or before `W`. Final results never change. VWAP is total price-times-size divided by total size; zero volume means unavailable, never zero price. Emit nonempty fixed windows; gaps are visible as missing intervals, not fabricated evidence of no trading.
- Rolling results cover exactly `[E − 60 seconds, E)`, with `E` equal to `W` rounded down to a whole second. Maintain one-second aggregate buckets and incremental totals; a fixed ring of 64 buckets per symbol suffices for this default horizon/lateness. Revalidate capacity if either parameter changes. Buckets at or after `E` are pending. Trade intensity is the interval count divided by 60; label warm-up until the horizon has advanced 60 seconds beyond the first incorporated event. Warm-up completion does not certify upstream completeness. This is an exact window at specified horizons; arbitrary-instant querying would require a different retention design. An event deque/reorder buffer is an alternative if that finer cadence becomes necessary.
- Idle trades, disconnects and replay EOF **freeze logical time**. Heartbeats and exchange event timestamps never advance watermarks. The browser may keep refreshing health, but event-as-of and analytics stay unchanged and visibly stale; rolling intensity does not decay according to wall time. Pending windows remain partial. Reconnection retains state and resumes event-time progress on eligible updates. Large jumps expire old state with bounded work and do not generate millions of empty intervals. A busy symbol never advances an idle symbol's watermark.
- Two seconds is a tunable starting assumption, not a Coinbase ordering guarantee or a consequence of its batching cadence. Max-observed timestamps cannot prove that earlier events have all arrived, and a future outlier can prematurely close windows. Initially reject a trade more than five seconds ahead of its captured receive UTC timestamp, with `future_timestamp_rejected`. Replay uses the saved receive timestamp and the same threshold. This guard relies on a reasonably synchronized local clock; smaller erroneous jumps can still cause loss. Measure observed lateness before tuning.

### Accounting, limits and replay contract

**Accepted** means a validated normalized match trade successfully admitted to the processor: queued in live mode or handed to the synchronous processor in replay. Acceptance is not a promise of analytics incorporation. Each admission gets one terminal outcome: incorporated, duplicate, late, numeric rejection, or shutdown drop. At a consistent checkpoint:

`accepted = incorporated + duplicate + late + numeric_rejected + shutdown_dropped + pending`

`pending` includes queued/in-flight admissions; after a successful drain it is zero. The incorporated outcome updates both analytic states, including pending buckets, or neither on a rejection. A coalesced browser snapshot need not display every incorporated trade. Invalid records, unsupported products, future timestamps, ignored last_match messages and queue-full events are accounted for before acceptance with separate counters. Heartbeats are not trade admissions. An undecodable message gets a message-error count; its unknown number of trades must not be invented. Keep upstream continuity warnings separate from counted local losses.

For duplicate IDs, start with an exact per-symbol cache of the most recent 10,000 incorporated IDs, deterministic oldest-inserted eviction, and `duplicates_dropped`. Keep it across reconnects, reset it on a new process/replay run. Detection is guaranteed only while an ID remains in the cache; an evicted or restarted ID can be counted again if it passes the time policy. ID uniqueness is an application assumption scoped to the chosen feed and symbol. Conflicting repeats warrant an integrity warning, not overwriting earlier analytics. This is bounded deduplication, not end-to-end exactly-once delivery. Reconnects neither recover missed upstream trades nor establish how many were missed.

Other initial caps: three configured symbols, 1 MiB WebSocket message/2 MiB JSONL line, 100 tape entries, 120 price points and 60 finalized windows per symbol, bounded latency histogram, five SSE clients with one pending snapshot each. These are starting engineering limits to test, not provider specifications. Aggregate numeric overflow must reject without partial mutation. Any future recording mode must also have file size/duration limits. Time horizons alone do not bound arbitrary trade lists or ID caches.

JSONL records contain a schema version, provider/feed identity, connection-session marker, local record order, captured receive UTC and one original inbound Exchange JSON message. Preserve event types, nested array order, trade strings and diagnostic metadata so live and replay share the same decoder and admission rules. The local record sequence is independent of Coinbase sequence values and trade IDs. Stream the file rather than loading it all. Include public trade/heartbeat messages and necessary control inputs; exclude outgoing authentication, headers, secrets or unrestricted session dumps. Small saved samples and human-authored synthetic fixtures are sufficient; an always-on recorder is outside the MVP. Existing trade-only fixtures, if any, need an explicit schema migration; do not silently reinterpret them.

Given identical records, order, configuration and numeric rules, a fresh replay reproduces aggregate results, terminal counters, snapshot-exclusion counts and partial/final status. It need not reproduce browser refreshes, operational throughput, latency, reconnect timing, or losses that occurred before recording. Do not sort messages or their nested trades: arrival order is part of the experiment. EOF does not force finalization; fixtures contain explicit later match trades when testing completed windows, and those trades also contribute to their own intervals. Compare ordered semantic results with documented floating-point tolerances; exclude local timings and unordered serialization. Paced replay sleeps between records but never uses that sleeping to advance analytic time. Replay does not run live heartbeat timeouts or reconnect loops; recorded session boundaries reset diagnostic context only. Measure replay performance with fresh per-run monotonic markers, never historical receive UTC timestamps. Reproducing live overload exactly would additionally require recorded admission decisions; that is deferred.

CLI-only replay may report its results and exit. Dashboard replay stops its source at EOF but keeps the HTTP service and final snapshot available until cancellation; otherwise a fast fixture could finish before the browser connects.

### Local development with Task

Use **Task (`go-task/task`)**, as in Jobit. The inspected `../jobit/Taskfile.yml` uses a discoverable default, `run`/`build`, and focused checks. Reuse the naming, not Jobit's application paths, CGO setting, hooks or script wrappers. The human adds a small Rill `Taskfile.yml` in Block 1 and grows it only when commands work. No Taskfile is created during this planning invocation. [Task guide](https://taskfile.dev/docs/guide/)

| Add when ready | Planned Task interface |
| --- | --- |
| Blocks 1–3 | `task` / `task --list`, `task run`, `task build`, `task fmt-check`, `task lint`, `task test`, `task check` |
| Blocks 4, 9 | `task replay -- <fixture>`, `task test:race` |
| Blocks 12–13 | `task web:dev`, `task web:check`, `task web:build` |
| Blocks 14–15 | `task bench`, `task profile:cpu`, `task profile:heap` |

Keep `check` source-preserving: formatting check, vet and tests. A human-only `task fmt` may format Go later; AI must not invoke it. Tasks wrap understandable Go/Bun commands and preserve exit codes. Start backend and frontend in two terminals; a cross-process supervisor or hot reload is unnecessary. Public live mode and replay must run without credentials. Document chosen endpoint/product configuration names and defaults when implemented; do not add unused secret variables. If authenticated access is introduced later, secrets stay in the human's local environment, never echoed by Task or exposed through `VITE_*`. No credentials are requested or inspected by Codex.

## The sixteen work blocks

Each checklist allocates all 120 minutes. Read only the resource sections needed for that block; the lists are not extra homework. All tasks involving Go code, tests or benchmarks are performed by the human.

### Block 1 — Explore Coinbase Exchange wire protocol

**Learn/look up:** Exchange feed overview, subscriptions, `matches`, first `last_match`, per-product `heartbeat`, message sequence, Go modules and Task.

**Goal and evidence:** Capture the actual subscribe acknowledgment, sample matches and heartbeat; verify BTC-USD rate without credentials; explain what is and isn't a fresh execution.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 2 — Define the source adapter and normalize matches

**Learn/look up:** Source-agnostic domain models, JSON parsing, exact trade IDs, maker/taker side, RFC3339 time, validation and error classification.

**Goal and evidence:** Show malformed-input fixtures and normalize an Exchange `match` without leaking wire-specific fields into analytics; exclude `last_match` explicitly.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 3 — Build a bounded, cancellable live pipeline

**Learn/look up:** Context cancellation, channels, goroutine lifecycle, socket receive buffers versus trade queues, backpressure, drop/block trade-offs.

**Goal and evidence:** Demonstrate shutdown, bounded capacity and measured overload; justify a starting queue size rather than accepting 1,024 on faith.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 4 — Reproduce inputs with JSONL replay

**Learn/look up:** JSONL framing, schema versioning, captured arrival time, determinism, replay clock semantics and adapter reuse.

**Goal and evidence:** Show that recorded `subscriptions`/`last_match`/`match`/`heartbeat` feed the same parser and repeated replay yields equal outcomes.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 5 — Implement event-time fixed windows

**Learn/look up:** UTC 10-second half-open windows, event-time watermark, two-second lateness heuristic, out-of-order data, immutable finalization.

**Goal and evidence:** Test boundaries, idle gaps and late-arrival exclusion; explain why heartbeats cannot advance event-time windows.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 6 — Add incremental rolling analytics

**Learn/look up:** Ring buffers, one-second buckets, rolling 60-second VWAP and trades/sec, warm-up, staleness, complexity bounds.

**Goal and evidence:** Compare incremental results to a simple reference for out-of-order and long idle input.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 7 — Account for every match and bound state

**Learn/look up:** Terminal-outcome invariants, finite numeric checks, bounded dedup, IDs and precision, queue overflow, state limits.

**Goal and evidence:** Reconcile accepted/incorporated/rejected/pending counts after drain; distinguish local drops from unknowable upstream misses.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 8 — Enable two products and fault-tolerant recovery

**Learn/look up:** Per-product sequence diagnostics, `heartbeat.last_trade_id`, exponential backoff and jitter, retry classification, resubscription, cancellation.

**Goal and evidence:** Test BTC/ETH independence; disconnects, stale heartbeat, subscription rejection and cancellation during backoff; no overlapping reader.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 9 — Use correctness and debugging reserve

**Learn/look up:** Table-driven adversarial fixtures, race detector, message ordering, duplicate and continuity warnings, deterministic fault injection.

**Goal and evidence:** Resolve or precisely record failures; avoid new features.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 10 — Measure engine behavior and define browser contract

**Learn/look up:** Message/sec versus trades/sec, observed event age, monotonic local processing latency, bounded histograms and immutable snapshots.

**Goal and evidence:** Provide an example snapshot showing transport health, per-product freshness, sequence uncertainty, queue saturation and local drops.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 11 — Publish HTTP/SSE safely

**Learn/look up:** SSE semantics, slow-consumer isolation, coalescing, cancellation and bounded fan-out.

**Goal and evidence:** Show that slow browsers cannot cause unbounded state or block trade ingestion.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 12 — Build the dashboard within frontend-only permission

**Learn/look up:** React/TypeScript, EventSource, bounded visualization histories, accessible stale/error states.

**Goal and evidence:** Codex builds only under `web/`; label Coinbase Exchange; human reviews contract and owns every Go change.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 13 — Integrate and debug live plus replay

**Learn/look up:** Contract tests, fault isolation, reproducible run commands, health under reconnect/shutdown.

**Goal and evidence:** Compare fixture → analytics → SSE → dashboard and demonstrate live BTC/ETH.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 14 — Benchmark individual-match throughput

**Learn/look up:** Go benchmarks, latency distributions, allocations, Little's law, synthetic burst/saturation testing.

**Goal and evidence:** Report match decoding separately from analytics and queue; do not invent Exchange (different API) batch workloads or equate synthetic with live throughput.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 15 — Profile and stress fault tolerance

**Learn/look up:** CPU and heap profiling, goroutine leaks, slow-consumer faults, reconnect storms, buffered backpressure and soak testing.

**Goal and evidence:** Produce bounded-heap and overload evidence, and record actual bottlenecks.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

### Block 16 — Document evidence and checkpoint

**Learn/look up:** Reproducibility, README/Task commands, feed delivery limits, connection state, replay assumptions, known uncertainties.

**Goal and evidence:** Mark MVP gates honestly at Friday checkpoint; use up to six optional weekend blocks for remaining core work only.

**Checkpoint:** Write down decisions, surprises and observed evidence; leave implementation details to your own research and Go code. Two focused hours; blockers move to reserved/optional time.

## Optional weekend contingency — October 24–25 (up to six 2-hour blocks)

Only schedule C1–C6 after the Friday Block 16 checkpoint. Prioritize unfinished Exchange protocol correctness and recovery; then replay/window correctness; then backpressure, dashboard integration, race checks, profiling and documentation. These are **not new features**. Record time spent and remaining blockers. If all acceptance gates pass Friday, stop Friday. Otherwise Sunday October 25 at 11:59 PM IST is the final checkpoint, and an incomplete gate must be reported honestly.

## Final verification checklist

- [ ] Public BTC-USD/ETH-USD updates and heartbeat observed without credentials; reconnect/cancellation checked; continuity uncertainty acknowledged.
- [ ] Fixed/rolling fixtures cover boundaries, ordering, lateness, duplicates, long gaps, stale symbols, zero volume and numeric rejection.
- [ ] Repeated message replays give the same results/counters, including ignored last_match messages and partial EOF; heartbeat-only input cannot advance analytics.
- [ ] Accepted-event accounting balances after draining; all state/history/client/message limits and overload outcomes are explicit.
- [ ] Relevant Task checks, race detection and frontend type/build checks pass; executed commands and limits of coverage are recorded.
- [ ] Live/replay UI, product selection, separate transport/freshness states and slow-client isolation checked at desktop/mobile sizes.
- [ ] Baseline includes workload, hardware/toolchain, throughput, latency/sample quality, allocations, CPU/heap findings and sustained-memory observations.
- [ ] README allows a new local start with Task; credentials and generated artifacts remain outside committed content.

## If progress falls behind

At each block boundary, record what passed and the next smallest failing case. Do not add hours invisibly. Use Block 9 for unfinished engine correctness and Block 13 for integration; the Block 16 hour is the last fix reserve. Shorten reading to the immediate question and remove abstraction work before removing tests.

Cut all stretch goals first, including deployment and optimization. Keep two symbols and defer SOL; simplify plots to a small SVG line and reduce styling, not required metrics. Keep one-second rolling cadence, two-Hz browser snapshots, two-terminal startup and fixture-based replay. Continuous capture, durable storage, catch-up replay for browsers, exact historical backfill and deployment remain outside the deadline.

If public feed access is blocked in Block 1, continue offline normalization/replay work and revisit the concrete network/provider blocker within scheduled integration time. Do not add account/JWT work without an evidenced need. Offline success does not satisfy live ingestion. If a mandatory gate still fails when the budget or revised deadline is reached, deliver a named incomplete checkpoint with evidence and a short continuation list; do not claim v0.1.0 or remove replay, rolling correctness, race testing or profiling to make the checklist look complete.

**Finished Rill v0.1.0 means:** all MVP gates and the final checklist pass, and the human can explain the connection lifecycle, state ownership, backpressure, event-time/EOF behavior, accounting and measured bottlenecks. It is a local analytics learning system with explicit loss and precision limits, not a claim of exchange completeness or a production trading platform.
