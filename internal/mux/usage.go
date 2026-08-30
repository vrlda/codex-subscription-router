package mux

import (
	"context"
	"encoding/json"
	"net/http"
)

const usageStatusURL = "https://chatgpt.com/backend-api/wham/usage"

// UsageStatus returns the full native Wham payload for one subscription so
// Codex can update model availability, including Luna Reserve state.
func (m *Multiplexer) UsageStatus(ctx context.Context, accountID string) (json.RawMessage, error) {
	account, err := m.accountForAuthenticatedRequest(accountID)
	if err != nil {
		return nil, err
	}
	return requestAccountJSON(
		ctx,
		m.profileClient,
		m.usageStatusEndpoint,
		http.MethodGet,
		account,
		nil,
		map[string]string{"OAI-App-Brand": "codex"},
	)
}
