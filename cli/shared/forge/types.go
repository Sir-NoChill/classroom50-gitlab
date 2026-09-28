package forge

// ForgeKind identifies which forge backend a Forge talks to. Today only GitHub
// is implemented; the type exists so config and code can name the backend as
// more are added (#1017). It is the domain-level counterpart of the optional
// `forge` schema discriminator (deferred to F6).
type ForgeKind string

const (
	// ForgeGitHub is the GitHub backend — the only one implemented today and
	// the default when a config omits the discriminator.
	ForgeGitHub ForgeKind = "github"
)

// User is a forge account (a person). Login is the forge-unique handle; ID is
// the forge's stable numeric identifier (0 when the endpoint doesn't expose
// one).
type User struct {
	Login string
	ID    int64
}

// Org is a forge organization / group namespace that owns repositories and
// teams. Login is its unique handle.
//
// (Grows in F4 as org operations migrate onto the Forge seam — kept minimal on
// purpose; see the package doc.)
type Org struct {
	Login string
}

// Repo is a forge repository. Owner is the org/user login; Name is the repo
// name within that owner; DefaultBranch is the branch the forge reports as
// default; Private reports visibility.
//
// (Grows in F4.)
type Repo struct {
	Owner         string
	Name          string
	DefaultBranch string
	Private       bool
}

// Team is a forge team/group within an org. Slug is the URL-safe identifier,
// Name the display name, ID the forge's numeric identifier (0 when unused).
//
// (Grows in F4.)
type Team struct {
	Org  string
	Slug string
	Name string
	ID   int64
}

// TeamRole is a member's role within a team (the two roles every forge models).
type TeamRole string

const (
	TeamRoleMember     TeamRole = "member"
	TeamRoleMaintainer TeamRole = "maintainer"
)

// RepoPermission is a collaborator's permission level on a repository. Forges
// differ in the exact ladder; these are the GitHub set and the neutral names
// the domain uses — a backend maps them to its own vocabulary.
type RepoPermission string

const (
	RepoPermissionPull     RepoPermission = "pull"
	RepoPermissionTriage   RepoPermission = "triage"
	RepoPermissionPush     RepoPermission = "push"
	RepoPermissionMaintain RepoPermission = "maintain"
	RepoPermissionAdmin    RepoPermission = "admin"
)
