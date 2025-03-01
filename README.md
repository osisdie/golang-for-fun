# golang-for-fun

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.23.4-blue.svg)](https://golang.org)

An simple testing project that tested in vscode.

<br>

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: Version 1.23 or higher.
- **Docker** (optional): For containerized builds and deployments.

<br>

## Folder Structure

```
   ├── app.go
    └── go.mod
    └── printrange/
        ├── go.mod
        ├── printrange.go
        └── printrange_test.go
```

<br>

## Installation

### **Install Go on Ubuntu for example**

1. **Update and Upgrade Packages**:

   ```sh
   sudo apt update && sudo apt upgrade -y
   ```

2. **Install Go**:
   - Download the latest version of Go:
     ```sh
     wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
     ```
   - Extract the downloaded file:
     ```sh
     sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
     ```

3. **Set Up Environment Variables**:
   - Open your shell configuration file (e.g., `.bashrc`, `.zshrc`, or `.profile`):
     ```sh
     nano ~/.bashrc
     ```
   - Add the following lines to the file:
     ```sh
     export PATH=$PATH:/usr/local/go/bin
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOPATH/bin
     ```
   - Save the file and reload the configuration:
     ```sh
     source ~/.bashrc
     ```

4. **Verify the Installation**:
   Check the Go version to ensure it’s installed correctly:
   ```sh
   go version
   # go version go1.23.4 linux/amd64
   ```

<br>

### **Install Additional Tools**

1. **Install `goimports`**:
   `goimports` is a tool for formatting Go code and managing imports:
   ```sh
   go install golang.org/x/tools/cmd/goimports@latest
   ```

2. **Install `gopls`**:
   `gopls` is the official Go language server for IDE integration:
   ```sh
   go install golang.org/x/tools/gopls@latest
   ```

3. **Install `go-outline`**:
   `go-outline` is an utility to extract JSON representation of declarations from a Go source file:
  ```sh
  go install github.com/ramya-rao-a/go-outline@latest
  ```

---

## **Run**

1. **Navigate to your workspace directory**:
     ```sh
     cd src/
     ```

2. **Create a New Go Module**:
   - Initialize a new Go module:
     ```sh
     go mod init example.com/app
     ```

3. **Build and Run the Module**:
   - Save the file and run it:
     ```sh
     go run app.go
     ```

4. **This should also work**:
    ```sh
    go run example.com/app
    ```

   You will see the output:
   ```
   Hello, Go go1.23.4
   # ...
   ```

<br>

## **Testing**

1. **Test the Module**:
   - Test specific module path:
     ```sh
     go test example.com/app/printrange
     ```

   - Test specific all files:
     ```sh
     go test ./... -count=1 # the cache is bypassed
     go test ./... -count=1 -v # -v with verbose
     ```

   You will see the output:
   ```
   ?       example.com/app [no test files]
   === RUN   TestPrintRangeNumbers
   === RUN   TestPrintRangeNumbers/_
   === RUN   TestPrintRangeNumbers/,_
   === RUN   TestPrintRangeNumbers/_|_
   --- PASS: TestPrintRangeNumbers (0.00s)
      --- PASS: TestPrintRangeNumbers/_ (0.00s)
      --- PASS: TestPrintRangeNumbers/,_ (0.00s)
      --- PASS: TestPrintRangeNumbers/_|_ (0.00s)
   PASS
   ok      example.com/app/printrange      0.004s
   ```

<br>

## **Run the Program**

1. **Compiles the Go application into an executable binary**:
   - Compiles the Go application into an executable binary named `app`
   ```sh
   go build -o app .
   ```

2. **Console run the binary app**:
   ```sh
   ./app
   ```

   You will see the same output like above.

<br>
<br>

---

**Enjoy this golang_for_fun project!**

---
