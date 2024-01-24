GOOS=darwin
GOARCH=arm64

go build -C ../bin -o npsg ../cmd/cli

cd ..

tar -czvf npsg-macos-silicon.tar.gz bin/

mv npsg-macos-silicon.tar.gz dist/