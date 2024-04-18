package main

type Product struct {
	Id   int
	Name string
	Cost float32
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		grapes := PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Grapes",
				Cost: 50,
			},
			Expiry: "2 Days",
		}
	*/
	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")

	prod := Product{}
	// prodPtr := &Product{}
	prodPtr := new(Product)
}
