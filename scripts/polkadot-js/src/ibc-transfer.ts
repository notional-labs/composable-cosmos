import { ApiPromise, WsProvider } from "@polkadot/api";
import { Keyring } from "@polkadot/keyring";
import BN from "bn.js";
import { KeyringPair } from "@polkadot/keyring/types";
import { getProvider, getWallets } from "../utils/indexer";

async function sendIbcFundsTx(
  api: ApiPromise,
  senderKeypair: KeyringPair,
  channelID: string,
  amount: { denom: string; amount: string; address: string },
  options: any
) {
  {
    // Ensure the API is connected
    if (!api.isConnected) {
      await api.connect();
    }

    // Fetch the latest metadata
    const meta = await api.rpc.state.getMetadata();

    // Convert asset number and amount
    const assetNum = new BN(amount.denom, 10);
    const amountBN = new BN(amount.amount, 10);

    // Construct the call to the IBC transfer function
    // Note: The exact method name and parameters depend on your chain's IBC implementation
    const call = api.tx.ibc.transfer({
      raw: 1, // Example parameter, adjust according to actual call structure
      size: amount.address.length * 4,
      to: amount.address,
      channel: new BN(channelID), // Assuming channelID is a string that can be converted
      timeout: 1,
      timestamp: 0, // Example, adjust as needed
      height: 3000, // Example, adjust as needed
      assetId: assetNum,
      amount: amountBN,
      memo: 0, // Example, adjust as needed
    });

    // Sign and send the transaction
    return await new Promise((resolve, reject) => {
      call
        .signAndSend(
          senderKeypair,
          { nonce: -1 },
          ({ status, dispatchError }) => {
            if (status.isInBlock || status.isFinalized) {
              if (dispatchError) {
                if (dispatchError.isModule) {
                  // For module errors, we have the section indexed, lookup
                  const decoded = api.registry.findMetaError(
                    dispatchError.asModule
                  );
                  const { docs, name, section } = decoded;
                  reject(new Error(`${section}.${name}: ${docs.join(" ")}`));
                } else {
                  // Other, CannotLookup, BadOrigin, no extra info
                  reject(new Error(dispatchError.toString()));
                }
              } else {
                resolve(status.asFinalized.toString());
              }
            }
          }
        )
        .catch(reject);
    });
  }
}
// Example usage
async function main() {
  const api = await getProvider();
  const wallets = getWallets();
  const senderKeypair = wallets.alice;

  const channelID = "channel-0"; // Example channel ID
  const amount = { denom: "10", amount: "1000", address: "targetAddress" }; // Example amount
  const options = {}; // Example options, adjust as needed

  try {
    const hash = await sendIbcFundsTx(
      api,
      senderKeypair,
      channelID,
      amount,
      options
    );
    console.log("Transaction hash:", hash);
  } catch (error) {
    console.error("Error sending IBC funds:", error);
  }
}

main().catch(console.error);
