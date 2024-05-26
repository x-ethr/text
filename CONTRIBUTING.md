# Contribution Guide

## Package Publication

_The following section refers to publishing package(s) to https://pkg.go.dev._

- See GO's [*Publishing a Module*](https://go.dev/doc/modules/publishing) for additional details.

1. Establish a [`LICENSE`](https://spdx.org/licenses/) to the project.
2. Ensure dependencies are updated.
    ```bash
    go mod tidy
    ```
3. Verify [`vendor`](https://go.dev/ref/mod#vendoring) requirements.
    ```bash
    go mod vendor
    ```
4. Sync the working tree's `HEAD` with its remote.
    ```bash
    git add .
    git commit --message "<commit-msg>"
    git push --set-upstream origin main
    ```
5. Assign a tag and push.
    ```bash
    git tag "v$(head VERSION)" && git push origin "v$(head VERSION)"
    ```
6. Make the module available, publicly.
    ```bash
    GOPROXY=proxy.golang.org go list -m "github.com/x-ethr/example@v$(head VERSION)"
    ```

### Pre-Commit

The following project makes use of `pre-commit` for local-development `git-hooks`. These hooks are useful
in cases such as preventing secrets from getting pushed into version-control.

See the [`.pre-commit-config.yaml`](.pre-commit-config.yaml) for implementation specifics.

#### Local Setup

1. Install pre-commit from https://pre-commit.com/#install.
2. Auto-update the config to the latest repos' versions by executing `pre-commit autoupdate`.
3. Install with `pre-commit install`.

#### General Command Reference(s)

**Update the configuration's upstreams**

```bash
pre-commit autoupdate
```

**Install `pre-commit` to local instance**

```bash
pre-commit install
```

## Documentation

Tool `godoc` is required to render the documentation, which includes examples.

- See [`doc.go`](./doc.go) for code-specific package documentation.

Installation Steps:

1. Install `godoc`.
    ```bash
    go install golang.org/x/tools/cmd/godoc@latest
    ```
1. Backup shell profile and update `PATH`.
    ```bash
    cp ~/.zshrc ~/.zshrc.bak
    printf "export PATH=\"\${PATH}:%s\"\n" "$(go env --json | jq -r ".GOPATH")/bin" >> ~/.zshrc
    source ~/.zshrc
    ```
1. Start the `godoc` server.
    ```bash
    godoc -http=:6060
    ```
1. Open the webpage.
    ```bash
    open "http://localhost:6060/pkg/"
    ```
