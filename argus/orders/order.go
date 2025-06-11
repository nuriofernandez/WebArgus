package orders

type Order struct {
	Url       string
	CheckType string
	Notify    NotificationRule
	Period    string
}

type NotificationRule struct {
	Phone   string
	Title   string
	Message string
}
