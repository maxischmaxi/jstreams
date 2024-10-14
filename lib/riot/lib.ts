import { GameMode, GameType } from "./types/match/match";

export enum REGIONAL_ROUTING_VALUES {
  EUROPE = "europe",
  AMERICAS = "americas",
  ASIA = "asia",
  ESPORTS = "esports",
}

export enum PLATFORM_ROUTING_VALUES {
  BR1 = "br1",
  EUN1 = "eun1",
  EUW1 = "euw1",
  JP1 = "jp1",
  KR = "kr",
  LA1 = "la1",
  LA2 = "la2",
  ME1 = "me1",
  NA1 = "na1",
  OC1 = "oc1",
  PH1 = "ph1",
  RU = "ru",
  SG2 = "sg2",
  TH2 = "th2",
  TR1 = "tr1",
  TW2 = "tw2",
  VN2 = "vn2",
}

export function getRiotApiUrl(
  region: REGIONAL_ROUTING_VALUES | PLATFORM_ROUTING_VALUES,
) {
  return `https://${region}.api.riotgames.com`;
}

export function accountRegionToTagline(
  region: REGIONAL_ROUTING_VALUES,
): string {
  switch (region) {
    case REGIONAL_ROUTING_VALUES.EUROPE:
      return "EUW";
    case REGIONAL_ROUTING_VALUES.AMERICAS:
      return "NA";
    case REGIONAL_ROUTING_VALUES.ASIA:
      return "KR";
    case REGIONAL_ROUTING_VALUES.ESPORTS:
      return "ESPORTS";
    default:
      return "EUW";
  }
}

export function accountRegionToString(region: REGIONAL_ROUTING_VALUES): string {
  switch (region) {
    case REGIONAL_ROUTING_VALUES.EUROPE:
      return "Europe West";
    case REGIONAL_ROUTING_VALUES.AMERICAS:
      return "North America";
    case REGIONAL_ROUTING_VALUES.ASIA:
      return "Asia";
    case REGIONAL_ROUTING_VALUES.ESPORTS:
      return "Esports";
    default:
      return "Europe";
  }
}

export function gameTypeToString(gameType: GameType): string {
  switch (gameType) {
    case GameType.MATCHED_GAME:
      return "Normales Spiel";
    default:
      return "Unbekannter Spieltyp";
  }
}

export function gameModeToString(gameMode: GameMode): string {
  switch (gameMode) {
    case GameMode.ARAM:
      return "ARAM";
    case GameMode.CLASSIC:
      return "Normales Spiel";
    default:
      return "Unbekannter Spielmodus";
  }
}
