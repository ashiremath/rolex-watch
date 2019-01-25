// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define watch structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
type watch struct {
	Name string `json:"name"`
	Timestamp string `json:"timestamp"`
	Qty  string `json:"qty"`
	Outlet  string `json:"outlet"`
}

/*
 * The Init method *
 called when the Smart Contract "watch-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "watch-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "querywatch" {
		return s.querywatch(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordwatch" {
		return s.recordwatch(APIstub, args)
	} 
	
	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The querywatch method *
Used to view the records of one particular watch
It takes one argument -- the key for the watch in question
 */
func (s *SmartContract) querywatch(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	watchAsBytes, _ := APIstub.GetState(args[0])
	if watchAsBytes == nil {
		return shim.Error("Could not locate watch")
	}
	return shim.Success(watchAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 watch catches)to our network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	watch := []watch{
		watch{Name: "Avi", Qty: "25", Outlet: "Hadapsar", Timestamp: "1504054225"},
	}

	i := 0
	for i < len(watch) {
		fmt.Println("i is ", i)
		watchAsBytes, _ := json.Marshal(watch[i])
		APIstub.PutState(strconv.Itoa(i+1), watchAsBytes)
		fmt.Println("Added", watch[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordwatch method *
Fisherman like Sarah would use to record each of her watch catches. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordwatch(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var watch = watch{ Name: args[1], Qty: args[2], Outlet: args[3], Timestamp: args[4] }

	watchAsBytes, _ := json.Marshal(watch)
	err := APIstub.PutState(args[0], watchAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record watch: %s", args[0]))
	}

	return shim.Success(nil)
}


/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}