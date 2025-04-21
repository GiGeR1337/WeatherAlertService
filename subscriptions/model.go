package subscriptions

type Subscription struct {
	Email     string `json:"email"`
	City      string `json:"city"`
	Condition string `json:"condition"`
}
