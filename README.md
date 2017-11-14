## Response Checker
[parliament.uk-response-checker](https://github.com/ukparliament/parliament.uk-response-checker) is a Go script designed to retrieve a list of working routes and check how these are responding, reporting back all URLs that have been visited, along with their status code.

### Requirements
Response Checker requires [Go](https://golang.org/doc/install) to be installed.

### How it Works
[parliament.uk-response-checker](https://github.com/ukparliament/parliament.uk-response-checker) executes five main steps:
1. Retrieve a list of parliament routes
2. Extract only the routes live on Beta
3. Replace ID placeholders with real IDs
4. Make a request to each route, recording each route that has been visited, along with their status code
5. Generate a HTML report (using `_report_template.html`) with the above results

### Usage
1. Set up this repository
```kernal
git clone https://github.com/ukparliament/parliament.uk-response-checker
```

2. Set base URL in `request.go` file

3. Amend source of URLs to check against, by modifying the `routeSource` variable in `request.go`

4. Update `resource_map.json` with any new required IDs

5. Build the executable
```kernal
go build .
```

6. This will generate an executable binary in your current directory, `parliament.uk-response-checker`

7. Run the executable binary
```kernal
./parliament.uk-response-checker
```

### Testing
Tests can be run by using the following command:
```kernal
go test -v
```

### Caveats
Currently, this script only supports the latin alphabet and not all Unicode characters.
