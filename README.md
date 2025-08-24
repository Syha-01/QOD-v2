# Go Web Server: Quote of the Day API

The goal of this project is to create a "Quote of the Day" (QOD) API, implementing modern best practices for configuration, logging, routing, and deployment.

**Go Version:** `1.25`

## Learning Objectives

This project serves as a practical example of how to:

-   Build a web server using Go's `net/http` package.
-   Serve JSON responses.
-   Configure a server using command-line flags.
-   Implement structured logging.
-   Handle graceful shutdown.
-   Implement Cross-Origin Resource Sharing (CORS).
-   Structure a Go application for containerization.

## Prerequisites

Before you begin, ensure you have the following installed on your local machine:
-   **Go** (version 1.25 or compatible)
-   **Git**

## Getting Started

Follow these steps to get the application running on your local machine.

### 1. Clone the Repository

First, clone this repository to your local machine:
```bash
git clone https://github.com/Syha-01/QOD-v2.git
cd qod
```

### 2. Install Dependencies

The project uses Go Modules to manage dependencies. Run the following command to download and install the required packages:
```bash
go mod tidy
```
This will install dependencies like `julienschmidt/httprouter` as defined in the `go.mod` file.

## Usage

### Running the Application

A `Makefile` is included to simplify common tasks. To start the web server, run:
```bash
make run/api
```
This will compile and run the application. By default, the server will start on `localhost:4000`.


### Testing the API

Once the server is running, you can test its endpoints.

#### Healthcheck Endpoint

Open a new terminal and use `curl` to send a request to the healthcheck endpoint:
```bash
curl -i localhost:4000/v1/healthcheck
```
You should see a `200 OK` status and a response body similar to this:
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 24 Aug 2025 00:13:24 GMT
Content-Length: 73

{"status": "available", "environment": "development", "version": "1.0.0"}
```

### Configuration

The application's configuration can be controlled via command-line flags.

-   `-port`: The port for the API server to listen on (default: `4000`).
-   `-env`: The operating environment (default: `development`).
