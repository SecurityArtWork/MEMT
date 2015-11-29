# MEMT Analyzer

Tool for generating the initial dataset from your samples. This tool classifies, all your samples and outputs a nice JSON ready to be imported into MongoDB, it also generates all the images regarding your samples.

## Building and installation

First we need to clone the project repository:

`git clone https://github.com/SecurityArtWork/MEMT`

After this, a symlink must be created to the `$GOPATH`, if you don't have a operative Go environment you should setup one before going further.

`ln -s /home/securityartwork/MEMT/cat /go/src/github.com/securityartwork/cat`

Now you can build the categorizer issuing:


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

