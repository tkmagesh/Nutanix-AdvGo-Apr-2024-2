# Go Intermediate

## Magesh Kuppan
- https://linkedin.com/in/tkmagesh

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (15 mins)
- Lunch Break   : 12:30 PM (45 mins)
- Tea Break     : 2:30 PM (15 mins)
- Wind up       : 4:00 PM

## Repo
- https://github.com/tkmagesh/nutanix-advgo-apr-2024-2

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code
- Go extension for VSCode

## Methodology
- No powerpoints
- Code & Discuss

## Modules & Packages
### Module
- Any code that need to versioned and deployed together
- Folder with go.mod file
- go.mod (manifest file)
    - Name
        - advisable to have the complete source control path where the code is maintained
    - go runtime version
    - dependencies
- Create a module
    > go mod init [module_name]
- Create a build
    > go build .

    > go build -o [binary_name] .
- Execute a module
    > go run .
- To use a 3rd party module
    > go get [module_name]
    - downloaded to the module cache ($GOPATH$/pkg/mod/...)
- To sync the dependency references in the go.mod file
    > go mod tidy
- To compile & use a 3rd party module as a CLI
    > go install [module_name]
- To download the dependencies documented in the go.mod file
    > go mod download
- To localize the dependencies
    > go mod vendor
- Other useful commands
    > go mod graph

    > go mod why [module_name]
- Reference
    > https://go.dev/ref/mod

### Package
- Internal organization of code in a module
- Typically a folder
- Code can be scattered across multiple files in a package folder, but they all will be indentified by package name (NOT the filename)
- Public entities should be named starting with an uppercase

## Cross Compilation
- Environment Variables - [GOOS, GOARCH]
- To get the list of supported platforms
> go tool dist list
- To cross compile
> GOOS=[target_os] GOARCH=[target_arch] go build -o [binary_name] .