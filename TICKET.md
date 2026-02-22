# PLATFORM-2887: Build windowed data aggregation worker

**Status:** In Progress · **Priority:** Medium
**Sprint:** Sprint 25 · **Story Points:** 5
**Reporter:** Raj Kapoor (Data Engineering Lead) · **Assignee:** You (Intern)
**Due:** End of sprint (Friday)
**Labels:** `backend`, `golang`, `streaming`, `feature`
**Task Type:** Feature Ship

---

## Description

The real-time aggregation system needs a window manager that groups incoming data into time windows and computes aggregates (sum, avg, count, min, max). A working `Aggregator` exists. Implement the TODO items in `WindowManager`.

## Acceptance Criteria

- [ ] `AddEvent()` assigns events to correct time window
- [ ] `CloseWindow()` computes aggregates for a completed window
- [ ] `GetActiveWindows()` returns currently open windows
- [ ] Sliding and tumbling window types supported
- [ ] All unit tests pass
