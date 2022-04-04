compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o release/arm/CoveStonks-arm main.go
	GOOS=linux GOARCH=arm64 go build -o release/arm/CoveStonks-arm64 main.go
	GOOS=linux GOARCH=386 go build -o release/linux/CoveStonks main.go
	GOOS=linux GOARCH=amd64 go build -o release/linux/CoveStonks-amd64 main.go
	GOOS=windows GOARCH=386 go build -o release/windows/CoveStonks.exe main.go
	GOOS=windows GOARCH=amd64 go build -o release/windows/CoveoStonks-amd64.exe main.go