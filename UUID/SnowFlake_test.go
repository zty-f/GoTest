package UUID

import (
	"testing"
	"time"
)

func TestIdGeneratorService(t *testing.T) {
	svr := &IdGeneratorService{
		dataCenterId: 1,
		workerId:     1,
		startTime:    time.Now().AddDate(-10, 0, -1),
	}
	id, err := svr.GenID()
	if err != nil {
		t.Errorf("GenID() error = %v", err)
		return
	}
	if id <= 0 {
		t.Errorf("GenID() id = %v, want > 0", id)
		return
	}
}
