# Reviewing human-written Go

This workflow operates under [AGENTS.md](../../AGENTS.md). It grants no implementation permission.

1. Establish the work block, intended behavior, and relevant diff or files. Read surrounding code and existing tests before drawing conclusions.
2. Inspect correctness, event accounting, concurrency, cancellation, resource ownership, bounded memory, backpressure, recovery, and reproducibility. For ingestion/replay changes, check the Coinbase policies in [todo.md](../../todo.md): batch membership, snapshot exclusion, trade-time semantics, heartbeat health and uncertain sequence scope. Treat performance concerns as hypotheses until evidence supports them.
3. Run relevant existing verification when useful. Prefer documented Task commands after inspecting their commands and dependencies. Do not run formatting, auto-fix, or generation steps that modify backend source. Never write missing tests or benchmarks; describe the cases the human should add.
4. Prioritize findings using exactly **Blocking** (correctness, safety, serious reliability), **Important** (meaningful engineering improvements), and **Optional** (style, alternatives, unproven optimizations). Omit categories with no findings.
5. For each finding, identify the location, triggering condition, consequence, and supporting evidence. Explain the conceptual fix direction and how the human can reproduce, investigate, and validate it. Do not provide or apply code fixes.
6. Report checks run and outcomes, checks not run, and any residual uncertainty. If no actionable issues are found, say so without claiming the code is proven correct. Avoid unrelated refactors and preference-based criticism.

Reusable request:

> Review my human-written changes for Block N against its definition of done and AGENTS.md. Inspect the diff and relevant surrounding files, prioritize actual findings as Blocking, Important, or Optional, and explain their consequences and how I should investigate and validate corrections. Run appropriate existing checks if useful. Do not edit backend files, generate Go code, write tests or benchmarks, stage changes, or create commits.
