package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing Assets
type SmartContract struct {
	contractapi.Contract
}

// Asset structure defines the asset data model
type Asset struct {
	DealerID    string  json:"dealerId"
	MSISDN      string  json:"msisdn"
	MPIN        string  json:"mpin"
	Balance     float64 json:"balance"
	Status      string  json:"status"
	TransAmount float64 json:"transAmount"
	TransType   string  json:"transType"
	Remarks     string  json:"remarks"
}

// CreateAsset adds a new asset to the world state
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, dealerId string, msisdn string, mpin string, balance float64, status string, transAmount float64, transType string, remarks string) error {

	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		DealerID:    dealerId,
		MSISDN:      msisdn,
		MPIN:        mpin,
		Balance:     balance,
		Status:      status,
		TransAmount: transAmount,
		TransType:   transType,
		Remarks:     remarks,
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateBalance updates the balance of an existing asset
func (s *SmartContract) UpdateBalance(ctx contractapi.TransactionContextInterface, id string, newBalance float64) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Balance = newBalance
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

// AssetExists checks if an asset exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, err
	}

	return assetJSON != nil, nil
}

// GetHistory retrieves the full transaction history for an asset
func (s *SmartContract) GetHistory(ctx contractapi.TransactionContextInterface, id string) ([]*Asset, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var history []*Asset
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		if response.Value == nil {
			continue
		}
		var asset Asset
		json.Unmarshal(response.Value, &asset)
		history = append(history, &asset)
	}

	return history, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode: %v", err))
	}

	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("Error starting chaincode: %v", err))
	}
}
