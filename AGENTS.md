# Repository Guidelines

## Project Structure & Module Organization
- `main.go` contains the CLI entrypoint and network flow for connecting to NicoLive.
- `proto/` holds the protobuf source definitions (`.proto`).
- `gen/pb/` contains generated Go protobuf types. Treat as generated output.
- `go.mod`/`go.sum` manage Go module dependencies.

## Build, Test, and Development Commands
- `go run main.go <lvid>` runs the client against a live ID (example: `lv123456`).
- `go build ./...` builds all packages in the module.
- `go test ./...` runs tests (none currently; expect “no test files”).

## Coding Style & Naming Conventions
- Go standard formatting: use `gofmt` (tabs for indentation).
- Follow Go naming conventions (CamelCase for exported identifiers, lowerCamel for unexported).
- Keep generated code in `gen/pb/`; do not hand-edit files there.
- Prefer small, focused functions for protocol handling and decoding logic.

## Testing Guidelines
- Add tests alongside code using `*_test.go` filenames.
- Use Go’s `testing` package; table-driven tests are preferred for protocol parsing.
- When adding fixtures, keep them in a new `testdata/` directory near the tests.

## Commit & Pull Request Guidelines
- Commits in history use short, imperative summaries (e.g., “Create .gitignore for Go project”).
- Keep commit subjects concise and descriptive; body optional when change is obvious.
- PRs should include: purpose, key changes, and how to run (`go run ...`, `go test ./...`).
- Link related issues or NicoLive protocol references when applicable.

## Configuration & Runtime Notes
- The client fetches live metadata and connects to NicoLive WebSocket endpoints.
- Avoid hard-coding stream IDs in code; pass them via CLI.
- If adding environment-based configuration, document expected vars in this file.
