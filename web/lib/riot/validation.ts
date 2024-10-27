import { z } from "zod";
import { RegionalRoutingValues } from "@maxischmaxi/jstreams-ts/account_pb";
import { PlatformRoutingValues } from "@maxischmaxi/jstreams-ts/summoner_pb";

export const getChampionMasteriesByPuuidValidation = z.object({
  puuid: z.string().min(1),
  region: z.nativeEnum(RegionalRoutingValues),
});

export const getChampionMateriesByPuuidForChampionValidation = z.object({
  puuid: z.string().min(1),
  region: z.nativeEnum(PlatformRoutingValues),
  championId: z.number().int(),
});

export const getAccountByGamenameAndTaglineValidation = z.object({
  gameName: z.string().min(1),
  region: z.nativeEnum(RegionalRoutingValues),
});

export const getAccountByGamenameAndTaglineResponseValidation = z.object({
  gameName: z.string(),
  tagLine: z.string(),
  puuid: z.string(),
});

export const getMatchByMatchIdValidation = z.object({
  matchId: z.string().min(1),
  region: z.nativeEnum(PlatformRoutingValues),
});
