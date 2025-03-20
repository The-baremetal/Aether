# FLUX ASSEMBLY

[![GitHub release](https://img.shields.io/github/release/platane/snk.svg?style=flat-square)](https://github.com/The-baremetal/FLUXASSEMBLY/releases/latest) [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=I%20Just%20Found%20This%20Named%20Aether&url=https://github.com/The-baremetal/FLUXASSEMBLY&hashtags=coding,programming,programminglanguage,aether,aetherlang,project,fyp)
[![GitHub license](https://img.shields.io/github/license/The-baremetal/FLUXASSEMBLY)](https://github.com/The-baremetal/FLUXASSEMBLY/blob/master/LICENSE)
[![GitHub contributors](https://img.shields.io/github/contributors/The-baremetal/FLUXASSEMBLY)](https://github.com/The-baremetal/FLUXASSEMBLY/graphs/contributors)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/m/The-baremetal/FLUXASSEMBLY)](https://github.com/The-baremetal/FLUXASSEMBLY/commits)


## Relativity
> Flux assembly not to be confused with flux lang is a language that converts directly to x86 assembly.
## Why Flux Assembly?
> Flux Assembly has been designed to be easy to learn. Its syntax is based on Lua, with all the characteristics which make it a good, 'first language' for immediate introduction to Assembly without having to get involved in all the complexities of low-level programming: self-contained and easily modifiable, personalized and flexible language.
> What's more, Flux Assembly is fully backward compatible with Lua. In other words, any pre-existent code in Lua can be easily ported into Flux Assembly, therefore easing the path to migrate from a higher-level, easier-to-handle scripting to an efficient, low-level assembly.
## Attribution
> The grammar and syntax of Flux Assembly are heavily inspired by Lua, a language renowned for its simplicity and ease of learning. Lua, developed by the Lua authors (https://www.lua.org/), has been a popular choice for developers looking for an accessible programming language.
> Flux Assembly is a derivative work of Lua's syntax, not intended to replace Lua but to offer an assembly-level option that retains the same beginner-friendly appeal. While this language shares the essence of Lua, it is not an official Lua implementation but a creative evolution designed to bridge the gap between high-level scripting and low-level assembly programming.
> By building on Lua's foundation, Flux Assembly ensures that developers can experience the power and flexibility of assembly without sacrificing the simplicity of their original coding language.
## Build instructions

### Prerequisites
Go 1.20+ for compilation of the project
All packages mentioned in go.sum
A machine that has Bash installed with SH support

> Step 1
### Clone the repository
```git clone https://github.com/The-baremetal/FLUXASSEMBLY.git```
```cd FLUXASSEMBLY```
> Step 2
### Gettings the dependencies
Go: Install go at: https://golang.org/doc/install or run ```sudo apt install golang```
LLVM: You can get llvm by running the command ```sudo apt install llvm```
> Step 3
### Actually building Flux Assembly
``` build.sh build```
Now that your Flux Assembly is installed, you can run ```./main```
Example files: ### input.lua
