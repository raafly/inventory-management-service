package db

import (
	"fmt"
	"testing"

	"github.com/raafly/invetory-management/config"
)




func TestPostgres(t *testing.T) {
	config, _ := config.NewAppConfig()

	_ = NewDB(config)
	fmt.Println("end task")
}