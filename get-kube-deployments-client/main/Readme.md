#Kube-app build for Platform Developer Test

Before proceeding with building the webapp, kindly deploy the kubernetes deployment using the following command.  The services file can be found oustide the main folder. please go a level back and run :
`kubectl apply -f ./services.yaml`

Confirm if the deployment is successful and if the pods are running. Please give it a couple of minutes for the pods to initialize
`kubectl -n default get pods`

Upon sucessful deployment you will see the pods running.

Move into the main folder to build the file. The go.mod file has the dependenices. simply run the following to installing them

`go mod download`

Once the dependencies are completely download, we can proceed with building the app

From the main folder simply run `go run .` or `go build` and `./main`

You can also go build -o app and run the binary. Both serves the same purpose. 

Once started please connect to localhost:8080 to use the following URL

localhost:8080/services

localhost:8080/services/alpha or beta or gamma

For reference regarding the responses you will find two png files in the main folder. kindly check if you get the results as seen. 



