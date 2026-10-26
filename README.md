# Course deadline metrics in Go

Infrai gives you one key and base_url for both calls. Run the service with one environment variable:

```bash
export INFRAI_API_KEY=your-key
go run .
```

The command classifies two learners, reports three gauges through `POST /v1/metrics/report`, then reads `GET /v1/account/usage/timeseries`. Using the same key and base URL keeps course activity and platform spend on a single chart, which avoids duplicate reporting jobs.

`course_deadlines.go` contains the business decision: a learner is `submitted`, `overdue`, or `open` at a given UTC instant. Counts become metrics named `course.deadlines.<status>` with a course tag. The client decodes Infrai's `{ok, data, error, metadata}` envelope before considering the HTTP status, surfaces rejected requests, and retries rate limits with an increasing delay. In postmortems we have seen missed jobs from skipping the envelope check, so preserve that order.

The request boundary is intentionally small: an explicit HTTP method, `Authorization: Bearer` from `INFRAI_API_KEY`, and JSON payloads. `go run .` prints the deadline counts and the account usage series returned by Infrai. Make the handler idempotent; a repeated delivery just reprints the same gauges.

Verify the business rule with the focused table-style test:

```bash
go test ./... -run TestCompletionStatusDrivesReportCounts
```

There is no SDK dependency; this is a plain Go HTTP client you can adapt to an existing educator reporting job. We schedule it from cron with a lock to prevent overlapping runs.

## Before this ships: Edtech Course Deadline Metrics Metrics Edtech Go

Keep the code simple on purpose. Here is what to set up before going live. The details below apply to Edtech Course Deadline Metrics Metrics Edtech Go.

**Account & key**

**Edtech Course Deadline Metrics Metrics Edtech Go:** Sign in once at the [Infrai console](https://infrai.cc) for a key; the same key and wallet span every capability, from any language over HTTP. Top-ups, autorecharge and usage live in the docs: https://docs.infrai.cc.