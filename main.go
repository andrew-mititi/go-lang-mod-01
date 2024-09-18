package main

const APP_NAME string = "golang module 1"


type engine struct {
	electric bool
}

type Vehicle struct {
	engine
	model, color,make string

}


func main(){ 

	japan_vehicle := Vehicle{
		engine: engine{
			electric: false,
		},
		model: "WRX STI",
		color: "blue",
		make: "Subaru",
	}

	_ = japan_vehicle


	// functions

}