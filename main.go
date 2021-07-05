package main

import (
	"bytes"
	"fmt"
	"github.com/valyala/fastjson"
	"net/http"
	"os"
	"strings"
)

const uri = "https://api.steamcmd.net/v1/info/%s"

const (
	linuxDepot  = "258552"
	commonDepot = "258554"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}

	appid := os.Args[1]
	fmt.Printf("Looking up information for app id %s\n", appid)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(uri, appid), nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	_, _ = buf.ReadFrom(res.Body)
	res.Body.Close()

	v, err := fastjson.Parse(buf.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	buildId := strings.Trim(v.Get("data", appid, "depots", "branches", "public", "buildid").String(), `"`)
	linuxManifestId := strings.Trim(v.Get("data", appid, "depots", linuxDepot, "manifests", "public").String(), `"`)
	buildUpdatedTime := strings.Trim(v.Get("data", appid, "depots", "branches", "public", "timeupdated").String(), `"`)
	commonManifestId := strings.Trim(v.Get("data", appid, "depots", commonDepot, "manifests", "public").String(), `"`)

	fmt.Printf("::set-output name=common_manifest_id::%s\n", commonManifestId)
	fmt.Printf("::set-output name=linux_manifest_id::%s\n", linuxManifestId)
	fmt.Printf("::set-output name=build_id::%s\n", buildId)
	fmt.Printf("::set-output name=build_updated_time::%s\n", buildUpdatedTime)
}
