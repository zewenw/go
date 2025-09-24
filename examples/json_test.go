package examples

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestJson(t *testing.T) {
	t.Run("Json test", func(t *testing.T) {
		bolB, _ := json.Marshal(true)
		fmt.Println(string(bolB))

		intB, _ := json.Marshal(1)
		fmt.Println(string(intB))

		fltB, _ := json.Marshal(2.34)
		fmt.Println(string(fltB))

		strB, _ := json.Marshal("gopher")
		fmt.Println(string(strB))

		slcD := []string{"apple", "peach", "pear"}
		slcB, _ := json.Marshal(slcD)
		fmt.Println(string(slcB))

		mapD := map[string]any{"apple": "5", "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		fmt.Println(string(mapB))

		res1 := &response1{
			Page: 1,
			Fruits: []string{
				"apple", "peach", "pear",
			},
		}
		res1B, _ := json.Marshal(res1)
		fmt.Println(string(res1B))

		res2 := &response2{
			Page: 1,
			Fruits: []string{
				"apple", "peach", "pear",
			},
		}
		res2B, _ := json.Marshal(res2)
		fmt.Println(string(res2B))

		byt := []byte(`{"num":{"key1" : "1", "key2" : 2.13},"strs":["a","b"]}`)

		var dat map[string]interface{}
		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)

		num := dat["num"].(map[string]interface{})["key2"].(float64)
		fmt.Println(num)

		str1 := dat["strs"].([]interface{})[0].(string)
		fmt.Println(str1)

		str := `{"page": 1, "fruits": ["apple", "peach"]}`
		res := response1{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println(res)
		fmt.Println(res.Fruits[0])

		enc := json.NewEncoder(os.Stdout)
		d := map[string]int{"apple": 5, "lettuce": 7}
		enc.Encode(d)

		dec := json.NewDecoder(strings.NewReader(str))
		res3 := response2{}
		dec.Decode(&res3)
		fmt.Println(res3)
	})
}
