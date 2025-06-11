package list

type OrderResponse struct {
	Id        string                   `json:"id"`
	Url       string                   `json:"url"`
	CheckType string                   `json:"checkType"`
	Notify    NotificationRuleResponse `json:"notify"`
	Period    string                   `json:"period"`
}

type NotificationRuleResponse struct {
	Phone   string `json:"phone"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
