# MEMT Analyzer

MEMT Analyzer is a μService intended to be the core binary analyzer and information extractor.

## Building and installation

First we need to clone the project repository:

`git clone https://github.com/SecurityArtWork/MEMT`

After this, a symlink must be created to the `$GOPATH`, if you don't have a operative Go environment you should setup one before going further.

`ln -s /home/securityartwork/MEMT/anal /go/src/github.com/securityartwork/anal`

Now you can build the analyzer μService issuing:


```go
go get
go build malAnalSvc.go
```


## Usage

```
./malAnalSvc -h

Usage of malAnalSvc:
  -bindst string
        Binary final folder. (default "/tmp/bin")
  -dbhost string
        Sets the mongodb address. (default "127.0.0.1")
  -dbport string
        Sets the mongodb port. (default "27017")
  -host string
        Sets the μservice address. (default "127.0.0.1")
  -imgdst string
        Images final folder. (default "/tmp/img")
  -info
        Shows about info.
  -pong
        Sends the result back to the requester.
  -port string
        Sets the μservice port. (default "31337")
```

