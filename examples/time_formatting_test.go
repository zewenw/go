package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormatting(t *testing.T) {

	t.Run("time formatting", func(t *testing.T) {
		p := fmt.Println
		now := time.Now()
		p(now.Format(time.RFC3339))

		t1, e := time.Parse(
			time.RFC3339,
			"2012-11-01T22:08:41+00:00")
		p(t1)

		// TODO this is interesting, if have time, read the source code
		p(now.Format("3:04PM"))
		p(now.Format("Mon Jan _2 15:04:05 2006"))
		p(now.Format("2006-01-02T15:04:05.999999-07:00"))
		form := "3 04 PM"
		t2, e := time.Parse(form, "8 41 PM")
		p(t2)

		fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second())

		ansic := "Mon Jan _2 15:04:05 2006"
		_, e = time.Parse(ansic, "8:41PM")
		p(e)
		p(e)
	})
}
