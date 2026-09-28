package githubapi

import "github.com/cli/go-gh/v2/pkg/api"

// rest recovers the concrete *api.RESTClient backing a Client. Every Client in
// either binary is the go-gh client from RequireAuthClient / NewClient (or the
// test fake in cli/shared/githubtest, which embeds one), so the assertion
// holds; it panics otherwise — a programming error, since the ghutil/gittree
// operations below require the concrete type and no other impl can satisfy them.
func rest(c Client) *api.RESTClient {
	rc, ok := c.(*api.RESTClient)
	if !ok {
		panic("githubapi: Client is not a *api.RESTClient; shared-module operations require the concrete go-gh client")
	}
	return rc
}

// REST is the exported escape hatch to the concrete *api.RESTClient, for the
// CLIs' own GitHub-specific operations that build on gittree/ghutil helpers
// taking the concrete client. Domain code should prefer the Client verbs.
func REST(c Client) *api.RESTClient { return rest(c) }
