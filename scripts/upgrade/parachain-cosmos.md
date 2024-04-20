This notes steps to test parachain - cosmos upgrade.

## Preupgrade Steps

1. Setup polkadot node

```bash
rm -rf /tmp/composable-devnet

cd composable-comsos
./scripts/upgrade/setup-polkadot.sh
```

2. Run cosmos chain and deploy wasm client

```bash
./scripts/upgrade/setup-old-cosmos-node.sh
```

3. Run relayer

```bash
./scripts/upgrade/setup-relayer.sh
```

4. IBC transfer from polkadot to cosmos

Transfer from polkdot to: `pica1hj5fveer5cjtn4wd6wstzugjfdxzl0xpas3hgy`

```bash
cd scripts/polkadot-js
ts-node src/ibc-transfer.ts
```

Then assert balances

```bash
ts-node getter/get_balances.ts
```

Prebalances should be:

```json
Account 5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY info: {
nonce: '27',
consumers: '2',
providers: '1',
sufficients: '0',
data: {
free: '1,149,227,213,405,701,585',
reserved: '3,000,000,000,000,000',
frozen: '0',
flags: '170,141,183,460,469,231,731,687,303,715,884,105,728'
}
}
```

Wait a bit for token transfer to completed

Then, check the balance on the cosmos side

```json
balances:
- amount: "995000000000000"
  denom: ibc/632DBFDB06584976F1351A66E873BF0F7A19FAA083425FEC9890C90993E5F0A4
- amount: "99999999989969990005572311"
  denom: ppica
pagination:
  next_key: null
  total: "0"
```

5. IBC transfer a half back to cosmos

```bash
_build/old/centaurid tx ibc-transfer transfer transfer channel-0 5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY 495000000000000ibc/632DBFDB06584976F1351A66E873BF0F7A19FAA083425FEC9890C90993E5F0A4 --from mykey --keyring-backend test --home mytestnet --chain-id centauri-dev --fees 200ppica -y
```

## Upgrade

```
./scripts/upgrade/upgrade.sh
```

## Postupgrade Steps

1. Stop the relayer

2. Update the account prefix

```bash
vim /tmp/composable-devnet/picasso-centauri-ibc/config-chain-b.toml

## Change `account prefix: centaurid -> pica
```

3. Restart the relayer

```bash
cd _build/composable
nix run .#picasso-centauri-ibc-relay
```

4. Test transfer back and forth

4.1 Centaurid -> Picasso

```bash
_build/new/picad tx ibc-transfer transfer transfer channel-0 5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY 495000000000000ibc/632DBFDB06584976F1351A66E873BF0F7A19FAA083425FEC9890C90993E5F0A4 --from mykey --keyring-backend test --home mytestnet --chain-id centauri-dev --fees 200ppica -y
```

4.2 Picasso -> Centaurid
Change address in file line 93 `scripts/polkadot-js/src/ibc-transfer.ts` to `pica1hj5fveer5cjtn4wd6wstzugjfdxzl0xpas3hgy` and run

```bash
ts-node src/ibc-transfer.ts
```
