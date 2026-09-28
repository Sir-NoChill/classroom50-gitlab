package githubapi

import (
	"time"

	"github.com/foundation50/classroom50-cli-shared/ghutil"
	"github.com/foundation50/classroom50-cli-shared/gittree"
)

// CurrentUser returns the authenticated user's login and id.
func CurrentUser(c Client) (login string, id int64, err error) {
	return ghutil.CurrentUser(rest(c))
}

// ReadTemplatePRBody reads the first existing native pull_request_template.md
// from owner/repo at ref as the Feedback PR body, or ("", false) to fall back
// to the built-in body. Best-effort/fail-open (see ghutil.ReadTemplatePRBody).
func ReadTemplatePRBody(c Client, owner, repo, ref string) (string, bool) {
	return ghutil.ReadTemplatePRBody(rest(c), owner, repo, ref)
}

// SetCollaborator adds username to owner/repo at the given permission,
// returning the resulting HTTP status.
func SetCollaborator(c Client, owner, repo, username, permission string) (int, error) {
	return ghutil.SetCollaborator(rest(c), owner, repo, username, permission)
}

// WaitForStableBranch polls until owner/repo's branch is readable after a
// templated-repo creation, smoothing replication lag.
func WaitForStableBranch(c Client, owner, repo, branch string) error {
	return ghutil.WaitForStableBranch(rest(c), owner, repo, branch)
}

// ResolveSettledDefaultBranch waits out the async template-copy lag and returns
// the branch that actually materialized (not the transiently-reported
// default_branch), falling back to `fallback`.
func ResolveSettledDefaultBranch(c Client, owner, repo, fallback string) string {
	return ghutil.ResolveSettledDefaultBranch(rest(c), owner, repo, fallback, 20, 250*time.Millisecond)
}

// UploadBlobs uploads file contents as git blobs, returning their tree entries.
func UploadBlobs(c Client, owner, repo string, files map[string]string) ([]gittree.TreeEntry, error) {
	return gittree.UploadBlobs(rest(c), owner, repo, files)
}

// CommitWithFreshRepoRetry commits a prepared entry set, retrying while a
// freshly created repo's git-data APIs lag.
func CommitWithFreshRepoRetry(
	c Client,
	owner, repo, branch, message string,
	entries []gittree.TreeEntry,
	cfg gittree.FreshRepoRetry,
) (string, error) {
	return gittree.CommitWithFreshRepoRetry(rest(c), owner, repo, branch, message, entries, cfg)
}
