package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/valyala/fastjson"
)

const uri = "https://api.steamcmd.net/v1/info/%s"

const (
	linuxDepot  = "258552"
	commonDepot = "258554"
)

var (
	appID  = flag.String("app-id", "", "App ID")
	branch = flag.String("branch", "public", "Branch")
)

func main() {
	flag.Parse()

	if *appID == "" {
		fmt.Println("App ID is required")
		os.Exit(-1)
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(uri, *appID), nil)
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

	buildId := strings.Trim(v.Get("data", *appID, "depots", "branches", *branch, "buildid").String(), `"`)
	linuxManifestId := strings.Trim(v.Get("data", *appID, "depots", linuxDepot, "manifests", *branch, "gid").String(), `"`)
	buildUpdatedTime := strings.Trim(v.Get("data", *appID, "depots", "branches", *branch, "timeupdated").String(), `"`)
	commonManifestId := strings.Trim(v.Get("data", *appID, "depots", commonDepot, "manifests", *branch, "gid").String(), `"`)

	fmt.Printf("::set-output name=common_manifest_id::%s\n", commonManifestId)
	fmt.Printf("::set-output name=linux_manifest_id::%s\n", linuxManifestId)
	fmt.Printf("::set-output name=build_id::%s\n", buildId)
	fmt.Printf("::set-output name=build_updated_time::%s\n", buildUpdatedTime)
}
