package entities

type Order struct {
	ID int32
	UserID int32
	Amount float32
	Status string
}

func NewOrder(userID int32, amount float32, status string) *Order{
	return &Order{UserID: userID, Amount: amount, Status: status}
}

func (o *Order) UpdateOrder(id int32, userID int32, amount float32, stauts string) {
	o.ID = id
	o.UserID = userID
	o.Amount = amount
	o.Status = stauts	
}