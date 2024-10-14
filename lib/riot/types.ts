import { ZodNativeEnum, ZodObject, ZodString } from "zod";
import { REGIONAL_ROUTING_VALUES, PLATFORM_ROUTING_VALUES } from "./lib";
import { Match } from "./types/match/match";
import { Timeline } from "./types/match/timeline";

export type GetAccountByGamenameAndTaglineRequest = {
  gameName: string;
  tagLine: string;
  region: REGIONAL_ROUTING_VALUES;
};

export type GetAccountByGamenameAndTaglineResponse = {
  puuid: string;
  gameName: string;
  tagLine: string;
};

export type GetChampionMateriesByPuuidRequest = {
  puuid: string;
  region: PLATFORM_ROUTING_VALUES;
};

export type ChampionMastery = {
  puuid: string;
  championId: number;
  championLevel: number;
  championPoints: number;
  lastPlayTime: number;
  championPointsSinceLastLevel: number;
  championPointsUntilNextLevel: number;
  markRequiredForNextLevel: number;
  tokensEarned: number;
  championSeasonMilestone: number;
  milestoneGrades: string[];
  nextSeasonMilestone: {
    requireGradeCounts: { [key: string]: number };
    rewardMarks: number;
    bonus: boolean;
    totalGamesRequires: number;
  };
};

export type GetChampionMasteriesByPuuidResponse = ChampionMastery[];

export type GetAccountByGamenameAndTaglineValidation = ZodObject<{
  gameName: ZodString;
  region: ZodNativeEnum<typeof REGIONAL_ROUTING_VALUES>;
}>;

export type GetChampionMasteriesByPuuidByChampionRequest = {
  puuid: string;
  championId: number;
  region: PLATFORM_ROUTING_VALUES;
};

export type GetChampionMasteriesByPuuidByChampionResponse = ChampionMastery;

export type GetMatchesByPuuidRequest = {
  puuid: string;
  region: REGIONAL_ROUTING_VALUES;
  startTime?: number;
  endTime?: number;
  queue?: number;
  type?: string;
  start?: number;
  count?: number;
};

export type GetMatchesByPuuidResponse = string[];

export type GetMatchByMatchIdRequest = {
  matchId: string;
  region: REGIONAL_ROUTING_VALUES;
};

export type GetMatchByMatchIdResponse = Match;

export type Selection = {
  perk: number;
  var1: number;
  var2: number;
  var3: number;
};

export type GetMatchTimelineRequest = {
  matchId: string;
  region: REGIONAL_ROUTING_VALUES;
};

export type GetMatchTimelineResponse = Timeline;

export type Image = {
  full: string;
  sprite: string;
  group: string;
  x: number;
  y: number;
  w: number;
  h: number;
};
