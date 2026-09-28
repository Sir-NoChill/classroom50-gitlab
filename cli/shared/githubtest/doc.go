// Package githubtest provides shared, exported test-support helpers for the
// CLIs' white-box tests: a real REST client wired to an httptest.Server, and an
// in-memory fake implementing the shared githubapi.Client seam. It (with
// githubapi) keeps the go-gh dependency confined so domain tests never
// construct a go-gh client directly.
package githubtest
