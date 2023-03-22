# pkgaffinity

[![Test](https://github.com/Syuparn/pkgaffinity/actions/workflows/test.yml/badge.svg)](https://github.com/Syuparn/pkgaffinity/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/syuparn/pkgaffinity)](https://goreportcard.com/report/github.com/syuparn/pkgaffinity)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

linter for Go source code that checks whether imports break encapsulation

# Usage
## Define Rules

`pkgaffinity` checks whether your code follows package import rules.

You can define rules in config files (see `.pkgaffinity.yaml` for details).

### Anti-Affinity Rules

Rules in `pkgaffinity` are described as package path anti-affinity in namespace.
Any package in anti-affinity rule are not allowed to import any other package in the rule.

#### Anti-Affinity List

Anti-affinity list forbids a package under a path prefix import packages under any other path prefixes.

```
# package tree
- github.com/foo/bar
  - baz
    - quux
  - pkg
    - hoge
      - fuga
```

```yaml
version: v1alpha1
antiAffinityRules:
  lists:
    - label: separate_hoge_and_baz
      pathPrefixes:
        - github.com/foo/bar/baz
        - github.com/foo/bar/pkg/hoge
```

Each path prefix is treated as import boundary.
Any import across the boundaries is forbidden (ex: import `github.com/foo/bar/baz/quux` from `github.com/foo/bar/pkg/hoge/fuga`).

#### Anti-Affinity Group

Anti-affinity group defines anti-affinity rules to all packages under *group* path prefix.

```
# package tree
- github.com/syuparn/pkgaffinity
  - pkg
    - config
      - domain
      - usecase
    - importchecker
      - domain
      - usecase
```

```yaml
version: v1alpha1
antiAffinityRules:
  groups:
    - pathPrefix: github.com/syuparn/pkgaffinity/pkg
```

Group `github.com/syuparn/pkgaffinity/pkg` treats subpackages directly under group (`config` and `importchecker`) as import boundaries.
Any import across the boundaries is forbidden (ex: import `github.com/syuparn/pkgaffinity/pkg/config/usecase` from `github.com/syuparn/pkgaffinity/pkg/importchecker/domain`).
This is suitable for microservices or modular monolith.

## Run command

This linter works as a single binary.

```bash
# install
$ go install github.com/syuparn/pkgaffinity/cmd/pkgaffinity@latest
```

If any violation found, it shows the reason.

```bash
# NOTE: this is only an example. Actuallly pkgaffinity follows all rules.
$ ./cmd/pkgaffinity/pkgaffinity ./...
package github.com/syuparn/pkgaffinity/pkg/config/domain: import "github.com/syuparn/pkgaffinity/pkg/importchecker/domain" breaks anti-affinity group rule `github.com/syuparn/pkgaffinity/pkg`
package github.com/syuparn/pkgaffinity/pkg/config/domain: import "github.com/syuparn/pkgaffinity/pkg/importchecker/domain" breaks anti-affinity group rule `github.com/syuparn/pkgaffinity/pkg`
pkgaffinity: check failed: violations found
pkgaffinity: check failed: violations found

$ echo $?
1
```

# Contributions

Any contributions are welcome!
