cd ~/fabric-samples
mkdir asset-transfer-internship
cd asset-transfer-internship
mkdir chaincode-go
cd chaincode-go

go mod init asset-chaincode
go get github.com/hyperledger/fabric-contract-api-go/contractapi
go mod tidy
