GOOS=linux
GOARCH=arm64

go build -C ../bin -o npsg ../cmd/cli

cd ..

tar -czvf npsg-linux-arm64.tar.gz bin/

mv npsg-linux-arm64.tar.gz dist/