/*
### API endpoints ###

* Search by hash
    => GET /api/v0/search/<hash>
    <= "http://xx.xx.xx.xx/api/v0/malware/info/<sha256:hash>"

    => GET /api/v0/search/<hash>/<type>
    <= "http://xx.xx.xx.xx/api/v0/malware/info/<sha256:hash>"


* Get info by hash
    => GET /api/v0/malware/info/<sha256:hash>
    <= Response:
    {
        "ecode": 302,
        "msg": "Asset already analysed",
        "data": {
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
    }

    {
        "ecode": 404
        "info": "This element does not exist"
        "state": ""
    }

    => GET /api/v0/malware/info/<sha256:hash>/<uuid:uuid>
    <= Response:
    {
        "ecode": 200
        "info": ""
        "state": ""
    }

    {
        "ecode": 404
        "info": "This element does not exist"
        "state": ""
    }



* Get info by artifact upload
    => POST (Multipart) /api/v0/malware/submit
    <= Response:
        {
            "ecode": 302,
            "msg": "Asset already analysed",
            "goto": url_for("MalwareView:info", hash=filename, type=type)
        }

        {
            "ecode": 200,
            "msg": "Analysis has been launch in background",
            "goto": url_for("MalwareView:info", hash=filename, type=type),
            "task_id": task_id.id
        }

*/

package api

import (
	"fmt"
	"net/http"
)

const (
	// API = "http://malhive.io"
	API = "http://127.0.0.1:8888"
)

func SendHash(hash string) {
	endpoint := "/api/v0/search/"
	url := API + endpoint + hash

	// Get response
	response, _, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(contents))
}

// func SendArtifactData(hash string) {
// 	endpoint := "malware/"

// 	var jsonStr = []byte(`{"sha256":"` + hash + `"}`)
// 	req, err := http.NewRequest(
// 		"POST",
// 		API+endpoint,
// 		bytes.NewBuffer(jsonStr),
// 	)

// 	req.Header.Set("User-Agent", "memt-cli")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("response Status:", resp.Status)
// 	fmt.Println("response Headers:", resp.Header)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println("response Body:", string(body))
// }
