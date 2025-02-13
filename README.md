# Getting Started with Go

## Install Go
Install Go via [Homebrew](https://brew.sh) or from the [official website](https://golang.org/dl/).

## Install GoLand IDE
Download and install [GoLand](https://www.jetbrains.com/go/).

## Setup Your Project

1. Open your terminal and navigate to your desired directory:
    ```sh
    cd
    ```

2. Create a new folder for your project:
    ```sh
    mkdir New_Folder
    cd New_Folder
    ```

3. Initialize a new Go module:
    ```sh
    go mod init New_Folder/hello
    ```
    This will create a `go.mod` file with the module name `New_Folder/hello`.

4. Create a new Go file named `hello.go` and add the following code:
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
    ```

5. Run your code:
    ```sh
    go run .
    ```

## Using External Packages

1. To use an external package, import it in your Go file:
    ```go
    import "rsc.io/quote"
    ```

2. Run the following command to add the package to your module:
    ```sh
    go mod tidy
    ```
    This will find and add the required module for `rsc.io/quote`.

3. Run your code again:
    ```sh
    go run .
    ```