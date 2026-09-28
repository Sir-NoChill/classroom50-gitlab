package githubapi

import (
	"github.com/foundation50/classroom50-cli-shared/ghutil"
	ghapi "github.com/foundation50/classroom50-cli-shared/githubapi"
)

// EnablePages configures a GitHub Pages site on owner/repo; alreadyEnabled
// reports a 409 (a site exists, left untouched). See ghutil.EnablePages.
func EnablePages(c ghapi.Client, owner, repo string, body ghutil.PagesCreateBody) (alreadyEnabled bool, err error) {
	return ghutil.EnablePages(ghapi.REST(c), owner, repo, body)
}
