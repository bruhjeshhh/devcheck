# Contributing

All contributions are welcome — new checks, bug fixes, docs, whatever.

## Setup

Fork the repo first, then clone your fork:

```bash
git clone https://github.com/YOUR_USERNAME/devcheck
cd devcheck
go build ./cmd/devcheck
go test ./...
```

## Adding a new check

Most contributions are new checks. Here's how:

1. Create `internal/checks/yourcheck.go`
2. Implement the `Check` interface:

```go
type YourCheck struct{}

func (c *YourCheck) Name() string { return "your check name" }

func (c *YourCheck) Run(ctx context.Context) check.Result {
    return check.Result{
        Name:    c.Name(),
        Status:  check.StatusPass,
        Message: "looks good",
        Fix:     "", // shown with --fix if this fails
    }
}
```

3. Register it in `internal/check/registry.go` under the right stack condition
4. Add a test in `internal/checks/yourcheck_test.go` — cover both pass and fail
5. Run `go test ./...` and make sure everything passes
6. Open a PR referencing the issue (e.g. "Closes #5")

Before you start on something, leave a comment on the issue so nobody duplicates work.

## PR checklist

- `go build ./...` passes
- `go test ./...` passes
- `go vet ./...` passes
- Test covers pass and fail cases

## Questions

Open a GitHub Discussion or leave a comment on the issue.
