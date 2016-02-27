package providers

import (
	"net/url"
)

type ProviderData struct {
	ProviderName   string
	ClientID       string
	ClientSecret   string
	DiscoveryURL   *url.URL
	LoginURL       *url.URL
	RedeemURL      *url.URL
	ProtectedResource *url.URL
	ProfileURL     *url.URL
	ValidateURL    *url.URL
	Scope          string
	ApprovalPrompt string
}

func (p *ProviderData) Data() *ProviderData { return p }
