package models

type PerformanceMetric struct {
    ID          string  `json:"id" db:"id"`
    UserID      string  `json:"user_id" db:"user_id"`
    CadenceID   string  `json:"cadence_id" db:"cadence_id"`
    ConversionRate float64 `json:"conversion_rate" db:"conversion_rate"`
    ResponseRate float64  `json:"response_rate" db:"response_rate"`
    FollowUpRate float64  `json:"follow_up_rate" db:"follow_up_rate"`
}

