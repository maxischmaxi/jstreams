import { Participant } from "./participant";
import { Team } from "./team";

export enum GameMode {
  CLASSIC = "CLASSIC",
  ARAM = "ARAM",
}

export enum GameType {
  MATCHED_GAME = "MATCHED_GAME",
}

export type Match = {
  metadata: {
    dataVersion: string;
    matchId: string;
    participants: string[];
  };
  info: {
    endOfGameResult: string;
    gameCreation: number;
    gameDuration: number;
    gameEndTimestamp: number;
    gameId: number;
    gameMode: GameMode;
    gameName: string;
    gameStartTimestamp: number;
    gameType: GameType;
    gameVersion: string;
    mapId: number;
    platformId: string;
    queueId: number;
    tournamentCode: string;
    participants: Participant[];
    teams: Team[];
  };
};
