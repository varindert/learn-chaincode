package main

import (
	"errors"
	"fmt"

	

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("mylogger")

type SampleChaincode struct {
}

//custom data models
type CompanyInfo struct {
	Companyname string `json:"companyname"`
	Companycontact  string `json:"companycontact"`
	Companybudget  int `json:"companybudget"`
	AssignedRole  string `json:"assignedRole"`
	CompanyID string `json:"companyid"`
}

type ContractorInfo struct {
	Contractorname string `json:"Contractorname"`	
	Contractorassignedto string `json:"contractorassignedto"`		// assigned to which project
	ContractorHourlyrate  string `json:"contractorHourlyrate"`
	AssignedRole  string `json:"assignedRole"`
	ContractorID string `json:"contractorid"`
}


type ManagerInfo struct {
	Managername string `json:"Contractorname"`	
	Managerassignedto string `json:"managerassignedto"`		// assigned to which project
	AssignedRole  string `json:"assignedRole"`
	ManagerID string `json:"managerid"`
}


func main() {
	err := shim.Start(new(SampleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}




// query function entry point
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "getcompanyinfo" {
		return t.GetCompanyInfo(stub, args)
	}else if function == "read" {
		return t.read(stub, args)
	}
	fmt.Println("Query did not find func: " + function)
	return nil, nil
}

// get companyinfo
func (t *SampleChaincode) GetCompanyInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
//func GetCompanyInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering GetCompanyInfo")

	if len(args) < 1 {
		logger.Error("Invalid number of arguments")
		return nil, errors.New("Missing Company ID")
	}

	var CompanyID = args[0]
	bytes, err := stub.GetState(CompanyID)
	if err != nil {
		logger.Error("Could not fetch company info with id "+CompanyID+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}


// Invoke entry point to invoke a chaincode function
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	 if function == "createcompany" {
		return CreateCompany(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}


func CreateCompany(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering CreateCompany")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for create company creation")
	}

	var companyId = args[0]
	var companyInput = args[1]

	err := stub.PutState(companyId, []byte(companyInput))
	if err != nil {
		logger.Error("Could not save company info to ledger", err)
		return nil, err
	}

	//var customEvent = "{eventType: 'loanApplicationCreation', description:" + loanApplicationId + "' Successfully created'}"
	//err = stub.SetEvent("evtSender", []byte(customEvent))
	
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully saved loan application")
	return nil, nil

}


// read - query function to read key/value pair
func (t *SampleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

// rename this file as chaincode_finished, build it and check into github finished branch, this way I don't have to register and quickly test if a company can be created.
// make post request from postman to test create company and return company