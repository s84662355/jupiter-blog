package helper

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/uniplaces/carbon"
	"io"
	"math"
	"math/rand"
	"net"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//计算某月的第一天
func GetMonthFirstDay(now time.Time) time.Time {
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	return firstOfMonth
}

//计算某月的最后一天
func GetMonthLastDay(now time.Time) time.Time {
	firstOfMonth := GetMonthFirstDay(now)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return lastOfMonth
}

//去除两边空格
func Trim(str *string) {
	*str = strings.Trim(*str, " ")
	*str = strings.Replace(*str, "\n", "", -1)
}

//去除emoji
func Emoji(str *string) {
	emojiRx := regexp.MustCompile("[^\u0000-\uFFFF]")
	*str = emojiRx.ReplaceAllString(*str, "")
}

func IsMobile(mm string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mm)
}

func Encryption(data string, secretkey string, hm uint64) string {
	return Md5(fmt.Sprintf("%s%s%d", data, secretkey, hm))
}

//生成32位md5字串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// 判断是否为slcie数据
func IsSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return val, ok
}

//是否是文件
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

//获取随机数  数字和文字
func GetRandomBoth(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取本机ip
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func Recover() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

//sha1加密
func Sha1En(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page int, prepage int, nums int64) map[string]int {

	var firstpage int //前一页地址
	var nextpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		nextpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		nextpage = page + 1
	default:
		/*
		   pages = make([]int, int(math.Min(5, float64(totalpages))))
		   for i, _ := range pages {
		       pages[i] = i + 1
		   }
		*/
		firstpage = int(math.Max(float64(1), float64(page-1)))
		nextpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]int)
	//paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["nextpage"] = nextpage
	paginatorMap["currpage"] = page
	paginatorMap["total"] = int(nums)
	paginatorMap["size"] = int(prepage)

	return paginatorMap
}

//截取指定字符子串
func SubstrContains(s, substr string) string {
	n := strings.Index(s, substr)
	return s[n:]
}

func GetScopeDay(start_time, stop_time string) (args []string) {

	tm1, _ := time.Parse("2006-01-02 00:00:00", start_time)
	tm2, _ := time.Parse("2006-01-02 00:00:00", stop_time)

	fmt.Println(tm1, tm2)

	sInt := tm1.Unix()
	eInt := tm2.Unix()
	if sInt > eInt {
		return
	}
	for {
		st := time.Unix(sInt, 0).Format("2006-01-02")
		args = append(args, st)
		sInt += 86400
		if sInt > eInt {
			return
		}
	}
}

func GetScopeMonth(start_time, stop_time string) (args []string) {
	tm1, _ := time.Parse("2006-01-02 00:00:00", start_time)
	tm2, _ := time.Parse("2006-01-02 00:00:00", stop_time)
	sInt := tm1.Unix()
	eInt := tm2.Unix()
	if sInt > eInt {
		return
	}
	sCarbon := carbon.NewCarbon(tm1)
	for {
		args = append(args, sCarbon.Format("2006-01"))
		sCarbon = sCarbon.AddMonth()
		if sCarbon.Unix() > eInt {
			return
		}
	}
	return
}

func GetScopeYear(start_time, stop_time string) (args []string) {
	tm1, _ := time.Parse("2006-01-02 00:00:00", start_time)
	tm2, _ := time.Parse("2006-01-02 00:00:00", stop_time)
	sInt := tm1.Unix()
	eInt := tm2.Unix()
	if sInt > eInt {
		return
	}
	sCarbon := carbon.NewCarbon(tm1)
	for {
		args = append(args, sCarbon.Format("2006"))
		sCarbon = sCarbon.AddYear()
		if sCarbon.Unix() > eInt {
			return
		}
	}
	return
}
