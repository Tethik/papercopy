# Go template

My own minimalist template for setting up a new Go project.

## Setup 🚀

- [ ] Besides the standard go tooling, ensure you have [goreleaser](https://goreleaser.com/) and `make` installed
- [ ] Copy the template (e.g. via github)
- [ ] Modify default code. Find and replace `github.com/Tethik/go-template`. Rename `example`.
- [ ] Modify `main: ./cmd/example/example.go` and `binary: app` in `.goreleaser.yaml`

### Github release workflow

Ensure you set `Actions -> Workflow permissions` to `Read and write permissions`

## Usage 🧑‍💻

- `make` or `make single-build` - build for just your arch. Outputs in `dist/`.
- `make build` - to build for all archs
- `make test` - to run tests

### Releases

To create a new release:

```sh
git tag -a vX.Y.Z # set your semantic version here
git push origin vX.Y.Z
```

Alternatively you can a manual release via make (not tested tbh)

`make release`

### References 📜

- [Repo by @nobe4 which this steals a bunch from](https://github.com/nobe4/safe)
- [Golang Standards Project Layout](https://github.com/golang-standards/project-layout)
