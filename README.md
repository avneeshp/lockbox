# lockbox

Steps to run lockbox application on simulated backend:
1. Build and run application
=> go build; ./lockbox

Steps to run lockbox application on test network:
1. Set up and start ethereum client
=> geth --testnet –genesis <genesis.json> -–networkid 321 -–nodiscover -–datadir <data_dir_path> –-port 30304 --maxpeers 0 console
2. Build and run application
=> go build; ./lockbox
