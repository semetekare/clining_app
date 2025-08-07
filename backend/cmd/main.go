package main

import (
	server "clining_app/app"

)

func main() {
	// viper.SetConfigFile(".env")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }
	app := new(server.App)
	app.Run("8000")
}
