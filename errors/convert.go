package errors

import "github.com/gogf/gf/util/gvalid"

func GValidError2MapData(err gvalid.Error) map[string][]string {
	data := make(map[string][]string)
	for _, d := range err.Items() {
		for k, v := range d {
			m := make([]string, 0)
			for _, val := range v {
				m = append(m, val)
			}
			data[k] = m
		}
	}
	return data
}
