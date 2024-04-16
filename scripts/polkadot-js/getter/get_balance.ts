import { ApiPromise, WsProvider } from "@polkadot/api";
import { getProvider } from "../utils/indexer";

// Put the address of the account you want to fetch info for here
const address = "5EZntS5SYEpBXmJpz8H6hXcGF86ukiugDuEn4si4usJ1bNkb";

async function fetchAccountInfo() {
  // Initialise the provider to connect to the local node

  // Create the API instance
  const api = await getProvider();

  try {
    // Fetch the account info
    const accountInfo = await api.query.system.account(address);

    console.log(`Account ${address} info:`, accountInfo.toJSON());
  } catch (error) {
    console.error("Error fetching account info:", error);
  } finally {
    // Disconnect the provider when done
    api.disconnect();
  }
}

fetchAccountInfo();
