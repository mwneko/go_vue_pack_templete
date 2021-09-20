package backend

// need separate to models
type company struct {
	Company_id   int    `json:"id"`
	Company_name string `json:"name"`
}

type customers struct {
	User_id      int    `json:"id"`
	Login        string `json:"login"`
	Password     string `json:"psw"`
	Name         string `json:"name"`
	Company_id   int    `json:"company_id"`
	Credit_cards string `json:"credit"`
}

type orders struct {
	ID          int    `json:"id"`
	Created_at  string `json:"created_at"`
	Order_name  string `json:"order_name"`
	Customer_id string `json:"customer_id"`
}
type deliveries struct {
	ID                 int `json:"id"`
	Order_item_id      int `json:"order_item_id"`
	Delivered_quantity int `json:"quantity"`
}
type order_items struct {
	ID             int     `json:"id"`
	Order_id       int     `json:"order_id"`
	Price_per_unit float64 `json:"price_per_unit"`
	Quantity       int     `json:"quantity"`
	Product        string  `json:"product"`
}

type sample struct {
	ID               int     `json:"id"`
	Order_name       string  `json:"order"`
	Customer_company string  `json:"company"`
	Customer_name    string  `json:"customer"`
	Order_date       string  `json:"date"`
	Amount           float64 `json:"amount"`
	Total            float64 `json:"total"`
}

// This part need get from db
// This is mockup data
var sampleData = []sample{
	{ID: 1, Order_name: "PO #001-I", Customer_company: "Roga & Kopyta", Customer_name: "ivan", Order_date: "2020-01-02", Amount: 10.1, Total: 50.9},
	{ID: 2, Order_name: "PO #002-I", Customer_company: "Roga & Kopyta", Customer_name: "ivan", Order_date: "2020-01-02", Amount: 10.1, Total: 50.9},
	{ID: 3, Order_name: "PO #003-I", Customer_company: "Roga & Kopyta", Customer_name: "ivan", Order_date: "2020-01-02", Amount: 20.1, Total: 250.9},
	{ID: 4, Order_name: "PO #004-I", Customer_company: "Roga & Kopyta", Customer_name: "ivan", Order_date: "2020-01-02", Amount: 20.1, Total: 150.9},
	{ID: 5, Order_name: "PO #005-I", Customer_company: "Roga & Kopyta", Customer_name: "ivan", Order_date: "2020-01-02", Amount: 30.1, Total: 150.9},
	{ID: 6, Order_name: "PO #001-I", Customer_company: "Pupkin & Co", Customer_name: "petr", Order_date: "2020-01-02", Amount: 40.1, Total: 150.9},
	{ID: 7, Order_name: "PO #002-I", Customer_company: "Pupkin & Co", Customer_name: "petr", Order_date: "2020-01-02", Amount: 50.1, Total: 450.9},
	{ID: 8, Order_name: "PO #003-I", Customer_company: "Pupkin & Co", Customer_name: "petr", Order_date: "2020-01-02", Amount: 60.1, Total: 350.9},
	{ID: 9, Order_name: "PO #004-I", Customer_company: "Pupkin & Co", Customer_name: "petr", Order_date: "2020-01-02", Amount: 80.1, Total: 250.9},
	{ID: 10, Order_name: "PO #005-I", Customer_company: "Pupkin & Co", Customer_name: "petr", Order_date: "2020-01-02", Amount: 100.1, Total: 150.9},
}
