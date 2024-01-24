GOOS=linux
GOARCH=amd64

go build -C ../bin -o npsg ../cmd/cli

cd ..

tar -czvf npsg-linux-amd64.tar.gz bin/

mv npsg-linux-amd64.tar.gz dist/