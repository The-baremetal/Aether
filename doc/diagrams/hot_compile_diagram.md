graph LR
    A[Source Code Change] --> B(Compiler);
    B --> B1(Lexical Analysis);
    B1 --> B2(Parsing);
    B2 --> B3(Semantic Analysis);
    B3 --> C{Check for Errors};
    C -- No Errors --> D[Generate Intermediate Representation];
    D --> E[Optimize IR];
    E --> F[Generate Machine Code];
    F --> G[Update Running Application];
    G --> K[Dynamic Linking]
    K --> M[Load New Code]
    M --> G;
    G --> L[Code Patching]
    L --> N[Replace Old Code]
    N --> G;
    C -- Errors --> H[Report Errors];
    style A fill:#f9f,stroke:#333,stroke-width:2px
    style B fill:#ccf,stroke:#333,stroke-width:2px
    style C fill:#ccf,stroke:#333,stroke-width:2px
    style D fill:#f9f,stroke:#333,stroke-width:2px
    style E fill:#f9f,stroke:#333,stroke-width:2px
    style F fill:#f9f,stroke:#333,stroke-width:2px
    style G fill:#ccf,stroke:#333,stroke-width:2px
    style H fill:#f9f,stroke:#333,stroke-width:2px
    style B1 fill:#f9f,stroke:#333,stroke-width:2px
    style B2 fill:#f9f,stroke:#333,stroke-width:2px
    style B3 fill:#f9f,stroke:#333,stroke-width:2px
    style K fill:#f9f,stroke:#333,stroke-width:2px
    style L fill:#f9f,stroke:#333,stroke-width:2px
    style M fill:#f9f,stroke:#333,stroke-width:2px
    style N fill:#f9f,stroke:#333,stroke-width:2px
