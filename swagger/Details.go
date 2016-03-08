package swagger

import (
)

type Details struct {
    Name  string  `json:"name,omitempty"`
    Phone  string  `json:"phone,omitempty"`
    Email  string  `json:"email,omitempty"`
    Customer  Customer  `json:"customer,omitempty"`
    LineItems  []LineItem  `json:"line_items,omitempty"`
    BillingAddress  BillingAddress  `json:"billing_address,omitempty"`
    
}
