/*
### API endpoints ###

* Search by hash
    GET /api/v0/search/<hash>
    GET /api/v0/search/<hash>/<type>

* Get info by hash
    GET /api/v0/malware/info/<sha256:hash>

== Response ==
{
    "ssdeep": ""
    "md5": ""
    "sha1": ""
    "sha256": ""
    "sha512": ""
    "format": ""
    "symbols": [""]
    "imports": [""]
    "sections": [""]
    "arch": ""
    "strain": ""
    "mutations": [""]
    "siblings": [""]
}

* Get info by artifact upload
    /api/v0/malware/submit (Multipart)

== Response
{
    "res": {
        "ecode": 302,
        "msg": "Asset already analysed",
        "goto": url_for("MalwareView:info", hash=filename, type=type)
    }
}

{
    "res": {
        "ecode": 200,
        "msg": "Analysis hsa been lunch in background",
        "goto": url_for("MalwareView:info", hash=filename, type=type),
        "task_id": task_id.id
    }
}


*/
package api

import (
	"fmt"
	"net/http"
)

const (
	API = "http://malhive.io/api/v0/"
)

func SendArtifactData(hash string) {
	endpoint := "malware/"

	var jsonStr = []byte(`{"sha256":"` + hash + `"}`)
	req, err := http.NewRequest(
		"POST",
		API+endpoint,
		bytes.NewBuffer(jsonStr),
	)

	req.Header.Set("User-Agent", "memt-cli")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
