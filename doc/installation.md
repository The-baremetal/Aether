# Install guide

### Prerequisites
Go 1.20+ for compilation of the project
All packages mentioned in go.sum
A machine that has Bash installed with make and LLVM
Atleast 50 mb of ram and 500mb of storage

# Estimates

### Estimate of storage
- LLVM: 1-3 gb of storage
- Go: 1-2 gb of storage
- Aether: 5-10 mb of storage

### Estimate of ram
- LLVM: 500-700mb
- Go: 600-900mb
- Aether: 10-20mb

# Installing Aether

## Step 1
### Clone the repository
```git clone https://github.com/The-baremetal/Aether.git && cd Aether```
## Step 2
### Gettings the dependencies
Go: Install go at: https://golang.org/doc/install or run ```sudo apt install golang```
LLVM: You can get llvm by running the command ```sudo apt install llvm```
## Step 3
### Building
```make build```
