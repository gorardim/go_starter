package component

import (
	"fmt"
	"sync/atomic"
	"time"

	"app/pkg/randx"
)

var id int32 = 0

func GenerateQuestionId() string {
	atomic.AddInt32(&id, 1)
	if id >= 100000 {
		id = 0
	}

	now := time.Now()
	return fmt.Sprintf("%s%03d%s%05d", now.Format("06"), now.YearDay(), randx.Digit(4), id)
}
