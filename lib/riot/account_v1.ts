import "server-only";
import { getRiotApiUrl, REGIONAL_ROUTING_VALUES } from "./lib";
import {
  GetAccountByGamenameAndTaglineRequest,
  GetAccountByGamenameAndTaglineResponse,
} from "./types";

export async function getAccountByPuuid(puuid: string) {
  const url = `${getRiotApiUrl(REGIONAL_ROUTING_VALUES.EUROPE)}/riot/account/v1/accounts/by-puuid/${puuid}?api_key=${process.env.RIOT_API_KEY}`;
  const response = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Origin: "https://developer.riotgames.com",
      "Accept-Charset": "application/x-www-form-urlencoded; charset=UTF-8",
    },
  });
  return response.json();
}

export async function getAccountByGamenameAndTagline(
  props: GetAccountByGamenameAndTaglineRequest,
): Promise<GetAccountByGamenameAndTaglineResponse> {
  const { gameName, tagLine, region } = props;
  const url = `${getRiotApiUrl(region)}/riot/account/v1/accounts/by-riot-id/${gameName}/${tagLine}?api_key=${process.env.RIOT_API_KEY}`;
  const response = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Origin: "https://developer.riotgames.com",
      "Accept-Charset": "application/x-www-form-urlencoded; charset=UTF-8",
    },
  });
  return response.json();
}
