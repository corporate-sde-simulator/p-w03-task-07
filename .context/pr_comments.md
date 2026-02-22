# PR Review - Real-time data aggregation worker (by Suresh)

## Reviewer: Priya Menon
---

**Overall:** Good foundation but critical bugs need fixing before merge.

### `aggregator.go`

> **Bug #1:** Tumbling window boundary calculation uses wrong modulo causing windows to overlap instead of being adjacent
> This is the higher priority fix. Check the logic carefully and compare against the design doc.

### `windowManager.go`

> **Bug #2:** Count aggregation counts expired events that should have been evicted from the window
> This is more subtle but will cause issues in production. Make sure to add a test case for this.

---

**Suresh**
> Acknowledged. I have documented the issues for whoever picks this up.
