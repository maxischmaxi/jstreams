import { Image } from "./types";

export type SummonerSpell = {
  cooldown: number[];
  cooldownBurn: string;
  cost: number[];
  costBurn: string;
  costType: string;
  datavalues: object;
  description: string;
  effect: Array<number[] | null>;
  effectBurn: Array<string | null>;
  id: string;
  image: Image;
  key: string;
  maxammo: string;
  maxrank: number;
  modes: string[];
  name: string;
  range: number[];
  rangeBurn: string;
  resource: string;
  summonerLevel: number;
  tooltip: string;
  vars: Array<{
    link: string;
    coeff: number;
    key: string;
  }>;
};

export type GetSummonerSpellsResponse = {
  type: "summoner";
  version: string;
  data: {
    [key: string]: SummonerSpell;
  };
};
export async function getSummonerSpells(
  patchVersion: string,
): Promise<GetSummonerSpellsResponse> {
  const uri = `https://ddragon.leagueoflegends.com/cdn/${patchVersion}/data/en_US/summoner.json`;
  const url = new URL(uri);
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
