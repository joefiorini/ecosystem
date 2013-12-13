# If we're building on Windows, specify an extension
EXTENSION=""
if [ "$(go env GOOS)" = "windows" ]; then
    EXTENSION=".exe"
fi

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

echo "--> Installing dependencies to speed up builds..."
go get ./...

# Build!
echo "--> Building..."
go build \
    -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" \
    -v \
    -o bin/eco${EXTENSION}
cp bin/eco${EXTENSION} $GOPATH/bin
