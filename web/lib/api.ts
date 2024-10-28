import { AccountService } from "@/account/v1/account_connect";
import { TierService } from "@/tier/v1/tier_connect";
import { EntriesService } from "@/entries/v1/entries_connect";
import { MatchesService } from "@/matches/v1/matches_connect";
import { VersionService } from "@/version/v1/version_connect";
import { ChampionsService } from "@/champions/v1/champions_connect";
import { SummonerService } from "@/summoner/v1/summoner_connect";
import { AssetsService } from "@/assets/v1/assets_connect";
import { MasteriesService } from "@/masteries/v1/masteries_connect";
import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ItemsService } from "@/items/v1/items_connect";

export const transport = createConnectTransport({
  baseUrl: process.env.GATEWAY ?? "",
  jsonOptions: {
    enumAsInteger: true,
  },
  fetch: (input: RequestInfo | URL, init?: RequestInit) => {
    return fetch(input, {
      ...init,
      redirect: "follow",
    });
  },
  useHttpGet: true,
});

export const account = createPromiseClient(AccountService, transport);
export const tier = createPromiseClient(TierService, transport);
export const entries = createPromiseClient(EntriesService, transport);
export const matches = createPromiseClient(MatchesService, transport);
export const version = createPromiseClient(VersionService, transport);
export const champions = createPromiseClient(ChampionsService, transport);
export const summoner = createPromiseClient(SummonerService, transport);
export const assets = createPromiseClient(AssetsService, transport);
export const masteries = createPromiseClient(MasteriesService, transport);
export const items = createPromiseClient(ItemsService, transport);
