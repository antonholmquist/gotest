
# Export GOPATH
Should be set to project directory to make everything simple

export GOPATH=/Users/antonholmquist/Projects/gotest
export PATH=$PATH:$GOPATH/bin

# Install modules
1. go get github.com/go-martini/martini
2. Remove .git and .ignore and commit it all

# Format code (fixes indentation and stuff)
go fmt

# Run
go run src/github.com/antonholmquist/gotest/app.go 

# go-go-get-go-install-local-packages-and-version-control
http://stackoverflow.com/questions/10130341/go-go-get-go-install-local-packages-and-version-control

http://stackoverflow.com/questions/17780754/automatically-defining-gopath-on-a-per-project-basis