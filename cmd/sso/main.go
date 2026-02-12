package sso

import (
	"fmt"

	"github.com/napryag/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
