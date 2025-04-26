# Aether roadmap

## First Priority: Reintegrate Lexer and Parser

1.  Examine the new lexer and parser implementations.
2.  Integrate the new lexer and parser into the existing codebase.
3.  Implement implicit borrowing and ownership (leash system) in `src/aether/codegen/codegen.go`.
4.  Implement atomic operations for mutable borrowing in `src/aether/codegen/codegen.go`.
5.  Update documentation (`doc/spec.md` and `doc/tutorial.md`) to reflect the changes.

## 1.1
> Operations (+, -, *, /)
> Custom Parser
> Custom Lexer

## 1.1.5
> Package system via import kw, folder based with the module file that has a name (path: (MODULE NAME)/(MODULE FOLDER)) Importing a parent folder also imports its child folders, if certain packages and functions arent used, they are not included in the final bundle.
> Leash system

## 1.2
> Concurrency (implicit)
> Assertion
> Efficient compiler
> Safety rules (borrowing and ownership (implicit))

## 1.3
> Native package manager integration with the mirror url "packages.aether.org" (official mirror compatible with APT, DNF)
> Imports can use the Aether root and the project root, Aether root contains the packages installed from packages.aether.org

## 1.4
> Hot reload
> Hot compile avoids compiling a chunk of code if it has already been compiled

## 1.5
> Aether rewrite in Aether
> More insentive to add features
> Maintainable codebase

## Compiler-Controlled Multithreading with Atomic Operations

1.  **Compiler-Controlled Multithreading:** The compiler will automatically handle multithreading by breaking down tasks and executing them in parallel as needed.
2.  **Atomic Operations for Memory Safety and Synchronization:** Aether will use atomic operations (e.g., CAS, fetch-and-add) to synchronize access to shared resources between threads, avoiding locks, mutexes, and message-passing.
3.  **Memory Safety via Immutability and Ownership:** Memory safety will be ensured through immutability and atomicity, with data treated as immutable by default when shared across tasks. There will be no explicit borrowing or ownership control.
4.  **Optimizing for Performance:** Atomic operations will minimize synchronization overhead, and the compiler will manage tasks and ensure safe concurrent execution for efficient multithreading and optimal memory usage.
5.  **Simplicity for the Programmer:** The programmer doesn’t have to worry about multithreading details, atomicity, or memory safety, allowing them to focus on writing code that expresses their logic without dealing with concurrency issues.
6.  **Future Considerations:** Fine-grained control could be added through compiler annotations, and more advanced strategies (such as memory pooling, advanced atomic operations, or load balancing) can be added to enhance performance. Metatables, functions and stuff.