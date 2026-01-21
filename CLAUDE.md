# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

DocVault is a self-hosted cloud document management system built with Go (backend) and React (frontend planned). The backend is implemented; the frontend is not yet started.

## Development Commands

### Backend (Go)
```bash
cd backend

# Download dependencies
go mod download

# Run locally (default: port 8080, storage: ./data/md)
go run cmd/server/main.go

# Run with custom document root
DOC_ROOT=/custom/path go run cmd/server/main.go

# Build executable
go build -o docvault-server cmd/server/main.go
```

### Environment Variables
- `DOC_ROOT`: Document storage directory (default: `./data/md`)
- Port is hardcoded to `8080` in `cmd/server/main.go`

## Architecture

The backend follows a clean three-layer architecture:

```
Handler Layer (HTTP) → Service Layer (Business Logic) → FS Layer (File System)
```

- **`internal/handler/`**: HTTP request/response handling using Gin framework
- **`internal/service/`**: Business logic for document operations
- **`internal/fs/`**: File system abstraction with path validation

### API Endpoints
Base URL: `http://localhost:8080/api`

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/docs` | List all documents |
| GET    | `/docs/:name` | Get document content |
| PUT    | `/docs/:name` | Update document |
| POST   | `/docs` | Create new document |
| DELETE | `/docs/:name` | Delete document |

## Key Implementation Details

- **Single-user system**: No database; all documents stored as markdown files
- **File extension**: All documents automatically get `.md` extension
- **Path security**: Path traversal protection blocks `..` in filenames
- **Entry point**: `backend/cmd/server/main.go` (30 lines, very minimal)

## Current Status

- **Backend**: Complete, ~209 lines across 4 Go files
- **Frontend**: Not implemented (empty `frontend/` directory)
- **Tests**: None exist yet
- **Docker**: Mentioned in README but no Dockerfile present
