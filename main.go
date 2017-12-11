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

var sim *backends.SimulatedBackend
var addr common.Address
var contract *LockboxApp
var auth *bind.TransactOpts
var err error

func deployLockbox() {
  key, _ := crypto.GenerateKey()
  auth = bind.NewKeyedTransactor(key)

  alloc := make(core.GenesisAlloc)
  alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
  sim = backends.NewSimulatedBackend(alloc)
  // Deploy LockboxApp contract
  addr, _, contract, err = DeployLockboxApp(auth, sim)
  if err != nil {
    log.Fatalf("### Could not deploy contract: %v", err)
  }
  sim.Commit()
}

func testlockbox() {
  deployLockbox()

  fmt.Println("# Mining...")
  // Simulate mining
  sim.Commit()
  fmt.Printf("# Contract deployed to %s\n", addr.String())

  fmt.Printf("# Sender: %v, \n", auth.From.String())

  lname := "avneeshp"
  // Register a lockbox 
  fmt.Println("\n# Registering a lockbox...")
  _, err = contract.CreateLockbox(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, lname)
  if err != nil {
    fmt.Printf("### Address %s already owns a lockbox\n", auth.From.String())
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Println("# Successfully created Lockbox") 
  fmt.Printf("# Lockbox address: %v, Lockbox name: %s\n", auth.From.String(), lname)

  fmt.Println("\n# Getting lockbox name for input user account")
  uname, err := contract.GetLockboxName(nil, auth.From)
  if err != nil {
    log.Fatalf("### No lockbox account found: %v\n", err)
  }
  fmt.Printf("# Lockbox account belongs to: %s\n", uname)

  fmt.Printf("\n# Instantiating contract at address: %s\n", auth.From.String())
  instContract, err := NewLockboxApp(addr, sim)
  if err != nil {
    log.Fatalf("### Could not instantiate contract: %v", err)
  }
  fmt.Printf("# Contract instatiated at address: %s\n", auth.From.String())

  fmt.Printf("\n# Adding personal info from address: %s\n", auth.From.String())
  ssn := "jun23433l"
  license := "9cas170"
  var dob int64 = 12111990
  _, err = instContract.AddPersonalInfo(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, ssn, license, big.NewInt(dob));
  if err != nil {
    log.Fatalf("### Could not add personal info: %v", err)
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Printf("# Personal info added for address: %s\n", auth.From.String())

  fmt.Println("\n# Getting SSN of input user account")
  ssn, err = instContract.GetSSN(nil, auth.From)
  if err != nil {
    log.Fatalf("### No ssn information found in lockbox account: %v\n", err)
  }
  fmt.Printf("# SSN: %s\n", ssn)

  fmt.Println("\n# Getting License Number of input user account")
  license, err = instContract.GetLicenseNo(nil, auth.From)
  if err != nil {
    log.Fatalf("### No ssn information found in lockbox account: %v\n", err)
  }
  fmt.Printf("# License No: %s\n", license)

  fmt.Println("\n# Getting DOB of input user account")
  sdob, err := instContract.GetDob(nil, auth.From)
  if err != nil {
    log.Fatalf("### No dob information found in lockbox account: %v\n", err)
  }
  fmt.Println("# DOB: ", sdob)


  fmt.Printf("\n# Adding credit card details from address: %s\n", auth.From.String())
  cname := "CHASE SAPPHIRE"
  var cardno uint64 = 9273131731489870343
  var edate int64 = 12112019 
  var cvv int64 = 375
  _, err = instContract.AddCardDetails(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, cname, cardno, big.NewInt(edate), big.NewInt(cvv));
  if err != nil {
    log.Fatalf("### Could not add card details: %v", err)
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Printf("# Card details added for address: %s\n", auth.From.String())

  fmt.Println("\n# Getting Card details of input user account")
  cdetails, err := instContract.GetCardDetails(nil, auth.From, cname)
  if err != nil {
    log.Fatalf("### Card details not found in lockbox account: %v\n", err)
  }
  fmt.Println("# Card Number: ", cdetails.Cardno, " Expiry Date: ", cdetails.Edate, " CVV: ", cdetails.Cvv)

  fmt.Printf("\n# Adding login details from address: %s\n", auth.From.String())
  service := "GMAIL"
  username := "ethblockchain"
  password:= "cs244bscpd"
  _, err = instContract.AddLoginCredentials(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, service, username, password); 
  if err != nil {
    log.Fatalf("### Could not add login details: %v", err)
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Printf("# Login details added for address: %s\n", auth.From.String())

  fmt.Println("\n# Getting login details of input user account")
  ldetails, err := instContract.GetLoginCredentials(nil, auth.From, service)
  if err != nil {
    log.Fatalf("### Login details not found in lockbox account: %v\n", err)
  }
  fmt.Println("# Service: ", ldetails.Service, " Username: ", ldetails.Uname, " Password: ", ldetails.Passwd)

}

func printLockboxOptions() {
  fmt.Println("\n")
  fmt.Println("============ Decentralized Lockbox ==============")
  fmt.Println("# Press 1 to create a lockbox")
  fmt.Println("# Press 2 to add personal info")
  fmt.Println("# Press 3 to add login credentials")
  fmt.Println("# Press 4 to add credit card details")
  fmt.Println("# Press 5 to run lockbox test utility")
  fmt.Println("# Press 6 to get login details for an account")
  fmt.Println("# Press 7 to get personal information")
  fmt.Println("# Press 8 to get card details")
  fmt.Println("# Press 0 to exit")
}

// function creates a new lockbox for user
func createLockbox() {
  fmt.Println("# Enter username for the lockbox")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  lname := scanner.Text()

  // Create lockbox
  fmt.Println("\n# Creating lockbox for user: ", lname)
  _, err := contract.CreateLockbox(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, lname)
  if err != nil {
    fmt.Printf("### Address %s already owns a lockbox\n", auth.From.String())
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Println("# Successfully created Lockbox") 
  fmt.Printf("# Lockbox address: %v, Lockbox name: %s\n", auth.From.String(), lname)
}

// function adds personal info in the user lockbox
func addPersonalInfo() {
  fmt.Println("# Enter personal details:")
  scanner := bufio.NewScanner(os.Stdin)
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

  _, err = contract.AddPersonalInfo(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, ssn, license, big.NewInt(dob));
  if err != nil {
    log.Fatalf("### Could not add personal info: %v", err)
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Printf("# Personal info added for address: %s\n", auth.From.String())
}

// function adds login details in the user lockbox
func addLoginCredentials() {
  fmt.Println("# Enter login credentials:")
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter service name:")
  scanner.Scan()
  service := scanner.Text()
  fmt.Println("# Enter username:")
  scanner.Scan()
  username := scanner.Text()
  fmt.Println("# Enter password:")
  scanner.Scan()
  password := scanner.Text()

  _, err = contract.AddLoginCredentials(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, service, username, password); 
  if err != nil {
    log.Fatalf("### Could not add login details: %v", err)
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Printf("# Login details added for address: %s\n", auth.From.String())

}

// function adds credit card details in the user lockbox
func addCreditCardDetails() {
  fmt.Println("# Enter card details:")
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("# Enter card name:")
  scanner.Scan()
  cname := scanner.Text()
  fmt.Println("# Enter card number:")
  scanner.Scan()
  cardno, serr := strconv.ParseUint(scanner.Text(), 10, 64)
  if serr != nil {
    log.Fatalf("### Invalid card number %v", serr)
  }
  fmt.Println("# Enter expiry date:")
  scanner.Scan()
  edate, serr := strconv.ParseInt(scanner.Text(), 10, 64)
  if serr != nil {
    log.Fatalf("### Invalid expiry date %v", serr)
  }
  fmt.Println("# Enter CVV:")
  scanner.Scan()
  cvv, serr := strconv.ParseInt(scanner.Text(), 10, 64)
  if serr != nil {
    log.Fatalf("### Invalid cvv number %v", serr)
  }

  _, err = contract.AddCardDetails(&bind.TransactOpts{
    From:     auth.From,
    Signer:   auth.Signer,
    GasLimit: big.NewInt(2381623),
    Value:    big.NewInt(10),
  }, cname, cardno, big.NewInt(edate), big.NewInt(cvv));
  if err != nil {
    log.Fatalf("### Could not add card details: %v", err)
  }
  fmt.Println("# Mining...")
  sim.Commit()
  fmt.Printf("# Card details added for address: %s\n", auth.From.String())
}

// function gets personal info from the user lockbox
func getPersonalInfo() {
  fmt.Println("\n# Getting SSN of input user account")
  ssn, err := contract.GetSSN(nil, auth.From)
  if err != nil {
    log.Fatalf("### No ssn information found in lockbox account: %v\n", err)
  }
  fmt.Printf("# SSN: %s\n", ssn)

  fmt.Println("\n# Getting License Number of input user account")
  license, err := contract.GetLicenseNo(nil, auth.From)
  if err != nil {
    log.Fatalf("### No ssn information found in lockbox account: %v\n", err)
  }
  fmt.Printf("# License No: %s\n", license)

  fmt.Println("\n# Getting DOB of input user account")
  sdob, err := contract.GetDob(nil, auth.From)
  if err != nil {
    log.Fatalf("### No dob information found in lockbox account: %v\n", err)
  }
  fmt.Println("# DOB: ", sdob)
}

// function gets login details from the user lockbox
func getLoginDetails() {
  fmt.Println("\n# Getting login details of input user account")
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("\n# Enter service")
  scanner.Scan()
  service := scanner.Text()
  ldetails, err := contract.GetLoginCredentials(nil, auth.From, service)
  if err != nil {
    log.Fatalf("### Login details not found in lockbox: %v\n", err)
  }
  fmt.Println("# Service: ", ldetails.Service, " Username: ", ldetails.Uname, " Password: ", ldetails.Passwd)
}

// function gets credit card details from the user lockbox
func getCardDetails() {
  fmt.Println("\n# Getting Card details of input user account")
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("\n# Enter account name")
  scanner.Scan()
  cname := scanner.Text()
  cdetails, err := contract.GetCardDetails(nil, auth.From, cname)
  if err != nil {
    log.Fatalf("### Card details not found in lockbox account: %v\n", err)
  }
  fmt.Println("# Card Number: ", cdetails.Cardno, " Expiry Date: ", cdetails.Edate, " CVV: ", cdetails.Cvv)
}

func main() {
  printLockboxOptions()
  deployLockbox()

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
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }
}
