
# Download from https://z-kit.cc/
cp_so:
	sudo cp liblcl.so /usr/lib/x86_64-linux-gnu/

run:
	go run main.go

win64-build-on-ubuntu:
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o govcl.exe
