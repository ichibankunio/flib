package jsstorage

import (
	"regexp"
	"strconv"
	"strings"
	"syscall/js"
)

/*
func Confirm(txt string) {
	if js.Global().Get("window").Call("confirm", txt).Bool() {
		js.Global().Get("location").Set("href", "index.html")
	}
}
*/

/*
func SuggestApp() {
	if strings.Contains(js.Global().Get("navigator").Get("userAgent").String(), "Android") {
		// if r.MatchString(js.Global().Get("navigator").Get("userAgent").String()) {
	// if js.Global().Get("navigator").Get("userAgent").match("/(iPhone|iPod|Android.*Mobile)/i)") {
		if js.Global().Get("window").Call("confirm", "無料のアプリ版をダウンロードしませんか？\n全画面でサクサク動きます！\nPlayストアで開きますか？").Bool() {
			js.Global().Get("location").Set("href", "https://play.google.com/store/apps/details?id=com.ku20298.hashidebashi&pcampaignid=pcampaignidMKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1")
			// if !js.Global().Get("window").Call("open", "https://play.google.com/store/apps/details?id=com.ku20298.hashidebashi").Bool() {
			// }
		}
	}
}
*/

func HideLoading() {
	js.Global().Get("document").Call("getElementById", "loading").Get("style").Set("display", "none")
}

func IsMobile() bool {
	r := regexp.MustCompile("iPhone|iPod|Android")
	return r.MatchString(js.Global().Get("navigator").Get("userAgent").String()) 
}

func locationTo() {
	js.Global().Get("location").Set("href", "index.html")
}	


// func ShowQR(this js.Value, args []js.Value) interface{} {
//     js.Global().Call("alert", "hello")
//     return nil
// }


func GetInt(tag string) int {
	i, _ := strconv.Atoi(GetItem(tag).(js.Value).String())
	return i
}

func GetItem(tag string) interface{} {
	window := js.Global().Get("window")
	localStorage := window.Get("localStorage")
	
	return localStorage.Get(tag)
}

func SetItem(tag string, a interface{}) {
	window := js.Global().Get("window")
	localStorage := window.Get("localStorage")
	localStorage.Set(tag, a)
}

func SetArray(tag string, arr []int) {
	arrStr := ""
	for _, v := range arr {
		arrStr += strconv.Itoa(v) + ","
	}

	SetItem(tag, arrStr)
}

func GetArray(tag string) []int {
	arrStr := GetItem(tag).(js.Value).String()
	
	arr := []int{}
	for _, v := range strings.Split(arrStr, ",") {
		i, _ := strconv.Atoi(v)
		arr = append(arr, i)
	}
	
	return arr
}

func SetArray2(tag string, arr []interface{}) {
	window := js.Global().Get("window")
	jsonJS := window.Get("JSON")
	jsonStr := jsonJS.Call("stringify", js.ValueOf(arr)).String()
	SetItem(tag, jsonStr)
}

func GetArray2(tag string) int {
	window := js.Global().Get("window")
	jsonJS := window.Get("JSON")
	
	return jsonJS.Call("parse", GetItem(tag).(js.Value).String()).Length()
}

func Clear() {
	window := js.Global().Get("window")
	localStorage := window.Get("localStorage")
	localStorage.Call("clear")
}
