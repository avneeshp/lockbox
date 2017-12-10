package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
  key, _ := crypto.GenerateKey()
  auth := bind.NewKeyedTransactor(key)

  alloc := make(core.GenesisAlloc)
  alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
  sim := backends.NewSimulatedBackend(alloc)

  // Deploy LockboxApp contract
  addr, _, contract, err := DeployLockboxApp(auth, sim)
  if err != nil {
    log.Fatalf("### Could not deploy contract: %v", err)
  }

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
