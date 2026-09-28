// Package githubapi is the shared GitHub REST transport seam for the
// Classroom 50 CLIs. It is the single place permitted to import
// github.com/cli/go-gh/v2/pkg/api (both gh-teacher and gh-student route
// through it via their own thin githubapi packages, which hold only
// forge-specific domain operations).
//
// The Client interface is transport-verb-level (Get/Post/Patch/Request/
// RequestWithContext), not per-operation — domain shaping belongs in the
// service layer. On top of the verbs this package provides the generic
// page/per_page pagination plumbing, the git-tree-commit helpers, and the
// authenticated-client constructors shared by both CLIs.
//
// It exists as F1 of the forge-agnostic foundation (#1017): consolidating the
// two previously-paralleled per-CLI transports here so a single GitHub Forge
// backend (F3) can be built on one implementation.
package githubapi
