import { Image } from "@/lib/riot/types";

export function getChampionSplashAssetUrl(
  championName: string,
  skinNumber: number,
): string {
  return `https://ddragon.leagueoflegends.com/cdn/img/champion/splash/${championName.replaceAll(" ", "")}_${skinNumber}.jpg`;
}

export function getChampionLoadingScreenAssetUrl(
  championName: string,
  skinNumber: number,
): string {
  return `https://ddragon.leagueoflegends.com/cdn/img/champion/loading/${championName.replaceAll(" ", "")}_${skinNumber}.jpg`;
}

type GetChampionSquareAssetUrlProps = {
  championName: string;
  patchVersion: string;
};

export function getChampionSquareAssetUrl(
  props: GetChampionSquareAssetUrlProps,
): string {
  const { championName, patchVersion } = props;
  return `https://ddragon.leagueoflegends.com/cdn/${patchVersion}/img/champion/${championName}`;
}

type GetChampionPassiveAssetUrlProps = {
  patchVersion: string;
  championName: string;
};
export function getChampionPassiveAssetUrl(
  props: GetChampionPassiveAssetUrlProps,
): string {
  const { patchVersion, championName } = props;
  return `https://ddragon.leagueoflegends.com/cdn/${patchVersion}/img/passive/${championName.replaceAll(" ", "")}_P.png`;
}

type GetChampionAbilityAssetUrlProps = {
  championName: string;
  abilityNumber: number;
  patchVersion: string;
};

export function getChampionAbilityAssetUrl(
  props: GetChampionAbilityAssetUrlProps,
): string {
  const { championName, patchVersion, abilityNumber } = props;
  return `https://ddragon.leagueoflegends.com/cdn/${patchVersion}/img/champion/${championName.replaceAll(" ", "")}_${abilityNumber}.png`;
}

type GetItemAssetUrlProps = {
  itemName: string;
  patchVersion: string;
};

type GetSummonerSpellIconProps = {
  patchVersion: string;
  image: Image;
};

type GetRuneIconProps = {
  patchVersion: string;
  style: number;
};

export function getRuneIcon(props: GetRuneIconProps): string {
  return `https://ddragon.leagueoflegends.com/cdn/img/${props.patchVersion}/img/${props.style}`;
}

export function getSummonerSpellIcon(props: GetSummonerSpellIconProps): string {
  return `https://ddragon.leagueoflegends.com/cdn/${props.patchVersion}/img/spell/${props.image.full}`;
}

export function getItemAssetUrl(props: GetItemAssetUrlProps): string {
  const { itemName, patchVersion } = props;
  return `https://ddragon.leagueoflegends.com/cdn/${patchVersion}/data/en_US/${itemName}.json`;
}

type GetSpellAssetUrlProps = {
  spellName: string;
  patchVersion: string;
};

export function getSpellAssetUrl(props: GetSpellAssetUrlProps): string {
  const { spellName, patchVersion } = props;
  return `https://ddragon.leagueoflegends.com/cdn/${patchVersion}/img/spell/${spellName}.png`;
}
