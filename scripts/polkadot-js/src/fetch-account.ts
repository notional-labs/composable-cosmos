const { ApiPromise, WsProvider } = require("@polkadot/api");

// Put the address of the account you want to fetch info for here
const address = "5YourAccountAddressHere";

async function fetchAccountInfo() {
  // Initialise the provider to connect to the local node
  const provider = new WsProvider("wss://polkadot.api.onfinality.io/public-ws");

  // Create the API instance
  const api = await ApiPromise.create({ provider });

  try {
    // Fetch the account info
    const accountInfo = await api.query.system.account(address);

    console.log(`Account ${address} info:`, accountInfo.toJSON());
  } catch (error) {
    console.error("Error fetching account info:", error);
  } finally {
    // Disconnect the provider when done
    provider.disconnect();
  }
}

fetchAccountInfo();
