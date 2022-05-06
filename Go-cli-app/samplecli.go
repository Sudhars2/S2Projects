package main

import (
  
  "log"
  "os"
  "fmt"
  "io/ioutil"
  "net/http"
	"strings"
  "github.com/urfave/cli"
  "encoding/json"
  //"bytes"
)

func main() {
	client := &http.Client {
		}
	app := cli.NewApp()
	app.Name = "API Lookup"
	app.Usage = "To query, update and filter user requested details"
	//var User string
	myFlags := []cli.Flag{
		&cli.StringFlag{
		Name:        "User",
		Value:       "Sudharsan_001",
        Usage:       "Enter the UserID to be searched",
		},
		&cli.StringFlag{
			Name:        "UserName",
			Value:       "SudharsanS",
			Usage:       "Enter the username",
			},
		&cli.StringFlag{
				Name:        "Email",
				Value:       "SudharsanS@nomail.com",
				Usage:       "Enter the user email",
				},
        &cli.StringFlag{
					Name:        "Country",
					Value:       "India",
					Usage:       "Enter the user Country",
					},
		&cli.StringFlag{
						Name:        "Filterval",
						Value:       "nil",
						Usage:       "Enter the exact filter details",
						},
			
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ID",
			Usage: "Search for the given user ID",
			Flags: myFlags,
			
			Action: func(c *cli.Context) error {
				//if c.NArg() > 0 {
					//fmt.Println(c.Args())
					userid := c.String("User")
					fmt.Println(userid)
					url := "https://services.odata.org/TripPinRESTierService/People('"+userid+"')"
					req, err := http.NewRequest("GET", url, nil)
					if err != nil {
						fmt.Println(err)
					}
					res, err := client.Do(req)
					if err != nil {
						fmt.Println(err.Error())
					}
					body, err := ioutil.ReadAll(res.Body)

					fmt.Println(string(body))
					defer res.Body.Close()
					
			
				return nil
			},
		},
		{
			Name:  "Add",
			Usage: "Add details for the given User ID ",
			Flags: myFlags,
			
			Action: func(c *cli.Context) error {
				    m := make(map[string]string)

					m["UserName"] = c.String("User")
					m["FirstName"] = c.String("UserName")
					m["Emails"] = c.String("Email")
					m["CountryRegion"] = c.String("Country")
					
					jsonString, err := json.Marshal(m)
					if err != nil {
						panic(err)
					}
					//fmt.Println(userid)
					url := "https://services.odata.org/TripPinRESTierService/People"
					//payload := strings.NewReader("{\n    \"UserName\":\"SudharsanS\",\n    \"FirstName\":\"S\",\n    \"LastName\":\"Suds\",\n    \"Emails\":[\n        \"sudharsan@example.com\"\n    ],\n    \"AddressInfo\": [\n    {\n      \"Address\": \"187 Suffolk Ln.\",\n      \"City\": {\n        \"Name\": \"Arbat\",\n        \"CountryRegion\": \"Belarus\",\n        \"Region\": \"ID\"\n      }\n    }\n    ]\n}")
					fmt.Println(string(jsonString))
					req, err := http.NewRequest("POST", url,strings.NewReader(string(jsonString)))
					if err != nil {
						fmt.Println(err)
					}
					req.Header.Add("Content-Type", "application/json")
					res, err := client.Do(req)
					if err != nil {
						fmt.Println(err.Error())
					}
					//body, err := ioutil.ReadAll(res.Body)

					fmt.Println(res.Status)
					defer res.Body.Close()
					
			
				return nil
			},
		},

		{
			Name:  "Filter",
			Usage: "Filter details based on given conditions ",
			Flags: myFlags,
			
			Action: func(c *cli.Context) error {
				    
					filter := c.String("Filterval")
					fmt.Println(filter)
					filterresult := strings.Replace(filter, " ", "+", -1)
					url := "https://services.odata.org/TripPinRESTierService/People?$filter="+filterresult
					req, err := http.NewRequest("GET", url, nil)
					if err != nil {
						fmt.Println(err)
					}
					res, err := client.Do(req)
					if err != nil {
						fmt.Println(err.Error())
					}
					body, err := ioutil.ReadAll(res.Body)

					fmt.Println(string(body))
					defer res.Body.Close()
					
			
				return nil
			},
		},
	}

	
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


/*

Instructions to run the app

Please install the module by 
go get github.com/urfave/cli
and $ GO111MODULE=on go get github.com/urfave/cli/v2 
go build samplecli.go

To get user details 
	./samplecli ID --User russellwhyte

To add
./samplecli Add --User sud --UserName Sudharsan --Email=sudh@gmail --Country India	

To filter
	./samplecli Filter --Filterval "FirstName eq 'Scott'"