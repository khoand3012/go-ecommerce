package main

import (
	"github.com/khoand3012/go-ecommerce/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}
