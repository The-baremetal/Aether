# Aether roadmap

## Milestone 1 (reached)
> Operations (+, -, *, /)
> Custom Parser
> Custom Lexer

## Milestone 2(incomplete)
> Package system via import kw, folder based with the module file that has a name (path: (MODULE NAME)/(MODULE FOLDER)) Importing a parent folder also imports its child folders, if certain packages and functions arent used, they are not included in the final bundle.

## Milestone 3
> Aether rewrite in Aether
> Rewrite using so libraries and slowly add features and move away from so libraries written in go.

## Milestone 4
> Native package manager integration with the mirror url "packages.aether.org" (official mirror compatible with APT, DNF)
> Imports can use the Aether root and the project root, Aether root contains the packages installed from packages.aether.org

## Milestone 5
> Hot reload
> Hot compile avoids compiling a chunk of code if it has already been compiled

## Milestone 6
> Concurrency (implicit)
> Assertion
> Efficient compiler
> Safety rules (borrowing and ownership (implicit))
