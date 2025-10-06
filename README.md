# Hyperledger Fabric Internship Assignment (Golang)

## 📌 Objective
Blockchain-based asset management system using Hyperledger Fabric.

## ⚙ Technologies
- Hyperledger Fabric v2.4
- Golang
- Docker
- Fabric CA

## 🚀 Setup

### 1. Install Fabric Samples
```bash
curl -sSL https://bit.ly/2ysbOFE | bash -s

### 2. Start test network
cd fabric-samples/test-network
./network.sh up createChannel -c mychannel -ca

### 3. Deploy Chaincode
./network.sh deployCC -ccn asset-transfer -ccp ../asset-transfer-internship/chaincode-go -ccl go

### 4. Invoke and Query
peer chaincode invoke -C mychannel -n asset-transfer -c '{"function":"CreateAsset","Args":["A1","D001","9999999999","1234","5000","ACTIVE","0","INIT","Created via Go chaincode"]}'
peer chaincode query -C mychannel -n asset-transfer -c '{"function":"ReadAsset","Args":["A1"]}'
