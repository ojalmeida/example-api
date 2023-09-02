//go:generate swag init
package main

import "example-api/cmd"



// @title			Example API
// @version			1.0
// @description		This is a sample api written in Golang

func main() {

	cmd.Execute()

}
