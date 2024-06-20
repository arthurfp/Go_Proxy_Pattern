# Proxy Pattern in Go

## Overview
This repository demonstrates the application of the Proxy design pattern in Go. The project highlights how to provide a surrogate or placeholder for another object to control access to it, showcasing flexibility and best practices in design patterns and unit testing.

## Pattern Description
The Proxy pattern is used to provide a surrogate or placeholder for another object to control access to it. This pattern is useful in various scenarios such as lazy initialization, access control, logging, and more. In this project, we have implemented the Proxy pattern using a `RealSubject` and a `Proxy` to manage access to it.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the creation and usage of the proxy.
- **pkg/proxy/**: Houses the Proxy implementation.
  - **subject.go**: Defines the `Subject` interface.
  - **real_subject.go**: Implements the `RealSubject`.
  - **proxy.go**: Implements the `Proxy` to control access to the `RealSubject`.
  - **real_subject_test.go**: Unit tests for `RealSubject`.
  - **proxy_test.go**: Unit tests for `Proxy`.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Proxy_Pattern.git
cd Go_Proxy_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp