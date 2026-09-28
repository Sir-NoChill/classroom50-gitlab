package forge

import "context"

// Forge is the forge-agnostic core seam every backend implements: which forge
// it is, plus identity. It is deliberately small.
//
// The teacher- and student-specific operation sets will EMBED Forge and grow as
// F4 migrates each domain package onto the seam, e.g.
//
//	type TeacherForge interface {
//		Forge
//		// ... org / team / repo / secret / release operations, added as the
//		//     domain packages that use them migrate.
//	}
//
// Defining them at the consumer (as each package migrates) rather than up front
// keeps every operation's shape validated by a real caller and the GitHub
// backend, instead of guessed — see the package doc.
type Forge interface {
	// Kind reports which forge backend this is.
	Kind() ForgeKind

	// CurrentUser returns the authenticated account.
	CurrentUser(ctx context.Context) (User, error)

	// GetUser looks up an account by its login handle.
	GetUser(ctx context.Context, login string) (User, error)
}
