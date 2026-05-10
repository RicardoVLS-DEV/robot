package main

import (
	"fmt"
	"robot/internal/app"
)

func main() {
	if err := app.LoadFromExcel("form.xlsx"); err != nil {
		fmt.Print(err)
		return
	}
}






