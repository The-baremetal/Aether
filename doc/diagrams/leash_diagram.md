Leash diagram
```mermaid
graph LR
    A[Variable] --> B{Ownership?};
    B -- Yes --> C[Owner];
    B -- No --> D[Borrowed];

    C --> E{Transfer Ownership?};
    E -- Yes --> F[New Owner];
    E -- No --> C;

    D --> G{Mutable?};
    G -- Yes --> H[Mutable Borrow (Exclusive)];
    G -- No --> I[Immutable Borrow (Shared)];

    H --> J{Lifetime End?};
    J -- Yes --> K[Invalidate Mutable Borrow];
    J -- No --> H;

    I --> L{Lifetime End?};
    L -- Yes --> M[Invalidate Immutable Borrow];
    L -- No --> I;

    C --> N{Scope End?};
    N -- Yes --> O[Ownership Ends, Memory Reclaimed];
    N -- No --> C;

    F --> N;
```

**Memory Management:**

*   By default, no explicit `free` is generated at the end of execution. The operating system takes the responsibiliy on terminating the process and all the memory that is allocated to it.
*   If the OS does not automatically free the memory, use the `--memfree_eox` flag to enable explicit freeing of `eox` (end of execution) variables.

**Description:**

Every variable has an owner, it can be borrowed or the owner can be changed.
```lua
function test_complex()
    root = { value = 100 }  -- Implicitly owned by function

    -- Nested ownership
    root.branch1 = { value = 50 }
    root.branch2 = root  -- Cloning (new reference, new owner)
    
    -- Immutable borrow (temporary read)
    borrowed_ref = root.branch1
    print(borrowed_ref.value) -- OK: Just reading (no possibility of a write)

    -- Mutable borrow (temporary write)
    borrowed_mut = root.branch1
    borrowed_mut.value = 75  -- OK: Still in scope

    -- Ownership transfer
    transferred = root -- `root` is now empty, `transferred` owns the value
    
    -- Nested borrowing
    leaf1 = transferred.branch1 -- Immutable borrow
    leaf2 = transferred.branch1 -- Mutable borrow
    
    -- Cloning a nested structure
    root.branch2.leaf3 = { value = 25 }

    print(leaf1.value)  -- OK: Immutable borrow
    leaf2.value = 90    -- OK: Mutable borrow
    
    return transferred -- Ownership transferred if assigned outside
end  -- Everything unreferenced is freed here
-- if there was more code here, it wouldn't free the variables that will be used.
-- no free code at the end as os would free it. If the os wouldn't free it, use the flag --memfree_eox (eox stands for END OF EXECUTION)
```


if a variable is const, do not generate free code for it, its not stored in memory but hardcoded in program binary
if a variable isnt manually declared a const but is never changed, turn it into a const