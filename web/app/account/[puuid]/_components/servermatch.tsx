import { RegionalRoutingValues } from "@/account/v1/account_pb";
import Image from "next/image";
import { Champion } from "@/champions/v1/champions_pb";
import { assets, matches } from "@/lib/api";
import { gameModeToString } from "@/lib/riot/lib";
import { cn } from "@/lib/utils";
import { SummonerSpell } from "@/summoner/v1/summoner_pb";
import dayjs from "dayjs";
import { headers } from "next/headers";
import Link from "next/link";
import { ChevronDown } from "lucide-react";
import de from "dayjs/locale/de";
import fromNow from "dayjs/plugin/relativeTime";
import duration from "dayjs/plugin/duration";

dayjs.locale(de);
dayjs.extend(fromNow);
dayjs.extend(duration);

type Props = {
  matchId: string;
  summonerSpells: { [key: string]: SummonerSpell };
  puuid: string;
  champions: { [key: string]: Champion };
}

export async function ServerMatch(props: Props) {
  const { matchId, summonerSpells, champions, puuid } = props;
  const h = headers();
  const version = h.get("x-version") ?? "";

  const match = await matches.getMatchByMatchId({
    matchId: matchId,
    region: RegionalRoutingValues.EUROPE,
  })

  if (!match.match) {
    return null;
  }

  const participant = match.match.info?.participants.find((participant) => {
    return participant.puuid = puuid;
  });

  if (!participant) {
    return null;
  }

  const champ = Object.values(champions).find((c) => {
    return c.key == participant.championId.toString();
  });

  if (!champ) {
    return null;
  }

  const spell1 = Object.values(summonerSpells).find((s) => {
    return s.key === participant.summoner1Id.toString();
  })

  const spell2 = Object.values(summonerSpells).find((s) => {
    return s.key === participant.summoner2Id.toString();
  });

  if (!spell1 || !spell2) {
    return null;
  }

  const squareAsset = await assets.getChampionSquareAssetUrl({
    patchVersion: version,
    championName: champ.image?.full ?? "",
  });

  const summonerSpellIcon1 = await assets.getSummonerSpellIcon({
    patchVersion: version,
    image: spell1.image,
  });

  const summonerSpellIcon2 = await assets.getSummonerSpellIcon({
    patchVersion: version,
    image: spell2.image,
  });

  const win = participant.win;
  const timeSinceEnd = dayjs(Number(match.match.info?.gameEndTimestamp)).fromNow();
  const gameDuration = dayjs
    .duration(match.match.info?.gameDuration ?? 0, "seconds")
    .format("mm:ss");

  const item0 = participant.item0 ?? await assets.getItemAssetUrl({
    patchVersion: version,
    itemId: participant.item0,
  }) ?? null;

  return (
    <li
      className={cn(
        "h-[100px] w-full bg-win rounded-md overflow-hidden flex flex-row justify-center items-center flex-nowrap",
      )}
    >
      <div className="h-full w-full p-2 flex flex-row flex-nowrap">
        <div className="flex flex-col justify-between h-full w-[110px] max-w-[110px] min-w-[110px]">
          <div>
            <p
              className={cn(
                "text-sm font-bold",
                win ? "text-win-foreground" : "text-lose-foreground",
              )}
            >
              {gameModeToString(match.match.info?.gameMode ?? "")}
            </p>
            <p className="text-xs">{timeSinceEnd}</p>
          </div>
          <div>
            <p className="text-sm font-bold">{win ? "Sieg" : "Niederlage"}</p>
            <p className="text-xs">{gameDuration}</p>
          </div>
        </div>

        <div className="flex flex-col justify-between w-full h-full">
          <div className="flex flex-row flex-nowrap gap-2">
            {champ && (
              <Image
                alt={champ.name}
                width={champ.image?.w}
                height={champ.image?.h}
                src={squareAsset.url ?? ""}
                className="w-12 h-12 rounded-full object-center object-cover"
              />
            )}
            <div className="grid grid-cols-2 gap-0">
              <Image
                alt={spell1.name}
                width={spell1.image?.w}
                height={spell1.image?.h}
                src={summonerSpellIcon1.url ?? ""}
                title={spell1.name}
                className="w-8 h-8 rounded-full object-center object-cover col-span-1"
              />
              <Image
                alt={spell2.name}
                width={spell2.image?.w}
                height={spell2.image?.h}
                src={summonerSpellIcon2.url ?? ""}
                title={spell2.name}
                className="w-8 h-8 rounded-full object-center object-cover col-span-1"
              />
            </div>
          </div>
          <div className="flex flex-row flex-nowrap">
            <Image src={participant.item0}
          </div>
        </div>

      </div>
      <button
        type="button"
        className="px-2 flex-1 h-full bg-[#2f436e] flex flex-col justify-end pb-4 items-end"
      >
        <ChevronDown
          className={cn(
            "w-4 h-4",
            win ? "text-win-foreground" : "text-lose-foreground",
          )}
        />
      </button>
    </li>
  )
}
