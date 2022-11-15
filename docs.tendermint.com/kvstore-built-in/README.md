# kvstore-built-in

with tendermint v0.34

https://docs.tendermint.com/v0.34/tutorials/go-built-in.html

```bash
# enable tendermint
# ref: https://docs.tendermint.com/v0.34/introduction/install.html
# in other directory
gh repo clone tendermint/tendermint
cd tendermint
git checkout v0.34.23
make install
make build
cd build
cp tendermint (application_directory)

# launch application
./tendermint init
go build
./kvstore-built-in --config $HOME/.tendermint/config/config.toml

# operation check
curl -s 'localhost:26657/broadcast_tx_commit?tx="tendermint=rocks"'
```
