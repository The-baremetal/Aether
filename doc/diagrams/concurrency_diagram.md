## Concurrency (Implicit) Diagram

This is a concurrency diagram

```mermaid
sequenceDiagram
    participant Main Thread
    participant Worker pool thread 1 (shared between coroutines)
    participant Worker pool thread 2 (shared between coroutines)

    Main Thread->>Main Thread: Encounter `async` keyword
    Main Thread->>Worker pool Dispatch task to coroutine 1
    Main Thread->>Worker pool Dispatch task to coroutine 2
    Main Thread->>Main Thread: Continue execution
    Worker Thread 1->>Main Thread: Task completed
    Worker Thread 2->>Main Thread: Task completed
    Main Thread->>Main Thread: Gather results (implicit join)
```

**Description:**

The feature dispatches threads automatically to handle coroutines.