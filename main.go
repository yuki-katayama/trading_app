package main

import (
	"fmt"

	"github.com/my_user/my_repo/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
