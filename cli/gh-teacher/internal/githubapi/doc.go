// Package githubapi holds gh-teacher's GitHub-specific domain operations that
// build on the shared transport seam (cli/shared/githubapi): org plan/billing
// reads and the optimistic tree-commit-with-rebase loop. The transport verbs,
// pagination, auth, and forge-neutral helpers live in the shared package; this
// package is what stays teacher-specific until the Forge interface (F2+) can
// model these operations.
package githubapi
