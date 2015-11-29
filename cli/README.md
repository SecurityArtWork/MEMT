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


```go
go get
go build memt.go
```


## Usage

```
./memt -h

Usage of memt:
  …
```

