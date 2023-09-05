package async

import (
	"sync"

	"gorm.io/gorm"
)

func AsyncDBQuery(f func() *gorm.DB, wait *sync.WaitGroup) *gorm.DB {
	var result *gorm.DB
	go func() {
		result = f()
		defer wait.Done()
	}()
	return result
}
