GOOS=darwin
GOARCH=amd4

go build -C ../bin -o npsg ../cmd/cli

cd ..

tar -czvf npsg-macos-intel.tar.gz bin/

mv npsg-macos-intel.tar.gz dist/
