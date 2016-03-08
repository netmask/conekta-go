package swagger

import (
)

type Client struct {
    Id  string  `json:"id,omitempty"`
    Object  string  `json:"object,omitempty"`
    Livemode  string  `json:"livemode,omitempty"`
    CreatedAt  string  `json:"created_at,omitempty"`
    Name  string  `json:"name,omitempty"`
    Email  string  `json:"email,omitempty"`
    Phone  string  `json:"phone,omitempty"`
    DefaultCard  string  `json:"default_card,omitempty"`
    BillingAddress  BillingAddress  `json:"billing_address,omitempty"`
    ShippingAddress  BillingAddress  `json:"shipping_address,omitempty"`
    Cards  []TokenizedCard  `json:"cards,omitempty"`
    Subscriptions  []Subscription  `json:"subscriptions,omitempty"`
    
}
