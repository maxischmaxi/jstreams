"use client";

import Image from "next/image";
import { Match as MatchType } from "@/lib/riot/types/match/match";
import dayjs from "dayjs";
import Link from "next/link";
import relativeTime from "dayjs/plugin/relativeTime";
import duration from "dayjs/plugin/duration";
import de from "dayjs/locale/de";
import { cn } from "@/lib/utils";
import { useMemo } from "react";
import { gameModeToString } from "@/lib/riot/lib";
import { ChevronDown } from "lucide-react";
import { Champion } from "@/lib/riot/champion";
import {
  getChampionSquareAssetUrl,
  getSummonerSpellIcon,
} from "@/lib/assets/riot/images";
import { SummonerSpell } from "@/lib/riot/summonerspell";

dayjs.locale(de);
dayjs.extend(relativeTime);
dayjs.extend(duration);

type Props = {
  match: MatchType;
  puuid: string;
  champions: { [key: string]: Champion };
  summonerSpells: { [key: string]: SummonerSpell };
};

type MatchInfos = {
  win: boolean;
  champ: Champion | null;
  spells: {
    spell1: SummonerSpell | null;
    spell2: SummonerSpell | null;
  };
  runes: {
    primary: number | null;
    secondary: number | null;
  };
};

export function Match(props: Props) {
  const { match, puuid, summonerSpells, champions } = props;

  const timeSinceEnd = dayjs(match.info.gameEndTimestamp).fromNow();
  const gameDuration = dayjs
    .duration(match.info.gameDuration, "seconds")
    .format("mm:ss");

  const { win, champ, spells } = useMemo<MatchInfos>(() => {
    const participant = match.info.participants.find((participant) => {
      return participant.puuid === puuid;
    });

    if (!participant) {
      return {
        win: false,
        spells: {
          spell1: null,
          spell2: null,
        },
        champ: null,
        runes: {
          primary: null,
          secondary: null,
        },
      };
    }

    return {
      win: participant.win,
      champ:
        Object.values(champions).find((champion) => {
          return champion.key === participant.championId.toString();
        }) ?? null,
      spells: {
        spell1:
          Object.values(summonerSpells).find((spell) => {
            return spell.key === participant.summoner1Id.toString();
          }) ?? null,
        spell2:
          Object.values(summonerSpells).find((spell) => {
            return spell.key === participant.summoner2Id.toString();
          }) ?? null,
      },
      runes: {
        primary: participant.perks.styles[0].style,
        secondary: participant.perks.styles[1].style,
      },
    };
  }, [champions, match.info.participants, puuid, summonerSpells]);

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
              {gameModeToString(match.info.gameMode)}
            </p>
            <p className="text-xs">{timeSinceEnd}</p>
          </div>
          <div>
            <p className="text-sm font-bold">{win ? "Sieg" : "Niederlage"}</p>
            <p className="text-xs">{gameDuration}</p>
          </div>
        </div>

        <div className="flex flex-row flex-nowrap gap-2">
          {champ && (
            <Image
              alt={champ.name}
              width={champ.image.w}
              height={champ.image.h}
              src={getChampionSquareAssetUrl({
                patchVersion: "14.20.1",
                championName: champ.image.full,
              })}
              className="w-12 h-12 rounded-full object-center object-cover"
            />
          )}
          <div className="grid grid-cols-2 gap-0">
            {spells.spell1 && (
              <Image
                alt={spells.spell1.name}
                width={spells.spell1.image.w}
                height={spells.spell1.image.h}
                src={getSummonerSpellIcon({
                  patchVersion: "14.20.1",
                  image: spells.spell1.image,
                })}
                title={spells.spell1.name}
                className="w-8 h-8 rounded-full object-center object-cover col-span-1"
              />
            )}
            {spells.spell2 && (
              <Image
                alt={spells.spell2.name}
                width={spells.spell2.image.w}
                height={spells.spell2.image.h}
                src={getSummonerSpellIcon({
                  patchVersion: "14.20.1",
                  image: spells.spell2.image,
                })}
                title={spells.spell2.name}
                className="w-8 h-8 rounded-full object-center object-cover col-span-1"
              />
            )}
          </div>
          <Link
            className="text-link hover:underline"
            href={`/match/${match.metadata.matchId}`}
          >
            {match.info.gameDuration}
          </Link>
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
  );
}
