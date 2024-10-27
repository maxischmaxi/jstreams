import { RegionalRoutingValues } from "@/account/v1/account_pb";

export function accountRegionToTagline(region: RegionalRoutingValues): string {
  switch (region) {
    case RegionalRoutingValues.EUROPE:
      return "EUW";
    case RegionalRoutingValues.AMERICAS:
      return "NA";
    case RegionalRoutingValues.ASIA:
      return "KR";
    case RegionalRoutingValues.ESPORTS:
      return "ESPORTS";
    default:
      return "EUW";
  }
}

export function accountRegionToString(region: RegionalRoutingValues): string {
  switch (region) {
    case RegionalRoutingValues.EUROPE:
      return "Europe West";
    case RegionalRoutingValues.AMERICAS:
      return "North America";
    case RegionalRoutingValues.ASIA:
      return "Asia";
    case RegionalRoutingValues.ESPORTS:
      return "Esports";
    default:
      return "Europe";
  }
}

export function gameTypeToString(gameType: string): string {
  switch (gameType) {
    case "MATCHED_GAME":
      return "Normales Spiel";
    default:
      return "Unbekannter Spieltyp";
  }
}

export function gameModeToString(gameMode: string): string {
  switch (gameMode) {
    case "ARAM":
      return "ARAM";
    case "CLASSIC":
      return "Normales Spiel";
    default:
      return "Unbekannter Spielmodus";
  }
}
