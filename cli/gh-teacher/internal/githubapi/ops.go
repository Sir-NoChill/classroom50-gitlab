package githubapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	ghapi "github.com/foundation50/classroom50-cli-shared/githubapi"
	"github.com/foundation50/classroom50-cli-shared/gittree"
	"github.com/foundation50/gh-teacher/internal/orgpolicy"
)

// OrgPlan reads GET /orgs/{org} and returns the billing plan name
// ("free"/"team"/"enterprise"), empty when the token lacks billing visibility.
// The error is returned raw so callers can classify it.
func OrgPlan(c ghapi.Client, org string) (string, error) {
	path := fmt.Sprintf("orgs/%s", url.PathEscape(org))
	var resp struct {
		Plan struct {
			Name string `json:"name"`
		} `json:"plan"`
	}
	if err := c.Get(path, &resp); err != nil {
		return "", err
	}
	return resp.Plan.Name, nil
}

// orgBudgetsPath is the org billing-budgets endpoint (list + create). Kept here
// so the read and write helpers can't drift.
func orgBudgetsPath(org string) string {
	return fmt.Sprintf("organizations/%s/settings/billing/budgets", url.PathEscape(org))
}

// budgetsListResponse is the GET /organizations/{org}/settings/billing/budgets
// envelope. GitHub returns the budgets under a "budgets" key; unknown fields
// are ignored so the reader tolerates schema growth.
type budgetsListResponse struct {
	Budgets []orgpolicy.Budget `json:"budgets"`
}

// ListOrgBudgets reads the org's billing budgets. The error is returned raw so
// callers can classify a 403/404 (no billing visibility / not entitled) as an
// advisory, not a hard failure. Needs org Administration: Read.
func ListOrgBudgets(c ghapi.Client, org string) ([]orgpolicy.Budget, error) {
	var resp budgetsListResponse
	if err := c.Get(orgBudgetsPath(org), &resp); err != nil {
		return nil, err
	}
	return resp.Budgets, nil
}

// CreateOrgActionsBudgetCap POSTs the desired $0 hard-stop Actions budget for
// the org, returning the HTTP status. Callers must only invoke this when no
// Actions budget exists (GitHub allows one budget per scope+SKU) — this helper
// never modifies an existing budget. Needs org Administration: Read and write.
func CreateOrgActionsBudgetCap(c ghapi.Client, org string) (int, error) {
	body, err := json.Marshal(map[string]any{
		"budget_amount":         0,
		"prevent_further_usage": true,
		"budget_scope":          orgpolicy.BudgetScopeOrg,
		"budget_type":           orgpolicy.BudgetTypeProductPricing,
		"budget_product_sku":    orgpolicy.BudgetProductSKUActions,
	})
	if err != nil {
		return 0, fmt.Errorf("encode body: %w", err)
	}
	resp, err := c.Request(http.MethodPost, orgBudgetsPath(org), bytes.NewReader(body))
	if err != nil {
		return 0, err
	}
	defer func() { _ = resp.Body.Close() }()
	_, _ = io.Copy(io.Discard, resp.Body)
	return resp.StatusCode, nil
}

// CommitWithRebase runs the optimistic tree-commit-with-rebase loop.
func CommitWithRebase(
	c ghapi.Client,
	owner, repo, branch, message string,
	build func(parentSHA string) (gittree.Change, error),
	classify404 func(error) error,
) (string, error) {
	return gittree.CommitWithRebase(ghapi.REST(c), owner, repo, branch, message, build, classify404)
}
