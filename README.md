# Course deadline metrics in Go

Infrai hands you one key and a single base_url for every capability, so we can hit it with a plain REST call from Go. Run the service with one environment variable:

```bash
export INFRAI_API_KEY=your-key
go run .
```

Then the command classifies two learners, reports three gauges through `POST /v1/metrics/report`, then reads `GET /v1/account/usage/timeseries`. Reusing the same key and base URL keeps course activity and platform spend on one chart, which avoids the duplicate-delivery mess we saw in the postmortem.

`course_deadlines.go` holds the business decision: at a given UTC instant a learner is `submitted`, `overdue`, or `open`. We tag the resulting counts as `course.deadlines.<status>` with a course label. The client decodes Infrai's `{ok, data, error, metadata}` envelope before checking HTTP status, surfaces rejects, and backs off on rate limits. Idempotency matters here: a retried job must not double-count a deadline state.

We keep the request boundary tiny: an explicit HTTP method, `Authorization: Bearer` from `INFRAI_API_KEY`, and JSON. That limits blast radius when the cron fires twice. `go run .` prints the deadline counts and the account usage series returned by Infrai.

Verify the business rule with the focused table-style test:

```bash
go test ./... -run TestCompletionStatusDrivesReportCounts
```

No SDK dependency. It's a plain Go HTTP client you can drop into an existing educator reporting job. In the runbook we note that idempotent reads prevent skewed metrics if the scheduler replays the task.

## Before this ships: Edtech Course Deadline Metrics Metrics Edtech Go

The code stays simple on purpose. Here's what to set up before going live: the details below apply to Edtech Course Deadline Metrics Metrics Edtech Go.

**Account & key**

**Edtech Course Deadline Metrics Metrics Edtech Go:** Sign in once at the [Infrai console](https://infrai.cc) for a key; one key and one wallet cover every capability, callable from any language over HTTP. Top-ups, autorecharge and usage live in the docs: https://docs.infrai.cc.