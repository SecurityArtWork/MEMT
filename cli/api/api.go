/*
### API endpoints ###

* Search by hash
    => GET /api/v0/search/<hash>
    <= Response:
    {
        "ecode": 302,
        "msg": "Asset already analysed",
        "goto": "http://xx.xx.xx.xx/api/v0/malware/info/<sha256:hash>"
    }

    {
        "ecode": 404
        "info": "This element does not exist"
        "state": ""
    }


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
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// MEMT = "http://malhive.io"
	MEMT = "http://127.0.0.1:8888"
	API  = "/api/v0/"
)

type ArtRes struct {
	Ssdeep    string   `json:"ssdeep"`
	Md5       string   `json:"md5"`
	Sha1      string   `json:"sha1"`
	Sha256    string   `json:"sha256"`
	Sha512    string   `json:"sha512"`
	Format    string   `json:"format"`
	Symbols   []string `json:"symbols"`
	Imports   []string `json:"imports"`
	Sections  []string `json:"sections"`
	Arch      string   `json:"arch"`
	Strain    string   `json:"strain"`
	Mutations []string `json:"mutations"`
	Siblings  []string `json:"siblings"`
}

type DataRes struct {
	Ecode int    `json:"ecode"`
	Msg   string `json:"msg"`
	Data  ArtRes `json:"data"`
}

type SearchRes struct {
	Ecode int    `json:"ecode"`
	Msg   string `json:"msg"`
	Goto  string `json:"goto"`
}

func SearchHash(hash string) (SearchRes, error) {
	var sr SearchRes
	endpoint := "search/"
	url := MEMT + API + endpoint + hash

	// Get response
	response, err := http.Get(url)
	if err != nil {
		return sr, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return sr, err
	}

	// Unmarshal json, and, if error return unprocessable entity
	if err := json.Unmarshal(body, &sr); err != nil {
		return sr, err
	}

	return sr, nil
}

func GetInfo(url string) (DataRes, error) {
	var dr DataRes

	// Get response
	response, err := http.Get(url)
	if err != nil {
		return dr, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dr, err
	}

	// Unmarshal json, and, if error return unprocessable entity
	if err := json.Unmarshal(body, &dr); err != nil {
		return dr, err
	}

	return dr, nil
}
