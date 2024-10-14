import "server-only";
import { getRiotApiUrl } from "./lib";
import {
  GetMatchByMatchIdRequest,
  GetMatchByMatchIdResponse,
  GetMatchesByPuuidRequest,
  GetMatchesByPuuidResponse,
  GetMatchTimelineRequest,
  GetMatchTimelineResponse,
} from "./types";

export async function getMatchTimeline(
  props: GetMatchTimelineRequest,
): Promise<GetMatchTimelineResponse> {
  const { matchId, region } = props;
  const url = new URL(
    `${getRiotApiUrl(region)}/lol/match/v5/matches/${matchId}/timeline`,
  );
  url.searchParams.append("api_key", process.env.RIOT_API_KEY ?? "");

  const response = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Origin: "https://developer.riotgames.com",
      "Accept-Charset": "application/x-www-form-urlencoded; charset=UTF-8",
    },
  });

  return await response.json();
}

export async function getMatchByMatchId(
  props: GetMatchByMatchIdRequest,
): Promise<GetMatchByMatchIdResponse> {
  const { matchId, region } = props;
  const uri = `${getRiotApiUrl(region)}/lol/match/v5/matches/${matchId}`;
  const url = new URL(uri);
  url.searchParams.append("api_key", process.env.RIOT_API_KEY ?? "");

  const response = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Origin: "https://developer.riotgames.com",
      "Accept-Charset": "application/x-www-form-urlencoded; charset=UTF-8",
    },
  });

  return await response.json();
}

export async function getMatchIdsByPuuid(
  props: GetMatchesByPuuidRequest,
): Promise<GetMatchesByPuuidResponse> {
  const { puuid, region, start, count, type, queue, endTime, startTime } =
    props;

  const uri = `${getRiotApiUrl(region)}/lol/match/v5/matches/by-puuid/${puuid}/ids`;
  const url = new URL(uri);
  url.searchParams.append("api_key", process.env.RIOT_API_KEY ?? "");

  if (start) {
    url.searchParams.append("start", start.toString());
  }

  if (count) {
    url.searchParams.append("count", count.toString());
  }

  if (type) {
    url.searchParams.append("type", type);
  }

  if (queue) {
    url.searchParams.append("queue", queue.toString());
  }

  if (endTime) {
    url.searchParams.append("endTime", endTime.toString());
  }

  if (startTime) {
    url.searchParams.append("startTime", startTime.toString());
  }

  const response = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Origin: "https://developer.riotgames.com",
      "Accept-Charset": "application/x-www-form-urlencoded; charset=UTF-8",
    },
  });
  return await response.json();
}
