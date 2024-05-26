# Contribution Guide

## Package Publication

_The following section refers to publishing package(s) to https://pkg.go.dev._

- See GO's [*Publishing a Module*](https://go.dev/doc/modules/publishing) for additional details.

1. Ensure a [`LICENSE`](https://spdx.org/licenses/) is added to the project.
2. Assign a tag to the package source and push it into VCS.
```bash

```


git commit --message "Release Version: v0.1.0"
git tag v0.1.0


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
