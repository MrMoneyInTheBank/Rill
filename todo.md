# Rill v0.1.0 — learning and delivery plan

**Objective:** Build and explain a single-process Go engine that ingests live crypto trades, computes correct bounded-memory analytics, replays recorded inputs deterministically, and exposes measured behavior through a small dashboard. All Go implementation, tests, and benchmarks belong to the human developer; Codex mentors and reviews, and implements only the frontend under `web/` in Block 12. See [AGENTS.md](AGENTS.md).

**Core checkpoint:** Friday, **October 23, 2026 at 11:59 PM IST**. **Optional final deadline:** Sunday, **October 25, 2026 at 11:59 PM IST**, only if core acceptance work needs more time. Roadmap preparation is October 19; core blocks run October 20–23.

**Roadmap-only automation:** Linear tracks scope, ownership, sizes and deadlines. All Git operations are human-only. No agent staging, commits, branches, tags, pushes, pull requests, hooks, repository integration, or GitHub Actions/CI setup is authorized. Linear completion must never trigger code execution or repository changes.

**Budget:** **16 core sequential blocks × 2 hours = 32 focused hours**, plus up to **6 optional contingency blocks × 2 hours = 12 hours** on October 24–25 (44 hours maximum), including reading, design, implementation, testing, debugging, AI interaction, frontend integration, and documentation. The overview uses four relative workdays, four blocks each; breaks are additional calendar time. This is a demanding eight focused hours per day, not a promise that every unfamiliar concept will fit. Preserve any completed blocks and account for migration work within remaining blocks; changing providers does not automatically reset the budget. Stop at each timebox and use the fallback policy below.

## MVP acceptance criteria

- Live Coinbase Exchange BTC-USD match ingestion: public subscription, heartbeat health, validated trade updates, cancellation, and basic reconnect with visible possible data gaps.
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

