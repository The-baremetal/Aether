## Assertion Diagram

This diagram is for the assertion feature

```mermaid
graph LR
    A[Source Code] --> B(Assertion Statement: assert(condition, message));
    B --> C{Condition Evaluation};
    C -- True --> D[Continue Execution];
    C -- False --> E[Raise Assertion Error: message];
    E --> F[Program Termination or Exception Handling];
```