package main

import "automatic_api/hvc"

func main() {
	err := hvc.InitDB(false)
	if err != nil {
		panic(err)
	}

}
