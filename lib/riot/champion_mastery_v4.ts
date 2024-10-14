import "server-only";
import { getRiotApiUrl } from "./lib";
import {
  GetChampionMasteriesByPuuidByChampionRequest,
  GetChampionMasteriesByPuuidByChampionResponse,
  GetChampionMasteriesByPuuidResponse,
  GetChampionMateriesByPuuidRequest,
} from "./types";

export async function getChampionMasteriesByPuuidByChampion(
  props: GetChampionMasteriesByPuuidByChampionRequest,
): Promise<GetChampionMasteriesByPuuidByChampionResponse> {
  const { puuid, championId, region } = props;
  const url = `${getRiotApiUrl(region)}/lol/champion-mastery/v4/champion-masteries/by-puuid/${puuid}/by-champion/${championId}?api_key=${process.env.RIOT_API_KEY}`;
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

export async function getChampionMateriesByPuuid(
  props: GetChampionMateriesByPuuidRequest,
): Promise<GetChampionMasteriesByPuuidResponse> {
  const { puuid, region } = props;
  const url = `${getRiotApiUrl(region)}/lol/champion-mastery/v4/champion-masteries/by-puuid/${puuid}?api_key=${process.env.RIOT_API_KEY}`;
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
