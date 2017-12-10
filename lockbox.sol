pragma solidity ^0.4.19;

contract LockboxApp {
  struct CreditCard {
    string cname;
    uint64 cardno;
    uint edate;
    uint cvv;
  }

  struct LoginInfo {
    string service;
    string uname;
    string passwd;
  }

  struct Lockbox {
    address addr;
    string name;
    string ssn;
    string license;
    uint dob; // MMDDYYYY
    mapping (bytes32 => CreditCard) cardsInfo;
    mapping (bytes32 => LoginInfo) loginCredentials;
    bool initialized;
  }

  address _admin;
  mapping (address => Lockbox) public lockboxAccounts;
  address[] public lockboxAddresses;

  event LockboxAdded(address addr, string name, bool initialized);
  event PersonalInfoAdded(address addr, string ssn, string license, uint dob);
  event CardDetailsAdded(address addr, string cname, uint64 cardno, uint edate, uint cvv);
  event LoginCredentialsAdded(address addr, string service, string uname, string passwd);

  // constructor
  function LockBoxApp() public {
    _admin = msg.sender;
  }

  function stringToBytes32(string memory source) public returns (bytes32 result) {
    bytes memory tempEmptyStringTest = bytes(source);
    if (tempEmptyStringTest.length == 0) {
      return 0x0;
    }

    assembly {
      result := mload(add(source, 32))
    }
  }
  function createLockbox(string lname) payable public returns (bool success){
    if (!lockboxAccounts[msg.sender].initialized) {
      lockboxAccounts[msg.sender].addr = msg.sender;
      lockboxAccounts[msg.sender].name = lname;
      lockboxAccounts[msg.sender].initialized = true;
      lockboxAddresses.push(msg.sender);
      LockboxAdded(msg.sender, lname, lockboxAccounts[msg.sender].initialized);
      return true;
    }
    return false;
  }

  function getLockboxName(address addr)  public constant returns (string name) {
    var lockbox = lockboxAccounts[addr];
    require(lockbox.initialized);
    return lockbox.name;
  }

  function addPersonalInfo(string ssn, string licenseno, uint dob) payable public returns (bool success) {
    var lockbox = lockboxAccounts[msg.sender];
    require(lockbox.initialized);
    lockbox.ssn = ssn;
    lockbox.license = licenseno;
    lockbox.dob = dob;
    PersonalInfoAdded(msg.sender, ssn, licenseno, dob);
    return true;
  }

  function addCardDetails(string cname, uint64 cardno, uint edate, uint cvv) payable public returns (bool success) {
    var lockbox = lockboxAccounts[msg.sender];
    require(lockbox.initialized);
    bytes32 key = stringToBytes32(cname);
    lockbox.cardsInfo[key].cname = cname;
    lockbox.cardsInfo[key].cardno = cardno;
    lockbox.cardsInfo[key].edate = edate;
    lockbox.cardsInfo[key].cvv = cvv;
    CardDetailsAdded(msg.sender, cname, cardno, edate, cvv);
    return true;
  }

  function addLoginCredentials(string service, string uname, string passwd) payable public returns (bool success) {
    var lockbox = lockboxAccounts[msg.sender];
    require(lockbox.initialized);
    bytes32 key = stringToBytes32(service);
    lockbox.loginCredentials[key].service = service;
    lockbox.loginCredentials[key].uname = uname;
    lockbox.loginCredentials[key].passwd = passwd;
    LoginCredentialsAdded(msg.sender, service, uname, passwd);
    return true;
  }

  function getSSN(address addr) public constant returns (string name) {
    var lockbox = lockboxAccounts[addr];
    require(lockbox.initialized);
    return lockbox.ssn;
  }

  function getLicenseNo(address addr) public constant returns (string name) {
    var lockbox = lockboxAccounts[addr];
    require(lockbox.initialized);
    return lockbox.license;
  }

  function getDob(address addr) public constant returns (uint num) {
    var lockbox = lockboxAccounts[addr];
    require(lockbox.initialized);
    return lockbox.dob;
  }
  
  function getCardDetails(address addr, string name) public constant returns (uint64 cardno, uint edate, uint cvv) {
    var lockbox = lockboxAccounts[addr];
    require(lockbox.initialized);
    bytes32 key = stringToBytes32(name);
    var ccInfo = lockbox.cardsInfo[key];
    return (ccInfo.cardno, ccInfo.edate, ccInfo.cvv);
  }

  function getLoginCredentials(address addr, string name) public constant returns (string service, string uname, string passwd) {
    var lockbox = lockboxAccounts[addr];
    require(lockbox.initialized);
    bytes32 key = stringToBytes32(name);
    var lc = lockbox.loginCredentials[key];
    return (lc.service, lc.uname, lc.passwd);
  }
}