**Read official Exchange documentation, not Advanced Trade docs:** [Overview](https://docs.cdp.coinbase.com/exchange/websocket-feed/overview), [Channels](https://docs.cdp.coinbase.com/exchange/websocket-feed/channels), [Best practices](https://docs.cdp.coinbase.com/exchange/websocket-feed/best-practices), [Errors](https://docs.cdp.coinbase.com/exchange/websocket-feed/errors), [Rate limits](https://docs.cdp.coinbase.com/exchange/websocket-feed/rate-limits). Validate live examples in Block 1 rather than trusting illustrative payloads.

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
- Process individual matches in received order. Keep the trade's exchange event time distinct from local receive time; no Advanced Trade message exchange event timestamp exists here. Preserve wire-level sequence and trade-ID diagnostics separately.
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

### Block 1 — Establish access and a small local workflow

**Objective:** Verify public Coinbase connectivity and create the smallest runnable Go project, managed through Task.

**Concepts to learn:** Go modules, configuration boundaries, public feed subscriptions, reproducible local commands.

**Architectural guidance:** Use the public Advanced Trade feed without credentials. Verify it with a trusted WebSocket client before building the Go lifecycle. Subscribe to the chosen trade product and heartbeat, and save bounded inbound examples; an unreachable endpoint should not consume the next three blocks.

**Tasks:**

- [ ] 20 min: Read the provider notes, recheck official rate limits, and check Go/Task/library compatibility. Note Bun as a Block 12 prerequisite.
- [ ] 40 min: Verify the subscription responses, BTC-USD match trades and heartbeat traffic; retain a few complete inbound messages with receive times and observed snapshot behavior. Close the probe afterward. Record access findings and unresolved quota details.
- [ ] 40 min: Initialize the module, write the minimal entry point yourself, and add working Task list/build/run/check tasks. Avoid empty future packages.
- [ ] 20 min: Run the skeleton and checks, confirm secret/binary ignore coverage, and record any access blocker with the next action.

**Definition of done:** Local commands work and there is observed BTC-USD access evidence, or an explicitly unresolved external blocker with offline work ready to proceed. Only the former satisfies the live-access gate.

**Decision to make:** Confirm the Advanced Trade public spot feed and update-only policy; defer optional JWT authentication.

**Resources:** [Go modules](https://go.dev/doc/tutorial/create-module), [Coinbase overview](https://docs.cdp.coinbase.com/coinbase-app/advanced-trade-apis/websocket/websocket-overview), [Task guide](https://taskfile.dev/docs/guide/).

### Block 2 — Normalize trades and test the boundary

**Objective:** Convert provider data into a small, validated domain model independently of the live socket.

**Concepts to learn:** Typed JSON decoding, exact identifiers, time parsing, validation, table-driven testing.

**Architectural guidance:** Keep provider fields at the boundary. Decode the message and then each nested event/trade, applying the update-only policy. Preserve event, server-send and receive timestamps with different meanings. Prefer explicit types over a generic event framework; control inputs and snapshots do not enter analytics.

**Tasks:**

- [ ] 20 min: Select domain fields, string-number parsing, rejection/exclusion reasons and timestamp meanings from the reference contract.
- [ ] 60 min: Implement message decoding and match-trade normalization in received order; classify snapshots, heartbeat, subscription/errors and unknown inputs.
- [ ] 30 min: Test malformed numeric strings, bad times, exact IDs, multiple events/trades, snapshot exclusion and heartbeat-only input.
- [ ] 10 min: Run `task check` and annotate what was intentionally excluded.

**Definition of done:** Every eligible trade in a saved batch normalizes correctly; invalid records have reasons; snapshots/heartbeat cannot create trades or advance time; IDs retain their identity.

**Conceptual checkpoint:** Why is an exchange timestamp unsuitable for measuring in-process execution duration?

**Resources:** [Coinbase AsyncAPI schema](https://docs.cdp.coinbase.com/api-reference/advanced-trade-api/advanced-trade-asyncapi.json), [JSON package](https://pkg.go.dev/encoding/json), [Go monotonic clocks](https://pkg.go.dev/time#hdr-Monotonic_Clocks).

### Block 3 — Connect the live pipeline with bounded concurrency

**Objective:** Receive normalized BTC-USD events through a controlled lifecycle that stops reliably.

**Concepts to learn:** Goroutines, channel ownership, contexts, blocking, backpressure, WebSocket message/control lifecycles.

**Architectural guidance:** One reader and one processor are sufficient. Learn cancellation before adding more components. Implement both subscriptions and the heartbeat watchdog on one connection; until Block 8, a watchdog failure terminates/reports that connection. The library handles WebSocket control frames; Coinbase application heartbeat are a separate health signal.

**Tasks:**

- [ ] 20 min: Read context/read/close guidance and sketch who creates, cancels, closes and waits for each resource.
- [ ] 60 min: Implement connect/subscribe/read, heartbeat monitoring and bounded update admission with minimal counters; subscribe promptly after connecting.
- [ ] 30 min: Exercise cancellation, full queues, a missing heartbeat and shutdown draining; use a deliberately slow consumer to expose batch pressure.
- [ ] 10 min: Check a short live run and run existing checks; leave retry work as a clearly identified Block 8 task.

**Definition of done:** Updates reach the processor; batches cannot bypass queue limits; cancellation leaks no reader or blocked send. Missing heartbeat and quiet trading produce distinct health states.

**Decision to make:** Can you explain what is lost when the queue is full, and what blocking would change?

**Resources:** [coder/websocket](https://github.com/coder/websocket), [Go pipelines and cancellation](https://go.dev/blog/pipelines), [context](https://pkg.go.dev/context).

### Block 4 — Build replay before market analytics

**Objective:** Make recorded and synthetic inputs usable offline through the exact same normalization boundary.

**Concepts to learn:** Streaming file I/O, fixture design, deterministic order, logical versus execution time.

**Architectural guidance:** Replay complete recorded messages through the live decoder and admission policy. A synchronous driver isolates correctness from scheduler/queue effects and will call the same analytics entry point as live ingestion. Preserve nested order and session boundaries; do not build a generalized source framework or durable recording service.

**Tasks:**

- [ ] 20 min: Define the JSONL record wrapper and prepare saved-message fixtures plus labeled synthetic snapshot/update/heartbeat cases.
- [ ] 60 min: Implement bounded line-by-line reading, the shared decoder/normalizer handoff and `task replay`; add pacing only after the core works.
- [ ] 30 min: Replay twice; compare order, rejection and snapshot-exclusion counts. Test malformed/oversize records, session boundaries and explicit EOF behavior.
- [ ] 10 min: Document fresh-state requirements, receive-time provenance and which operational values cannot be deterministic.

**Definition of done:** Replay needs no credentials, preserves message/nested order and exact IDs, and reproduces classification/normalization outcomes. It is ready to feed the analytics added next.

**Conceptual checkpoint:** Faster replay may reproduce aggregates while failing to simulate a live connection's timing and pressure. Explain why.

**Resources:** [Go testing](https://pkg.go.dev/testing).

### Block 5 — Implement fixed windows with a finalization policy

**Objective:** Produce correct 10-second count, volume and VWAP results, including out-of-order and late inputs.

**Concepts to learn:** Half-open intervals, event-time watermarks, bounded lateness, immutable results.

**Architectural guidance:** Implement the documented per-symbol watermark now, alongside fixed windows. Avoid wall-clock timers inside analytics. Track only open state and bounded recent results; empty gaps and partial tails are different from complete windows.

**Tasks:**

- [ ] 20 min: Work through boundary examples on paper, including `t = W`, a finalized interval and a timestamp jump.
- [ ] 60 min: Implement aggregates, late exclusion, finalization and zero-volume semantics; connect replay to this same processor.
- [ ] 30 min: Write fixtures with hand-calculated VWAP, exact boundaries, admissible reordering, late drops and an EOF partial tail.
- [ ] 10 min: Compare repeated replay results and run checks; document the two-second default's limitations.

**Definition of done:** Expected count/volume/VWAP match; finalized output remains unchanged after later arrivals; late events increment the counter exactly once.

**Decision to make:** Accept the shared strict lateness boundary or deliberately document an alternative before implementing it.

### Block 6 — Add incremental rolling analytics

**Objective:** Compute 60-second rolling VWAP and trade intensity with bounded storage and clear as-of semantics.

**Concepts to learn:** Circular buffers, incremental aggregates, expiration, computational complexity versus measurements.

**Architectural guidance:** Use the one-second bucket contract and maintain totals as buckets enter/leave the rolling interval. Avoid rescanning raw trade history. A large time jump must cost work bounded by retained state, not elapsed seconds. Keep pending buckets out of completed rolling totals.

**Tasks:**

- [ ] 20 min: Draw interval membership, pending buckets, warm-up and ring reuse across the 60-second boundary.
- [ ] 60 min: Implement bounded buckets and rolling totals without changing fixed-window semantics.
- [ ] 30 min: Check boundaries, wraparound, reordering, zero volume and long gaps; heartbeat-only input and EOF must leave the watermark unchanged.
- [ ] 10 min: Replay known fixtures and explain the observed operation costs without claiming a speedup yet.

**Definition of done:** Rolling values match a simple human-calculated/reference result at chosen horizons; state stays bounded and stale input does not move logical time.

**Conceptual checkpoint:** Which question does a rolling value answer when the newest trade is several minutes old?

### Block 7 — Close accounting gaps and enforce limits

**Objective:** Make every admission's outcome explainable and every retained collection bounded.

**Concepts to learn:** Accounting invariants, bounded deduplication, partial failure, numerical safety.

**Architectural guidance:** Treat incorporation as one state transition across both analytic views. Add the bounded exact-ID cache and explicit rejection precedence; a duplicate is counted once, not also counted as late. Expose the limits of cache eviction and unknown upstream loss.

**Tasks:**

- [ ] 15 min: Audit per-message versus per-trade counters, last_match exclusions and accepted outcomes; list every growing collection.
- [ ] 60 min: Implement bounded duplicate handling, missing caps and finite-aggregate checks; ensure rejected events cannot partially update windows or watermarks.
- [ ] 35 min: Test duplicate repeats/eviction, equal IDs on distinct symbols where supported by the model, numeric rejection, queue overflow and shutdown accounting.
- [ ] 10 min: Verify the accounting equation after a drain; record limits and deduplication guarantees.

**Definition of done:** Each admitted event has one outcome; pending reaches zero after normal shutdown; no collection grows with unbounded event count or elapsed time.

**Decision to make:** Why is a detected local drop count unable to establish the number of trades missed upstream?

### Block 8 — Add a second symbol and basic recovery

**Objective:** Maintain independent BTC-USD and ETH-USD analytics and recover a lost live connection.

**Concepts to learn:** State partitioning, retry classification, backoff, reconnect gaps.

**Architectural guidance:** Keep one symbol-state owner and one public connection. Reconnect serially with cancellable exponential backoff/jitter (start near one second, cap near 30 seconds). Re-establish both subscriptions promptly; no authentication step is needed. Retain analytics/dedupe, reset session diagnostics and ignore reconnect last_match. Retry network/heartbeat failures; surface invalid products/subscriptions, and back off on throttling. Sequence anomalies flag uncertain continuity, not automatic reconnect storms or exact missing-trade counts.

**Tasks:**

- [ ] 30 min: Enable two products with independent analytic state; compare sequence behavior across products/channels without assuming its scope.
- [ ] 50 min: Implement reconnect, session diagnostics and continuity indicators; reset backoff after a stable subscription with healthy heartbeat.
- [ ] 30 min: Test interleaved products, heartbeat loss, reconnect last_match overlap and cancellation during retry; verify public BTC/ETH updates.
- [ ] 10 min: Run checks and note that missed market events are not backfilled.

**Definition of done:** One product cannot close another's windows; reconnect restores trades and heartbeat without duplicate readers; snapshots do not alter retained state. Unresolved sequence scope is documented rather than guessed.

**Decision to make:** Keep SOL deferred if adding it would consume validation time; two symbols satisfy the minimum multi-symbol gate.

### Block 9 — Spend the correctness and debugging reserve

**Objective:** Stabilize the engine before exposing it to a browser. This block introduces no new feature scope.

**Concepts to learn:** Reproducible failures, reference results, race detection, evidence-driven debugging.

**Architectural guidance:** Tiny explicit fixtures catch semantic errors more clearly than live observation. Compare counts exactly and numeric aggregates with justified tolerances. A passing race run covers exercised paths, not every possible execution.

**Tasks:**

- [ ] 30 min: Assemble earlier regression cases, including individual match messages, reconnect last_match, heartbeat-only idle, future outliers and repeated replay.
- [ ] 25 min: Add `task test:race`, run checks and reproduce lifecycle/overload failures without needing the provider to misbehave.
- [ ] 50 min: Debug and correct the highest-priority problems yourself; Codex may review and suggest investigations.
- [ ] 15 min: Re-run affected cases and record unresolved blockers before scheduling UI work.

**Definition of done:** Core correctness fixtures, accounting and exercised race checks pass. Remaining defects are explicit; the block is not silently spent on features.

**Conceptual checkpoint:** Can you reproduce a late-event or cancellation failure offline with a small input?

**Resources:** [Go race detector](https://go.dev/doc/articles/race_detector).

### Block 10 — Instrument the engine and define the API contract

**Objective:** Measure the engine and specify the minimal snapshot the frontend will consume.

**Concepts to learn:** Throughput versus latency, histogram precision, observational overhead, immutable snapshots.

**Architectural guidance:** Report decoded match trades/sec before validation/admission as ingestion throughput, and incorporated trades/sec as processing throughput. Keep message rate, last_match exclusions and rolling market intensity separate. Measure local receive-to-processing-completion duration including decoding and queue wait; it excludes provider batching/network transit. Use a bounded histogram for approximate p50/p95/p99 and report population, sample count/range/overflow. Do not present p99 as stable with sparse data. Snapshot copies must not race with the state owner.

**Tasks:**

- [ ] 20 min: Choose latency boundaries, histogram precision and reporting period; define ingestion/processing rates and accepted, incorporated, rejection, queue and reconnect metrics.
- [ ] 55 min: Implement metrics and a coherent snapshot with heartbeat age/transport health, per-product staleness, continuity warnings, ignored last_match messages, backlog, analytics and capped histories.
- [ ] 30 min: Document `/healthz` and `/events`, units, schema version, string IDs, unavailable values and full-snapshot reconnect behavior. Compare snapshot fields to the dashboard checklist.
- [ ] 15 min: Validate a fixture snapshot and metric counts; verify that instrumentation has bounded memory.

**Definition of done:** A sample snapshot covers every required display and distinguishes market time from operational rates/durations. Zero samples appear unavailable, not as invented percentiles.

**Decision to make:** Fix a useful approximate histogram precision now; a telemetry stack is outside this budget.

**Resources:** [Go monotonic clocks](https://pkg.go.dev/time#hdr-Monotonic_Clocks), [Go diagnostics](https://go.dev/doc/diagnostics).

### Block 11 — Deliver snapshots through HTTP and SSE

**Objective:** Make engine state available to a browser without letting a client slow market processing.

**Concepts to learn:** HTTP streaming, SSE framing/flushing, client lifecycle, bounded fan-out.

**Architectural guidance:** SSE suits one-way updates and browser reconnection; upstream remains WebSocket. Browser WebSockets become useful if bidirectional messaging is later required. Publish at most twice per second with a one-slot latest-snapshot mailbox per client; replace stale pending snapshots and count coalescing separately from market drops. Use capped clients and bounded writes/cancellation. Keep the API local and let Vite proxy `/events` and `/healthz` during development. [SSE behavior](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events)

**Tasks:**

- [ ] 20 min: Read HTTP flush/cancellation and SSE guidance; confirm the snapshot contract and connection limits.
- [ ] 60 min: Implement health and SSE routes, controlled publishing, write deadlines, disconnect cleanup and HTTP shutdown.
- [ ] 30 min: Verify a normal client, reconnect, five-client cap and a deliberately slow/stalled reader while processing replay traffic.
- [ ] 10 min: Confirm complete snapshots on reconnect and that no event-log catch-up guarantee is implied.

**Definition of done:** Updates arrive without buffering indefinitely; a slow or disconnected client cannot block the processor or cause unbounded queues/goroutines.

**Conceptual checkpoint:** A skipped dashboard update and a skipped market event are different losses. Can the counters distinguish them?

**Resources:** [net/http](https://pkg.go.dev/net/http), [Using SSE](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events).

### Block 12 — Generate the focused dashboard with Codex

**Objective:** Build the permitted frontend under `web/` against the established API, within two hours including review.

**Concepts to learn:** Typed API consumption, render cadence, connection state, accessible data presentation.

**Architectural guidance:** This is the first frontend implementation block. Codex may use Vite, React, TypeScript, Tailwind CSS and Bun under `web/`. Prefer simple SVG price movement, compact tables/cards and a single EventSource. Pair selection uses the backend's product keys and filters subscribed data locally; label the source Coinbase Exchange. Show transport health separately from market freshness; omit trade-direction coloring unless explicitly labeled as maker side. The human handles any required backend contract change.

**Tasks:**

- [ ] 15 min: Give Codex the contract and sample snapshot; agree on the minimal responsive dark layout and bounded display histories.
- [ ] 65 min: Have Codex scaffold and implement pair selection, tape, price movement, fixed/rolling analytics, ingestion/processing throughput, latency, health and all drop counters, including empty/stale/reconnecting states.
- [ ] 25 min: Inspect mobile/desktop rendering, units, unavailable values, ID precision and EventSource cleanup; run frontend type/build checks.
- [ ] 15 min: Human adds Task wrappers for web development/check/build, and lists integration issues for Block 13.

**Definition of done:** The dashboard renders a contract fixture cleanly and its checks pass. All generated application code is under `web/`; Go remains human-authored.

**Decision to make:** Stop visual polish once every required value and failure state is readable.

**Resources:** [Vite with Bun](https://bun.sh/guides/ecosystem/vite), [Tailwind with Vite](https://tailwindcss.com/docs/installation/using-vite).

### Block 13 — Integrate and debug the complete application

**Objective:** Spend the second full debugging reserve on the real backend/browser boundary and operating failures.

**Concepts to learn:** Contract debugging, layered fault isolation, user-visible staleness, coordinated shutdown.

**Architectural guidance:** Start with replay to establish known results, then live input to test the network. AI fixes remain confined to frontend code; backend corrections remain human tasks. Freeze feature scope after this integration gate.

**Tasks:**

- [ ] 25 min: Start backend and web via Task; compare fixture values across processor output, SSE and displayed results.
- [ ] 25 min: Exercise BTC/ETH switching, public updates, quiet-market/heartbeat health, browser reconnect, upstream disconnect and process cancellation.
- [ ] 55 min: Fix contract, rendering, lifecycle or accounting defects; test a slow browser alongside normal processing.
- [ ] 15 min: Re-run affected checks and record actual run/configuration commands for the README.

**Definition of done:** Live and replay data reach all required dashboard views; stale/disconnected state is honest; startup and shutdown are reproducible.

**Conceptual checkpoint:** Identify which layer is wrong before changing either side of an API mismatch.

### Block 14 — Establish a performance baseline

**Objective:** Produce repeatable measurements before attempting optimization.

**Concepts to learn:** Benchmark scope, allocations/op, throughput, latency distributions, workload validity.

**Architectural guidance:** Benchmark normalized analytics separately from match-message decode-plus-processing and the live-style queue. Keep fixture creation/reporting outside timed work. Go benchmark `ns/op` is not p99. Use controlled offered load with queue metrics and completion latency to show saturation; a synchronous driver alone hides queueing. Provider batching affects delivery cadence independently of Rill's measured internal latency. No arbitrary low-latency target is a release requirement.

**Tasks:**

- [ ] 20 min: Fix a workload spanning windows/products with unique IDs, varying trades per batch and bursty arrivals; record machine, Go version and limits.
- [ ] 45 min: Write Go benchmarks yourself; add `task bench` around benchmark/benchmem commands and collect repeated baseline runs.
- [ ] 40 min: Exercise the existing queue with controlled replay load and record events/sec, latency samples/percentiles, queue occupancy and drops at each offered rate.
- [ ] 15 min: Save a compact baseline report with commands, workload size, sample count, variation and scope; do not optimize yet.

**Definition of done:** Someone can reproduce the workload and understand what was timed; throughput and tail latency are separately reported with allocations and loss.

**Decision to make:** Do results show CPU saturation, queueing, allocation pressure or insufficient samples? A hypothesis is not a conclusion.

**Resources:** [Go benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks), [Go diagnostics](https://go.dev/doc/diagnostics).

### Block 15 — Profile and test reliability under sustained load

**Objective:** Explain resource use and check that sustained operation respects the declared bounds.

**Concepts to learn:** CPU/heap profiles, allocation churn versus retained heap, profiler overhead, soak testing.

**Architectural guidance:** Prefer profile files from the existing benchmarks to adding a public profiling endpoint. Keep baseline, race and profiled runs separate because instrumentation changes timings. Optimize only an evidenced bottleneck if the reliability work is already complete; documenting a bottleneck is an acceptable outcome.

**Tasks:**

- [ ] 20 min: Add Task CPU/heap profile wrappers using Go test profiling flags and learn the relevant `go tool pprof` views.
- [ ] 30 min: Collect and inspect CPU/heap profiles; connect top costs to the data path and distinguish retained memory from total allocations.
- [ ] 45 min: Run a bounded 15–20 minute soak while exercising overload, a stalled SSE reader and shutdown; compare heap/goroutine/backlog trends, not just one sample.
- [ ] 25 min: Resolve a blocking reliability defect if found and repeat its check; otherwise document the baseline explanation and one justified future improvement.

**Definition of done:** Profiles exist with interpretable findings; retained application state stays within limits after warm-up; drops and shutdown behave as documented. No unsupported performance claim is made.

**Conceptual checkpoint:** Lower allocation count does not automatically mean lower tail latency. What does this workload's evidence show?

**Resources:** [runtime/pprof](https://pkg.go.dev/runtime/pprof), [Go diagnostics](https://go.dev/doc/diagnostics).

### Block 16 — Finish documentation and run acceptance

**Objective:** Deliver a reproducible, explainable v0.1.0 or an accurately labeled incomplete checkpoint.

**Concepts to learn:** Release evidence, documented limitations, useful reproducibility instructions.

**Architectural guidance:** No new MVP features now. README documentation should connect pipeline behavior to commands and evidence. Git operations and GitHub Actions/CI are outside this roadmap; retain local verification.

**Tasks:**

- [ ] 60 min: Use the final fix reserve on outstanding blockers and targeted rechecks. If none remain, rehearse the clean-start path and repeat the weakest acceptance evidence.
- [ ] 35 min: Complete README prerequisites, versions, credential-free public startup, configuration/Task commands, heartbeat/snapshot/clock/replay policies and benchmark interpretation; include known limits.
- [ ] 20 min: Run the final checklist below, inspect the diff for scope/artifacts/secrets, and record pass/fail evidence with no credentials.
- [ ] 5 min: Mark each acceptance gate honestly and note follow-up work. The human handles any Git commit/tag/push separately; agents perform none.

**Definition of done:** The final checklist passes and the human can explain the design/results. Otherwise list missing gates; a deadline does not make an unfinished MVP complete.

**Decision to make:** Publish the measured limitations in the README rather than hiding them behind a feature list.

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
