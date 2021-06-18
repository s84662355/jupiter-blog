package helper

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

//字符串格式1,7,5,7,8 以切片操作 json序列化为[1,4,5,6]
type SplitIntSlices []int

func (t SplitIntSlices) New(a []int) *SplitIntSlices {
	var tt SplitIntSlices = a
	return &tt
}

func (t SplitIntSlices) MarshalJSON() ([]byte, error) {
	if len(t) == 0 {
		return []byte("[]"), nil
	}
	data, err := json.Marshal([]int(t))
	return data, err
}

func (t *SplitIntSlices) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*t = []int{}
		return nil
	}
	tt := []int{}
	err := json.Unmarshal(data, &tt)
	*t = tt
	return err
}

func (t SplitIntSlices) String() string {
	if len(t) == 0 {
		return ""
	}
	arr := make([]string, len(t))
	for index, v := range t {
		arr[index] = fmt.Sprintf("%d", v)
	}
	res := strings.Join(arr, ",")
	return res
}

func (t SplitIntSlices) Value() (driver.Value, error) {
	return t.String(), nil
}

func (t *SplitIntSlices) Scan(v interface{}) error {
	vv, _ := v.([]byte)
	ss := string(vv)
	ss = strings.Replace(ss, " ", "", -1)
	if ss == "" {
		*t = []int{}
		return nil
	}
	arr := strings.Split(ss, ",")
	tt := make([]int, len(arr))
	for index, v := range arr {
		tt[index] = Str2Int(v)
	}
	*t = tt
	return nil
}

func (t SplitIntSlices) Get() []int {
	return []int(t)
}

func (t SplitIntSlices) Len() int {
	return len(t)
}
