package uploader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type IPBoundTransport string

type IPv4Addr string

func (a IPv4Addr) Network() string {
	return "tcp"
}

func (a IPv4Addr) String() string {
	return string(a)
}

//
// func telemFileToQueryString(file FileInfo) string {
// 	b, err := ioutil.ReadFile(file.GetTextPath())
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return telemToQueryString(string(b))
// }
//
// func telemToQueryString(contents string) string {
// 	lines := strings.Split(contents, "\n")
//
// 	v := url.Values{}
// 	for _, line := range lines {
// 		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
// 		if len(parts) != 2 {
// 			continue
// 		}
// 		v.Set(parts[0], parts[1])
// 	}
//
// 	return v.Encode()
// }

func telemFileToJSONString(file FileInfo) string {
	b, err := ioutil.ReadFile(file.GetTextPath())

	if err != nil {
		panic(err)
	}

	return telemToJSONString(string(b))
}

func telemToJSONString(contents string) string {
	lines := strings.Split(contents, "\n")

	m := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
		if len(parts) != 2 {
			continue
		}
		m[parts[0]] = parts[1]
	}

	b, _ := json.Marshal(m)

	return string(b)
}

// Upload - upload a file
func Upload(f FileInfo, bindIP string) error {

	fmt.Printf("local address: %s\n", bindIP)

	client := http.Client{
		Transport: &http.Transport{
			Proxy: nil,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				LocalAddr: IPv4Addr(bindIP),
			}).Dial,
		},
	}

	url := "http://homer.compassuav.com:3000/uploadImage?telem=" + url.QueryEscape(telemFileToJSONString(f))

	imageFile, err := os.Open(f.GetImagePath())
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	//resp, err := client.Post(url, "image/jpeg", nil)

	request, _ := http.NewRequest("PUT", url, imageFile)
	//request.Header.Set("Content-length", imageFile.Length())
	request.Header.Set("Content-type", "image/jpeg")

	resp, err := client.Do(request)

	if err == nil {
		fmt.Println("SUP")
		fmt.Println("Status: " + resp.Status)
		fmt.Printf("ContentLength: %d\n", resp.ContentLength)
	}

	return err
}
