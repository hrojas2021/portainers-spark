package main

import "spark/challenge/client/app"

const PORT = ":10000"

func main() {
	a := app.App{}
	a.Initialize()
	a.Run(PORT)
}
