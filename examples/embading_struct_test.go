package examples

import (
	"fmt"
	"testing"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func TestEmbedStruct(t *testing.T) {
	t.Run("embedding struct", func(t *testing.T) {
		co := container{
			base: base{
				num: 1,
			},
			str: "some name",
		}
		fmt.Printf("co = {num: %v, str: %v}\n", co.num, co.str)

		fmt.Println("also num:", co.base.num)

		fmt.Println("describe:", co.describe())
		
		type describer interface {
			describe() string
		}

		var d describer = co
		fmt.Println("describer:", d.describe())

	})
}
