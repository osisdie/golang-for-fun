# golan-for-fun

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.23.4-blue.svg)](https://golang.org)

An simple testing project that tested in vscode.

---

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: Version 1.23 or higher.
- **Docker** (optional): For containerized builds and deployments.

---

## Folder Structure

```
   ├── app.go
    └── go.mod
    └── printrange/
        ├── go.mod
        ├── printrange.go
        └── printrange_test.go
```

---

## Installation

### **Install Go on Ubuntu for example**

1. **Update and Upgrade Packages**:

   ```bash
   sudo apt update && sudo apt upgrade -y
   ```

2. **Install Go**:
   - Download the latest version of Go:
     ```bash
     wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
     ```
   - Extract the downloaded file:
     ```bash
     sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
     ```

3. **Set Up Environment Variables**:
   - Open your shell configuration file (e.g., `.bashrc`, `.zshrc`, or `.profile`):
     ```bash
     nano ~/.bashrc
     ```
   - Add the following lines to the file:
     ```bash
     export PATH=$PATH:/usr/local/go/bin
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOPATH/bin
     ```
   - Save the file and reload the configuration:
     ```bash
     source ~/.bashrc
     ```

4. **Verify the Installation**:
   Check the Go version to ensure it’s installed correctly:
   ```bash
   go version
   # go version go1.23.4 linux/amd64
   ```

---

### **Install Additional Tools**

1. **Install `goimports`**:
   `goimports` is a tool for formatting Go code and managing imports:
   ```bash
   go install golang.org/x/tools/cmd/goimports@latest
   ```

2. **Install `gopls`**:
   `gopls` is the official Go language server for IDE integration:
   ```bash
   go install golang.org/x/tools/gopls@latest
   ```

3. **Install `go-outline`**:
   `go-outline` is an utility to extract JSON representation of declarations from a Go source file:
  ```bash
  go install github.com/ramya-rao-a/go-outline@latest
  ```

---

### **Set Up a Go Workspace**

1. **Navigate to your workspace directory**:
     ```bash
     cd src/
     ```

2. **Create a New Go Module**:
   - Initialize a new Go module:
     ```bash
     go mod init example.com/app
     ```

3. **Build and Run the Module**:
   - Save the file and run it:
     ```bash
     go run app.go
     ```

4. **This should also work**:
    ```bash
    go run example.com/app
    ```

   You will see the output:
   ```
   Hello, Go go1.23.4
   # ...
   ```

---

### **Testing**

1. **Test the Module**:
   - Test specific module path:
     ```bash
     go test example.com/app/printrange -count=1
     ```

   - Test specific all files:
     ```bash
     go test ./... -count=1
     ```

   You will see the output:
   ```
   ?       example.com/app [no test files]
   ok      example.com/app/printrange      0.005s
   ```
