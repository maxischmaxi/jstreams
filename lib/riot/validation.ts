import { z } from "zod";
import { REGIONAL_ROUTING_VALUES, PLATFORM_ROUTING_VALUES } from "./lib";
import { GetAccountByGamenameAndTaglineValidation } from "./types";

export const getChampionMasteriesByPuuidValidation = z.object({
  puuid: z.string().min(1),
  region: z.nativeEnum(PLATFORM_ROUTING_VALUES),
});

export const getChampionMateriesByPuuidForChampionValidation = z.object({
  puuid: z.string().min(1),
  region: z.nativeEnum(PLATFORM_ROUTING_VALUES),
  championId: z.number().int(),
});

export const getAccountByGamenameAndTaglineValidation: GetAccountByGamenameAndTaglineValidation =
  z.object({
    gameName: z.string().min(1),
    region: z.nativeEnum(REGIONAL_ROUTING_VALUES),
  });

export const getAccountByGamenameAndTaglineResponseValidation = z.object({
  gameName: z.string(),
  tagLine: z.string(),
  puuid: z.string(),
});

export const getMatchByMatchIdValidation = z.object({
  matchId: z.string().min(1),
  region: z.nativeEnum(REGIONAL_ROUTING_VALUES),
});
