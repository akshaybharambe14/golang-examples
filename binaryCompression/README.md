# Guide to compress GoLang binaries

## Why we need compression

- The size of binary may not be the issue in today's use cases but, we may need this in cases like running your app on raspberry pi.

- By default, go includes some information in binaries that can be omitted.

- Ultimately, a Go binary can be compressed to 1/4th of its actual size.

To know, how this works, you can read an excellent article [here](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)

The following table summarizes the compression with different combinations of go commands and [UPX](https://upx.github.io/) on a 64bit Windows machine. **However the compression increases startup time for your app**

| Command                                         | Size                       | Compression ratio |
| ----------------------------------------------- | -------------------------- | ----------------- |
| go build                                        | 14.2 MB (14,905,856 bytes) | 100.00            |
| go build -ldflags="-w"                          | 11.3 MB (11,938,304 bytes) | 79.57             |
| go build -ldflags="-s"                          | 10.5 MB (11,089,408 bytes) | 73.94             |
| go build -ldflags="-s -w"                       | 10.5 MB (11,089,408 bytes) | 73.94             |
| go build / upx -1 -k fileName.exe               | 8.23 MB (8,631,808 bytes)  | 57.95             |
| go build / upx -9 -k fileName.exe               | 6.67 MB (7,000,064 bytes)  | 46.97             |
| go build -ldflags="-s" / upx -1 -k fileName.exe | 4.87 MB (5,116,416 bytes)  | 34.29             |
| go build -ldflags="-s" / upx -9 -k fileName.exe | 3.42 MB (3,591,680 bytes)  | 24.08             |
