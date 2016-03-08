package swagger

import (
)

type Charge struct {
    Id  string  `json:"id,omitempty"`
    MonthlyInstallments  string  `json:"monthly_installments,omitempty"`
    Refunds  *[]Refunds  `json:"refunds,omitempty"`
    Livemode  bool  `json:"livemode,omitempty"`
    Cash  *Cash  `json:"cash,omitempty"`
    Description  string  `json:"description,omitempty"`
    Amount  int32  `json:"amount,omitempty"`
    Currency  string  `json:"currency,omitempty"`
    ReferenceId  string  `json:"reference_id,omitempty"`
    Bank  *ChargeBank  `json:"bank,omitempty"`
    Card  string  `json:"card,omitempty"`
    FailureCode  string  `json:"failure_code,omitempty"`
    FailureMessage  string  `json:"failure_message,omitempty"`
    PaymentMethod  *PaymentMethod  `json:"payment_method,omitempty"`
    Details  Details  `json:"details,omitempty"`
    
}
