package inttegro

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v7/payout"
)

func TestPayoutsUseCanonicalTypedContracts(t *testing.T) {
	var setDestinationsBody struct {
		Destinations payout.Destinations `json:"destinations"`
	}
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/payouts/set_destinations":
			if err := json.NewDecoder(r.Body).Decode(&setDestinationsBody); err != nil {
				t.Fatalf("decode set destinations request: %v", err)
			}
			_, _ = io.WriteString(w, `{"settings":{"destinations":{"ghs":"fa_ghs"},"id":"ps_123"}}`)
		case "/payouts/settings":
			_, _ = io.WriteString(w, `{"settings":{"destinations":{"ghs":"fa_ghs"},"fx_enabled":true,"schedule":{"aging_spec":{"abide":"strict","label":"Seven days","t_plus":"168h"},"description":"Weekly payouts","interval":"weekly","name":"Weekly","schedule_on":"monday","type":"automatic"}}}`)
		case "/payouts/page":
			_, _ = io.WriteString(w, `{"page":{"number":1,"size":1,"payouts":[`+canonicalPayoutJSON+`]}}`)
		default:
			_, _ = io.WriteString(w, `{"payout":`+canonicalPayoutJSON+`}`)
		}
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	got, err := client.Payouts.Lookup(ctx, "po_123")
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if got.Error == nil || got.Error.Type != "network_error" || got.Error.OccurredAt.IsZero() {
		t.Fatalf("typed payout error = %#v", got.Error)
	}
	if len(got.BalanceTransactions) != 1 || got.BalanceTransactions[0] != "bt_123" {
		t.Fatalf("balance transactions = %#v", got.BalanceTransactions)
	}
	if got.CustomData["batch"] != "weekly" || got.SentAt == nil || got.FailedAt == nil {
		t.Fatalf("typed payout metadata/timestamps = %#v", got)
	}

	settings, err := client.Payouts.Settings(ctx)
	if err != nil {
		t.Fatalf("Settings() error = %v", err)
	}
	if settings.Destinations.GHS != "fa_ghs" || settings.Schedule == nil || settings.Schedule.AgingSpec.TPlus != "168h" {
		t.Fatalf("typed lookup settings = %#v", settings)
	}

	mutation, err := client.Payouts.SetDestinations(ctx, payout.Destinations{GHS: "fa_ghs"})
	if err != nil {
		t.Fatalf("SetDestinations() error = %v", err)
	}
	if setDestinationsBody.Destinations.GHS != "fa_ghs" || mutation.Destinations == nil || mutation.Destinations.GHS != "fa_ghs" {
		t.Fatalf("typed destination request/response = %#v / %#v", setDestinationsBody, mutation)
	}

	page, err := client.Payouts.Page(ctx, payout.PageParams{PageNumber: 1, PageSize: 256})
	if err != nil {
		t.Fatalf("Page() error = %v", err)
	}
	if page.Number != 1 || page.Size != 1 || len(page.Payouts) != 1 || page.Payouts[0].ID != "po_123" {
		t.Fatalf("typed payout page = %#v", page)
	}
}

const canonicalPayoutJSON = `{
  "amount":{"currency":"ghs","value":12500},
  "balance_transactions":["bt_123"],
  "custom_data":{"batch":"weekly"},
  "destination_id":"fa_ghs",
  "error":{"cause":"provider unavailable","message":"Payout failed","occurred_at":"2026-09-14T09:05:00Z","type":"network_error"},
  "execute_after":"2026-09-14T09:00:00Z",
  "failed_at":"2026-09-14T09:05:00Z",
  "id":"po_123",
  "initiated_at":"2026-09-14T08:55:00Z",
  "max_amount":{"currency":"ghs","value":12500},
  "sent_at":"2026-09-14T09:01:00Z",
  "status":"invalid"
}`
