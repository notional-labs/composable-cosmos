import { ApiPromise, WsProvider } from "@polkadot/api";

export const getProvider = async (): Promise<ApiPromise> => {
  // Initialise the provider to connect to the local node
  const wsEndpoint = "ws://65.21.224.114:9944";
  const provider = new WsProvider(wsEndpoint);

  console.log(`connection to provider at ${wsEndpoint}`);

  // Create the API and wait until ready
  const api = await ApiPromise.create({ provider });

  return api;
};
