peer chaincode invoke -o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com --tls \
--cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" \
-C mychannel -n asset-transfer \
--peerAddresses localhost:7051 \
--tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" \
-c '{"function":"CreateAsset","Args":["A1","D001","9999999999","1234","5000","ACTIVE","0","INIT","Created via Go chaincode"]}'

peer chaincode query -C mychannel -n asset-transfer -c '{"function":"ReadAsset","Args":["A1"]}'

peer chaincode invoke -C mychannel -n asset-transfer -c '{"function":"UpdateBalance","Args":["A1","6000"]}'

peer chaincode query -C mychannel -n asset-transfer -c '{"function":"GetAllAssets","Args":[]}'

peer chaincode query -C mychannel -n asset-transfer -c '{"function":"GetHistory","Args":["A1"]}'
