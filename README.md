# Contextive Demo - Go - Service

This repository illustrates the use of Contextive in an environment where multiple repositories are part of the same `bounded context` and need access to the same set of definitions.

In this pattern, the definitions are contained within a `common` repository - e.g. https://github.com/dev-cycles/contextive-demo-go-common

Native package management and distribution techniques are leveraged to ensure the definitions are available to your IDE while working on the service repo.

## How does it work?

1. Add your definitions file to a 'shared' repository for the bounded context - something which stores any common context-wide code, such as entities, contexts, constants etc.
   1. In this demo, it's [contextive-demo-go-common](https://github.com/dev-cycles/contextive-demo-go-common)
2. Ensure the definitions file is included in the distributed package.
   1. For `go` this is as simple as having it in the repository
3. Depend on the common package using standard package techniques.
   1. In this case `go get`.
   2. See [go.mod](go.mod) for the reference to [contextive-demo-go-common](https://github.com/dev-cycles/contextive-demo-go-common)
4. Set the `contextive.path` setting in the service repository as a workspace setting and leverage a shell command to locate the `contextive-demo-go-common` package:
   1. `$(go list -m -f '{{.Dir}}' github.com/dev-cycles/contextive-demo-go-common)/.contextive/definitions.yml`
   2. See [./.vscode/settings.json](./.vscode/settings.json) for an example

When you clone this current repository, after running `go build` or `go test`, the dependencies will be downloaded and the definitions defined in [contextive-demo-go-common](https://github.com/dev-cycles/contextive-demo-go-common) should be available in the UI.

## Notes

* When the definitions are updated in the `common` package, it will need republishing, and the services will need to update to the latest version of the package to pick up the updated definitions.

* Contextive currently only activates in VSCode on the presence of a `.contextive` folder in the repo. To that end, a dummy [`README.md`](./.contextive/README.md) is included in that folder to ensure Contextive activates.