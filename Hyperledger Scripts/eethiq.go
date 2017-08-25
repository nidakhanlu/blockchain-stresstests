package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Initialize the chaincode
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var A, Counter, Projects string            // Entities
	var Aval, numberdonors, numberprojects int // Asset holdings
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	Counter = "counter"
	Projects = "projects"
	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	numberdonors, err = strconv.Atoi("0")
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	numberprojects, err = strconv.Atoi("0")
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	fmt.Printf("Aval = %d\n", Aval)
	err = stub.PutState(Counter, []byte(strconv.Itoa(numberdonors)))
	if err != nil {
		return nil, err
	}
	err = stub.PutState(Projects, []byte(strconv.Itoa(numberprojects)))
	if err != nil {
		return nil, err
	}
	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Registers donors for Eethiq
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}
	if function == "gettime" {
		// Deletes an entity from its state
		return t.gettime(stub, args)
	}

	if function == "buycoins" {
		// Deletes an entity from its state
		return t.buycoins(stub, args)
	}
	if function == "donate" {
		// Deletes an entity from its state
		return t.donate(stub, args)
	}
	if function == "registerproject" {
		// Registers a project
		return t.registerproject(stub, args)
	}

	var Counter string
	var numberdonors int
	Counter = "counter"
	numberdonorsbytes, err := stub.GetState(Counter)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	var Donor string // Entities
	var Aval int     // Asset holdings

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	Donor = args[0]
	numberdonors, _ = strconv.Atoi(string(numberdonorsbytes))
	numberdonors = numberdonors + 1

	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	Aval, err = strconv.Atoi("0")
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	fmt.Printf("Aval = %d\n", Aval)

	// Write the state to the ledger
	err = stub.PutState(Donor, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(Counter, []byte(strconv.Itoa(numberdonors)))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Number of registered Donors = %d\n", numberdonors)
	ts, err2 := stub.GetTxTimestamp()
	if err2 != nil {
		fmt.Printf("Error getting transaction timestamp: %s", err2)
	}
	var tm, ID string

	fmt.Printf("Transaction Time: %v", ts)
	tm = ts.String()
	ID = strconv.Itoa(numberdonors)
	//ID = "id"
	err = stub.PutState(ID, []byte(tm))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Registers a project
func (t *SimpleChaincode) registerproject(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	var P string
	P = args[0]
	var err error
	var Aval int
	Aval, err = strconv.Atoi("0")
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	fmt.Printf("Aval = %d\n", Aval)

	// Write the state to the ledger
	err = stub.PutState(P, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	var Projects string
	var numberprojects int
	Projects = "projects"
	numberprojectsbytes, err := stub.GetState(Projects)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}

	numberprojects, _ = strconv.Atoi(string(numberprojectsbytes))
	numberprojects = numberprojects + 1
	err = stub.PutState(Projects, []byte(strconv.Itoa(numberprojects)))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Number of registered Projects = %d\n", numberprojects)
	return nil, nil

}

// Deletes an entity from state
func (t *SimpleChaincode) gettime(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}
	ts, err2 := stub.GetTxTimestamp()
	if err2 != nil {
		fmt.Printf("Error getting transaction timestamp: %s", err2)
	}
	fmt.Printf("Transaction Time: %v", ts)

	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Lets a donor buy coins
func (t *SimpleChaincode) buycoins(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return nil, errors.New("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Lets a donor donate coins
func (t *SimpleChaincode) donate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return nil, errors.New("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}
	var Counter, id string
	var number int
	Counter = "counter"
	A = args[0]
	numberbytes, err := stub.GetState(Counter)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	number, _ = strconv.Atoi(string(numberbytes))

	if A == "id" {

		id = strconv.Itoa(number)
		tmbytes, err := stub.GetState(id)
		if err != nil {
			return nil, errors.New("Failed to get state")
		}
		return tmbytes, nil
	}
	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
