## Package System Diagram

This diagram describes the package system of Aether

```mermaid
graph LR
    A[Source Code] --> B(Import Statement: import "module/package");
    B --> C{Module Resolution};
    C --> C1{Check if module is in STD};
    C1 -- Yes --> H(Import module and its sub-packages);
    C1 -- No --> C2["Find module file (module.flux) in module/package"];
    C2 -- Found --> E(Parse module.flux for module name);
    C2 -- Not Found --> F[Error: Module not found];
    E --> G["Check if module name matches expected path"];
    G -- Match --> H;
    G -- No Match --> I[Error: Module name mismatch];
    H --> J{Tree shaking};
    J --> K["Include only used packages and functions in final bundle"];
    K --> L[Final Bundle];
```

**Description:**

This package system imports packages when the ``import`` keyword is used.

**USAGE:**

```aether 
import math```

```aether
import user32.dll -- this will be resolved using build flags ex. -i math.dll```