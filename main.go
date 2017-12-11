package main

import (
  "bufio"
  "os"
	"fmt"
	"log"
	"math/big"
  "strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
  "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ContractData struct {
  sim *backends.SimulatedBackend
  addr common.Address
  contract *LockboxApp
  auth *bind.TransactOpts
}

var err error
var userMap map[string]*ContractData

func deployLockbox(uname string) {
  cd := new(ContractData)
  key, _ := crypto.GenerateKey()
  cd.auth = bind.NewKeyedTransactor(key)

  alloc := make(core.GenesisAlloc)
  alloc[cd.auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
  cd.sim = backends.NewSimulatedBackend(alloc)
  // Deploy LockboxApp contract
  cd.addr, _, cd.contract, err = DeployLockboxApp(cd.auth, cd.sim)
  if err != nil {
    log.Fatalf("### Error Deploying Contract: %v", err)
  }
  cd.sim.Commit()
  userMap[uname] = cd
}

func testlockbox() {
  uname := "randomuser"
  deployLockbox(uname)

  cd := userMap[uname]

  fmt.Println("# Mining...")
  // Simulate mining
  cd.sim.Commit()
  fmt.Printf("# Contract Deployed to %s\n", cd.addr.String())

  fmt.Printf("# Sender: %v, \n", cd.auth.From.String())

  lname := "cs244b"
  // Register a lockbox 
  fmt.Println("\n# Creating a lockbox for user: ", lname)
  _, err = cd.contract.CreateLockbox(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, lname)
  if err != nil {
    fmt.Printf("### Address %s already owns a lockbox\n", cd.auth.From.String())
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Println("# Successfully created Lockbox") 
  fmt.Printf("# Lockbox address: %v, Lockbox name: %s\n", cd.auth.From.String(), lname)

  fmt.Println("\n# Getting lockbox name for input user account")
  uname, err := cd.contract.GetLockboxName(nil, cd.auth.From)
  if err != nil {
    log.Fatalf("### No lockbox account found: %v\n", err)
  }
  fmt.Printf("# Lockbox Account belongs to: %s\n", uname)

  fmt.Printf("\n# Instantiating contract at address: %s\n", cd.auth.From.String())
  instContract, err := NewLockboxApp(cd.addr, cd.sim)
  if err != nil {
    log.Fatalf("### Could not instantiate contract: %v", err)
  }
  fmt.Printf("# Contract instatiated at address: %s\n", cd.auth.From.String())

  fmt.Printf("\n# Adding personal info from address: %s\n", cd.auth.From.String())
  ssn := "jun23433l"
  license := "9cas170"
  var dob int64 = 12111990
  _, err = instContract.AddPersonalInfo(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, ssn, license, big.NewInt(dob));
  if err != nil {
    log.Fatalf("### Error adding Personal Info: %v", err)
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Printf("# Personal Info added for address: %s\n", cd.auth.From.String())

  fmt.Println("\n# Getting SSN of input user account")
  ssn, err = instContract.GetSSN(nil, cd.auth.From)
  if err != nil {
    log.Fatalf("### Error getting SSN information in lockbox account: %v\n", err)
  }
  fmt.Printf("# SSN: %s\n", ssn)

  fmt.Println("\n# Getting License Number of input user account")
  license, err = instContract.GetLicenseNo(nil, cd.auth.From)
  if err != nil {
    log.Fatalf("### Error getting License Number information in lockbox account: %v\n", err)
  }
  fmt.Printf("# License Number: %s\n", license)

  fmt.Println("\n# Getting DOB of input user account")
  sdob, err := instContract.GetDob(nil, cd.auth.From)
  if err != nil {
    log.Fatalf("### Error getting DOB information in lockbox account: %v\n", err)
  }
  fmt.Println("# DOB: ", sdob)


  fmt.Printf("\n# Adding Credit Card details from address: %s\n", cd.auth.From.String())
  cname := "CHASE SAPPHIRE"
  var cardno uint64 = 9273131731489870343
  var edate int64 = 12112019 
  var cvv int64 = 375
  _, err = instContract.AddCardDetails(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, cname, cardno, big.NewInt(edate), big.NewInt(cvv));
  if err != nil {
    log.Fatalf("### Error adding Credit Card details: %v", err)
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Printf("# Credit Card details added for address: %s\n", cd.auth.From.String())

  fmt.Println("\n# Getting Card details of input user account")
  cdetails, err := instContract.GetCardDetails(nil, cd.auth.From, cname)
  if err != nil {
    log.Fatalf("### Card details not found in lockbox account: %v\n", err)
  }
  fmt.Println("# Card Number: ", cdetails.Cardno, " Expiry Date: ", cdetails.Edate, " CVV: ", cdetails.Cvv)

  fmt.Printf("\n# Adding login details from address: %s\n", cd.auth.From.String())
  service := "GMAIL"
  username := "ethblockchain"
  password:= "cs244bscpd"
  _, err = instContract.AddLoginCredentials(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, service, username, password); 
  if err != nil {
    log.Fatalf("### Error adding login details: %v", err)
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Printf("# Login details added for address: %s\n", cd.auth.From.String())

  fmt.Println("\n# Getting login details of input user account")
  ldetails, err := instContract.GetLoginCredentials(nil, cd.auth.From, service)
  if err != nil {
    log.Fatalf("### Login details not found in lockbox account: %v\n", err)
  }
  fmt.Println("# Service: ", ldetails.Service, " Username: ", ldetails.Uname, " Password: ", ldetails.Passwd)

}

func printLockboxOptions() {
  fmt.Println("\n")
  fmt.Println("=================================================")
  fmt.Println("============ Decentralized Lockbox ==============")
  fmt.Println("=================================================")
  fmt.Println("# Press 1 to create a lockbox")
  fmt.Println("# Press 2 to add personal info")
  fmt.Println("# Press 3 to add login credentials")
  fmt.Println("# Press 4 to add credit card details")
  fmt.Println("# Press 5 to run lockbox test utility")
  fmt.Println("# Press 6 to get login credentials")
  fmt.Println("# Press 7 to get personal information")
  fmt.Println("# Press 8 to get credit card details")
  fmt.Println("# Press 0 to exit")
}

// function creates a new lockbox for user
func createLockbox() {
  fmt.Println("# Enter Username for the lockbox")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  lname := scanner.Text()

  deployLockbox(lname)
  cd := userMap[lname]

  // Create lockbox
  fmt.Println("\n# Creating lockbox for user: ", lname)
  _, err := cd.contract.CreateLockbox(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, lname)
  if err != nil {
    fmt.Printf("### Error creating lockbox. Address %s already owns a lockbox\n", cd.auth.From.String())
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Println("# Transaction Complete. Lockbox successfully created") 
  fmt.Printf("# Lockbox address: %v, Lockbox name: %s\n", cd.auth.From.String(), lname)
}

// function adds personal info in the user lockbox
func addPersonalInfo() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter Username:")
  scanner.Scan()
  uname := scanner.Text()
  cd := userMap[uname]
  fmt.Println("# Enter Personal Details:")
  fmt.Println("# Enter SSN:")
  scanner.Scan()
  ssn := scanner.Text()
  fmt.Println("# Enter license number:")
  scanner.Scan()
  license := scanner.Text()
  fmt.Println("# Enter DOB:")
  scanner.Scan()
  dob, serr := strconv.ParseInt(scanner.Text(), 10, 64)
  if serr != nil {
    log.Fatalf("### Invalid Date of Birth %v", serr)
  }

  _, err = cd.contract.AddPersonalInfo(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, ssn, license, big.NewInt(dob));
  if err != nil {
    log.Fatalf("### Error adding Personal Info: %v", err)
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Println("# Transaction Completed")
  fmt.Printf("# Personal Info added for address: %s\n", cd.auth.From.String())
}

// function adds login details in the user lockbox
func addLoginCredentials() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter Username:")
  scanner.Scan()
  uname := scanner.Text()
  cd := userMap[uname]
  fmt.Println("# Enter Login Credentials:")
  fmt.Println("# Enter Account:")
  scanner.Scan()
  service := scanner.Text()
  fmt.Println("# Enter Username:")
  scanner.Scan()
  username := scanner.Text()
  fmt.Println("# Enter Password:")
  scanner.Scan()
  password := scanner.Text()

  _, err = cd.contract.AddLoginCredentials(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, service, username, password); 
  if err != nil {
    log.Fatalf("### Could not add login details: %v", err)
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Println("# Transaction Completed")
  fmt.Printf("# Login details added for address: %s\n", cd.auth.From.String())

}

// function adds credit card details in the user lockbox
func addCreditCardDetails() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter Username:")
  scanner.Scan()
  uname := scanner.Text()
  cd := userMap[uname]
  fmt.Println("# Enter Card Details:")
  fmt.Println("# Enter Card Name:")
  scanner.Scan()
  cname := scanner.Text()
  fmt.Println("# Enter Card Number:")
  scanner.Scan()
  cardno, serr := strconv.ParseUint(scanner.Text(), 10, 64)
  if serr != nil {
    log.Fatalf("### Invalid Card Number %v", serr)
  }
  fmt.Println("# Enter Expiry Date:")
  scanner.Scan()
  edate, serr := strconv.ParseInt(scanner.Text(), 10, 64)
  if serr != nil {
    log.Fatalf("### Invalid Expiry Date %v", serr)
  }
  fmt.Println("# Enter CVV:")
  scanner.Scan()
  cvv, serr := strconv.ParseInt(scanner.Text(), 10, 64)
  if serr != nil {
    log.Fatalf("### Invalid CVV Number %v", serr)
  }

  _, err = cd.contract.AddCardDetails(&bind.TransactOpts{
    From:     cd.auth.From,
    Signer:   cd.auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, cname, cardno, big.NewInt(edate), big.NewInt(cvv));
  if err != nil {
    log.Fatalf("### Error adding card details: %v", err)
  }
  fmt.Println("# Mining...")
  cd.sim.Commit()
  fmt.Println("# Transaction Completed")
  fmt.Printf("# Card Details added for address: %s\n", cd.auth.From.String())
}

// function gets personal info from the user lockbox
func getPersonalInfo() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter Username:")
  scanner.Scan()
  uname := scanner.Text()
  cd, ok := userMap[uname]
  if !ok {
    fmt.Println("\n### Lockbox account not found for user: ", uname)
    return
  }

  fmt.Println("\n# Getting SSN of input user account")
  ssn, err := cd.contract.GetSSN(nil, cd.auth.From)
  if err != nil {
    log.Fatalf("### No ssn information found in lockbox account: %v\n", err)
  }
  if ssn == "0" {
    fmt.Println("# No SSN information found in lockbox account");
  }
  fmt.Printf("# SSN: %s\n", ssn)

  fmt.Println("\n# Getting License Number of input user account")
  license, err := cd.contract.GetLicenseNo(nil, cd.auth.From)
  if err != nil {
    fmt.Println("### Error getting license number information in lockbox account")
  }
  if license == "0" {
    fmt.Println("# No License Number information found in lockbox account");
  }
  fmt.Printf("# License No: %s\n", license)

  fmt.Println("\n# Getting DOB of input user account")
  sdob, err := cd.contract.GetDob(nil, cd.auth.From)
  if err != nil {
    log.Fatalf("### Error getting DOB information in lockbox account: %v\n", err)
  }
  fmt.Println("# DOB: ", sdob)
}

// function gets login details from the user lockbox
func getLoginDetails() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter Username:")
  scanner.Scan()
  uname := scanner.Text()
  cd, ok := userMap[uname]
  if !ok {
    fmt.Println("\n### Lockbox account not found for user: ", uname)
    return
  }

  fmt.Println("\n# Getting login details of input user account")
  fmt.Println("\n# Enter Account")
  scanner.Scan()
  service := scanner.Text()
  ldetails, err := cd.contract.GetLoginCredentials(nil, cd.auth.From, service)
  if err != nil {
    fmt.Println("### Error getting Login Details in lockbox")
  }
  fmt.Println("# Service: ", ldetails.Service, " Username: ", ldetails.Uname, " Password: ", ldetails.Passwd)
}

// function gets credit card details from the user lockbox
func getCardDetails() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter Username:")
  scanner.Scan()
  uname := scanner.Text()
  cd, ok := userMap[uname]
  if !ok {
    fmt.Println("\n### Lockbox account not found for user: ", uname)
    return
  }

  fmt.Println("\n# Getting Card details of input user account")
  fmt.Println("\n# Enter account name")
  scanner.Scan()
  cname := scanner.Text()
  cdetails, err := cd.contract.GetCardDetails(nil, cd.auth.From, cname)
  if err != nil {
    log.Fatalf("### Card details not found in lockbox account: %v\n", err)
  }
  fmt.Println("# Card Number: ", cdetails.Cardno, " Expiry Date: ", cdetails.Edate, " CVV: ", cdetails.Cvv)
}

func main() {
  printLockboxOptions()
  userMap = make(map[string]*ContractData)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    if line == "0" {
      os.Exit(0)
    }
    if line == "1" {
      createLockbox()
      printLockboxOptions()
    }
    if line == "2" {
      addPersonalInfo()
      printLockboxOptions()
    }
    if line == "3" {
      addLoginCredentials()
      printLockboxOptions()
    }
    if line == "4" {
      addCreditCardDetails()
      printLockboxOptions()
    }
    if line == "5" {
      testlockbox()
      printLockboxOptions()
    }
    if line == "6" {
      getLoginDetails()
      printLockboxOptions()
    }
    if line == "7" {
      getPersonalInfo()
      printLockboxOptions()
    }
    if line == "8" {
      getCardDetails()
      printLockboxOptions()
    }
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Error reading standard input: ", err)
  }
}
