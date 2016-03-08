package swagger

import (
)

type Plan struct {
    Id  string  `json:"id,omitempty"`
    Object  string  `json:"object,omitempty"`
    Livemode  bool  `json:"livemode,omitempty"`
    CreatedAt  string  `json:"created_at,omitempty"`
    Name  string  `json:"name,omitempty"`
    Amount  string  `json:"amount,omitempty"`
    Currency  string  `json:"currency,omitempty"`
    Interval  int32  `json:"interval,omitempty"`
    Frequency  int32  `json:"frequency,omitempty"`
    IntervalTotalCount  int32  `json:"interval_total_count,omitempty"`
    TrialPeriodDays  int32  `json:"trial_period_days,omitempty"`
    
}
