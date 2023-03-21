# pkgaffinity
linter for Go source code that checks whether imports break encapsulation

# Usage
## Define Rules

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

