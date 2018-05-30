package time_test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.Local, time.UTC)
	t.Log(time.Now().UTC(), time.Now().Local(), time.Now().Location())
	t.Log(time.Now())
}
