package githubapi

import (
	"github.com/spf13/cobra"

	"github.com/foundation50/classroom50-cli-shared/contract"
	"github.com/foundation50/classroom50-cli-shared/ghauth"
)

// requiredScopes is the unified OAuth scope set both CLIs request, so
// authenticating for one covers the other (contract.RequiredOAuthScopes,
// issue #246). delete_repo stays opt-in for teardown.
var requiredScopes = contract.RequiredOAuthScopes()

// RequiredScopes returns the OAuth scopes the CLIs request beyond gh's
// defaults. Exposed for each CLI's login command.
func RequiredScopes() []string { return append([]string(nil), requiredScopes...) }

// RequireAuthClient returns a REST client, auto-running `gh auth login` when no
// token is set. commandName is the invoking CLI's user-facing name (e.g.
// "gh teacher" / "gh student"), surfaced in the auth prompt. Returned as the
// Client seam so domain code never names the concrete go-gh type.
func RequireAuthClient(cmd *cobra.Command, commandName string) (Client, error) {
	return ghauth.RequireClient(cmd.OutOrStdout(), cmd.ErrOrStderr(), ghauth.Options{
		RequiredScopes: requiredScopes,
		CommandName:    commandName,
	})
}
