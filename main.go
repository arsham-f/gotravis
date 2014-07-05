package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
)

var (
	path   = flag.String("path", "", "Local path of target file")
	target = flag.String("target", "", "Remote path of target file")

	keys = make(map[string]string, 4)
)

func main() {

	flag.Parse()

	required_env_keys := []string{
		"ARTIFACTS_S3_BUCKET",
		"ARTIFACTS_AWS_REGION",
		"ARTIFACTS_AWS_ACCESS_KEY_ID",
		"ARTIFACTS_AWS_SECRET_ACCESS_KEY",
	}

	for _, k := range required_env_keys {
		keys[k] = os.Getenv(k)
		if keys[k] == "" {
			panic("Getting ENV variables", fmt.Sprintf("Missing environment key: %s", k))
		}
	}

	Init(
		keys["ARTIFACTS_AWS_ACCESS_KEY_ID"],
		keys["ARTIFACTS_AWS_SECRET_ACCESS_KEY"],
		keys["ARTIFACTS_S3_BUCKET"],
		keys["ARTIFACTS_AWS_REGION"],
	)

	upload()

}

func upload() {
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		panic("Reading local file", err.Error())
	}

	ext := filepath.Ext(*path)
	ctype := mime.TypeByExtension(ext)
	if err = UploadFile(*target, data, ctype); err != nil {
		panic("Uploading file", err.Error())
	}
}

func panic(label string, msg string) {
	fmt.Printf("PANIC [%s]: %s\n", label, msg)
	os.Exit(-1)
}
