# Papercopy

Tool to generate a papercopy of cryptographic keys in a PDF format.

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


### Ideas

* Encrypt before using key or password derived key

* Shell autocomplete on mnemonic characters

* Logo in QR Code for style

* Form to gather input to customize the pdf

### References 📜

- [Repo by @nobe4 which this steals a bunch from](https://github.com/nobe4/safe)
- [Golang Standards Project Layout](https://github.com/golang-standards/project-layout)
