## Adaptive Floating-Point Numbers Diagram


```mermaid
graph LR
    A[Floating-Point Expression] --> B{Potential Precision Loss?};
    B -- Yes --> C[Arbitrary Precision Calculation];
    B -- No --> D[IEEE-754 Calculation];
    C --> E[Result (Arbitrary Precision)];
    E --> F(Optimize Representation);
    F --> G[Result (Optimized Arbitrary Precision)];
    D --> H[Result (IEEE-754)];
    G --> I(Output);
    H --> I;
```

**Description:**

This system automatically decides between IEEE-754 or arbitrary precision adjusting to different scenarios. For example comparison...
```lua
if 0.1+0.2 == 0.3 then
    print('Arbitrary precision was used')
end
```
```output
Arbitrary precision was used
```
In this scenario, the system was automatically switched to arbitrary precisions as using IEEE-754 would result in something breaking
In this new scenario, the system compares it to a bigger number
```lua
if 0.1+0.2 < 0.4 then
    print('The system would use the IEEE-754 as using the arbitrary system would not impact the result.')
end
```
```output
The system would use the IEEE-754 as using the arbitrary system would not impact the result.
```

Plan: 1.7 will introduce this system
The compiler would precompute the variables, by default, llvm will precompute the floatingpoint arithmetic using IEEE-754, instead, the compiler not the backend will precompute the variables.
This system only applies to static variables that will not be changed after, if it does, it will default to IEEE-754, the programmer will be able to explicitly use arbitrary precision.
This also only applies to const variables or variables that wont be changed later.