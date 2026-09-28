// Package forge is the forge-agnostic domain seam for the Classroom 50 CLIs:
// neutral types (User, Org, Repo, Team, ...) and the Forge interface that a
// per-forge backend implements. It lets the domain packages talk about
// classroom operations without naming GitHub, so a second backend (GitLab,
// Gitea, ...) can be added by implementing Forge rather than editing every
// call site (#1017).
//
// F2 of the foundation deliberately defines only the neutral types plus a small,
// grounded core interface (identity). The teacher- and student-specific
// operation sets embed Forge and GROW package-by-package as F4 migrates each
// domain package onto the seam — so every operation's shape is validated by a
// real consumer (and the GitHub backend in F3) instead of being guessed up
// front. This is why the awkward operations (the tree-commit rebase loop, secret
// encryption, PR-as-issue labels, template-generate+stabilize, rulesets) are NOT
// modelled here yet: they arrive with the consumer that needs them.
package forge
