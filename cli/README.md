# MEMT Command line interface.

MEMT Command Line Interface allows you to interact with the MEMT server API from the system
console.

WARNING: MEMT-CLI sends the artifact object to the server for an
exhaustive analysis, if you don't like to do so, you can disable this
behavior with the -nosend flag.



## operating modes:

* Artifact analysis mode: This mode is triggered with the -file flag, it will
send the artifact to MEMT server and analyze the binary returning all the
available information about this artifact.

* [NYI] Directory watch mode: This mode is triggered with the -watch flag and
setting the watch mode with the -path flag. In this operation mode MEMT-CLI
will watch all the new files dropped in the designed folder and upload them
to MEMT server, returning detailed information about the artifact in case
of being malware.

## Building and installation

First we need to clone the project repository:

`git clone https://github.com/SecurityArtWork/MEMT`

After this, a symlink must be created to the `$GOPATH`, if you don't have a operative Go environment you should setup one before going further.

`ln -s /home/securityartwork/MEMT/cli /go/src/github.com/securityartwork/cli`

Now you can build the categorizer issuing:


```bash
go get
go build memt.go
```


## Usage

```bash
./memt -h

Usage of memt:
  …
```

## Testing the server

```bash
➤ $ curl http://127.0.0.1:8888
Pong

➤ $ curl http://127.0.0.1:8888 -XPOST
{"msg":"Pong"}

➤ $ curl http://127.0.0.1:8888 -XPUT
Method not allowed

➤ $ curl http://127.0.0.1:8888/api/v0/search/1234567890abcdef
{"ecode":302,"msg":"Asset already analysed","goto":"http://127.0.0.1:8888/api/v0/malware/info/1234567890abcdef"}

➤ $ curl http://127.0.0.1:8888/api/v0/search/1234567890abcdefa
{"ecode":200,"msg":"Analysis has been launch in background","goto":"http://127.0.0.1:8888/api/v0/malware/info/1234567890abcdefa/UUID-XXXX-YYYYYYYYY"}

➤ $ curl http://127.0.0.1:8888/api/v0/search/1234567890abcdef -XPOST
404 page not found

➤ $ curl http://127.0.0.1:8888/api/v0/malware/info/1234567890abcdef
{"ecode":302,"msg":"Asset already analysed","data":{"ssdeep":"1234567890","md5":"1234567890","sha1":"1234567890","sha256":"1234567890","sha512":"1234567890","format":"pe","symbols":["a","b"],"imports":["a","b"],"sections":["a","b"],"arch":"amd64","strain":"","mutations":["0987654321","5647382910","4536789013"],"siblings":[""]}}

➤ $ curl http://127.0.0.1:8888/api/v0/malware/info/1234567890abcdefs
{"ecode":404,"msg":"This element does not exist","goto":""}
```
