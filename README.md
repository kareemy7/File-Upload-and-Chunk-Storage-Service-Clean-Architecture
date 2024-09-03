# File Upload and Chunk Storage Service

## Overview

This project implements a File Upload and Chunk Storage Service using Clean Architecture principles. It allows users to upload large files in chunks and download them later. The service uses LevelDB for efficient storage and retrieval of file chunks and metadata.

## Features

- Upload large files in chunks
- Download previously uploaded files
- Efficient storage using LevelDB
- Clean Architecture design for maintainability and scalability

## Technology Stack

- Go (Golang)
- Gin Web Framework
- LevelDB
- Clean Architecture

## Project Structure

The project follows Clean Architecture principles:

- `entities`: Core business objects
- `usecases`: Application-specific business rules
- `repositories`: Data access layer interfaces and implementations
- `controllers`: HTTP request handlers
- `infrastructure`: Database and server setup

## Getting Started

### Prerequisites

- Go 1.16 or later
- LevelDB

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/kareemy7/File-Upload-and-Chunk-Storage-Service-Clean-Architecture.git
   ```

2. Navigate to the project directory:
   ```
   cd File-Upload-and-Chunk-Storage-Service-Clean-Architecture
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Service

1. Start the server:
   ```
   go run main.go
   ```

2. The server will start on `localhost:8080` by default.

## API Endpoints

- `POST /upload`: Upload a file
- `GET /download/:file_id`: Download a file by its ID

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for providing excellent libraries and tools.
- This project structure is inspired by Clean Architecture principles as described by Robert C. Martin.