graph LR
    A[File Change] --> B(File Watcher);
    B --> C{Check if Hot Reloading Enabled};
    C -- Yes --> D[Invalidate Cache];
    D --> E[Recompile Changed Modules];
    E --> F[Update Running Application];
    F --> H[Modules]
    F --> J[Module Replacement]
    J --> K[Unload Old Module]
    K --> L[Load New Module]
    L --> H;
    E --> G[Dependency Graph];
    G --> E;
    C -- No --> I[Full Application Restart];
    I --> F;
    style A fill:#f9f,stroke:#333,stroke-width:2px
    style B fill:#ccf,stroke:#333,stroke-width:2px
    style C fill:#ccf,stroke:#333,stroke-width:2px
    style D fill:#f9f,stroke:#333,stroke-width:2px
    style E fill:#f9f,stroke:#333,stroke-width:2px
    style F fill:#ccf,stroke:#333,stroke-width:2px
    style G fill:#f9f,stroke:#333,stroke-width:2px
    style H fill:#f9f,stroke:#333,stroke-width:2px
    style I fill:#ccf,stroke:#333,stroke-width:2px
    style J fill:#f9f,stroke:#333,stroke-width:2px
    style K fill:#f9f,stroke:#333,stroke-width:2px
    style L fill:#f9f,stroke:#333,stroke-width:2px
