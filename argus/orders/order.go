package orders

type Order struct {
	Id        string
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
