import { AccountService } from "@maxischmaxi/jstreams-ts/account/v1/account_connect";
import { TierService } from "@maxischmaxi/jstreams-ts/tier/v1/tier_connect";
import { EntriesService } from "@maxischmaxi/jstreams-ts/entries/v1/entries_connect";
import { MatchesService } from "@maxischmaxi/jstreams-ts/matches/v1/matches_connect";
import { VersionService } from "@maxischmaxi/jstreams-ts/version/v1/version_connect";
import { ChampionsService } from "@maxischmaxi/jstreams-ts/champions/v1/champions_connect";
import { SummonerService } from "@maxischmaxi/jstreams-ts/summoner/v1/summoner_connect";
import { AssetsService } from "@maxischmaxi/jstreams-ts/assets/v1/assets_connect";
import { MasteriesService } from "@maxischmaxi/jstreams-ts/masteries/v1/masteries_connect";
import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

export const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  jsonOptions: {
    enumAsInteger: true,
  },
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
