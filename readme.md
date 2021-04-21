
The app is written in golang

To download go please follow the link here = https://golang.org/doc/install

Instructions to run the app

Prerequesite:
Please install the module by 

    go get github.com/urfave/cli
    and $ GO111MODULE=on go get github.com/urfave/cli/v2 
    go build samplecli.go



To get user details 
	./samplecli ID --User russellwhyte

To add users
./samplecli Add --User sud --UserName Sudharsan --Email=sudh@gmail --Country India	

To filter
	./samplecli Filter --Filterval "FirstName eq 'Scott'"