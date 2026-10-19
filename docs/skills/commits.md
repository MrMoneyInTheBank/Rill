# Human-only Git workflow

All Git operations belong to the human under [AGENTS.md](../../AGENTS.md). Linear is a roadmap, not an execution or Git automation system.

1. AI may inspect relevant files and read-only diffs, identify unrelated changes, and suggest logical commit boundaries.
2. AI may run existing source-preserving verification when requested; it reports failures for the human to fix.
3. The human alone stages, creates branches, commits, tags, opens pull requests, pushes, amends or otherwise manages Git history.
4. AI must not configure hooks, repository integrations or GitHub Actions/CI, or delegate Linear issues to autonomous coding agents.
5. A roadmap issue, review request or issue-status change never grants Git permission. A future explicit override must identify the particular Git mutation authorized.

Reusable request:

> Review my diff and suggest coherent commit boundaries and messages. Report relevant verification results. Leave all Git actions to me: do not stage, commit, branch, tag, push, create pull requests, configure hooks/CI, or change history. Do not write or edit backend code.
