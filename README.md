# goRestAPI

## Instructions

- CD into the goRestAPI directory and run "go run *.go" and it should run the GoServer.

- No need to worry about a db because I create the structs as tables when the app builds and
its cloud hosted through elephantsql. ENJOY!

# BackStory
- Your goal is to create a service, preferably written in Go (https://golang.org), that will accept new
true-to-size data through the Hypertext Transfer Protocol (HTTP) and store this data in a relational
database, preferably Postgres. Additionally, this service should be able to return a shoe’s
TrueToSizeCalculation , defined below, through an HTTP request. True to size entries range between
1 and 5, inclusive, where 1: really small, 2: small, 3: true to size, 4: big and 5: really big.
