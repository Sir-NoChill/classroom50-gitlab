// Package githubapi holds gh-student's GitHub-specific domain operations that
// build on the shared transport seam (cli/shared/githubapi): enabling a Pages
// site on a student repo. The transport verbs, pagination, auth, and
// forge-neutral helpers live in the shared package; this package is what stays
// student-specific until the Forge interface (F2+) can model these operations.
package githubapi
