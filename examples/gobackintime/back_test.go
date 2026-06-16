package gobackintime

import (
	"fmt"
	"testing"
	"testing/synctest"
	"time"
)

func TestTimeMovesBackwards(t *testing.T) {
	t.Skip()
	ts := time.Now()
	for {
		time.Sleep(time.Second)
		elapsed := time.Until(ts)
		synctest.Test(t, func(t *testing.T) {
			time.Sleep(time.Until(ts.Add(elapsed)))
			fmt.Printf("Hello from %s!\n", time.Now().Truncate(time.Second).Format(time.Stamp))
		})
	}
}
