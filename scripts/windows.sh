GOOS=windows
GOARCH=amd64

go build -C ../bin -o npsg ../cmd/cli

cd ..

tar -czvf npsg-windows.tar.gz bin/

mv npsg-windows.tar.gz dist/