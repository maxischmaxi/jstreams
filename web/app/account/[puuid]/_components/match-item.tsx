import "server-only";
import { SummonerSpellIcon } from "@/components/icons/summonerspell-icon";
import { Champion } from "@/champions/v1/champions_pb";
import { gameModeToString } from "@/lib/riot/lib";
import { cn } from "@/lib/utils";
import { SummonerSpell } from "@/summoner/v1/summoner_pb";
import dayjs from "dayjs";
import { ChevronDown } from "lucide-react";
import de from "dayjs/locale/de";
import fromNow from "dayjs/plugin/relativeTime";
import duration from "dayjs/plugin/duration";
import { ItemIcon } from "@/components/icons/item-icon";
import { ChampionIcon } from "@/components/icons/champion-icon";
import { Match } from "@/matches/v1/matches_pb";
import { ParticipantListItem } from "./participant-list-item";

dayjs.locale(de);
dayjs.extend(fromNow);
dayjs.extend(duration);

type Props = {
  summonerSpells: { [key: string]: SummonerSpell };
  puuid: string;
  champions: { [key: string]: Champion };
  match: Match;
  patchVersion: string;
};

export async function MatchItem(props: Props) {
  const { match, summonerSpells, patchVersion, champions, puuid } = props;

  const participant = match.info?.participants.find((participant) => {
    return participant.puuid === puuid;
  });

  if (!participant) {
    return null;
  }

  const win = participant.win;
  const timeSinceEnd = dayjs(Number(match.info?.gameEndTimestamp)).fromNow();
  const gameDuration = dayjs
    .duration(match.info?.gameDuration ?? 0, "seconds")
    .format("mm:ss");
  const kda =
    Math.round(
      ((participant.kills + participant.assists) /
        Math.max(1, participant.deaths)) *
        100,
    ) / 100;

  return (
    <li
      className={cn(
        win
          ? "bg-win-background border-win-foreground"
          : "bg-lose-background border-lose-foreground",
        "h-[100px] w-full rounded-md overflow-hidden flex flex-row justify-center items-center flex-nowrap border-l-4",
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
              {gameModeToString(match.info?.gameMode ?? "")}
            </p>
            <p className="text-xs">{timeSinceEnd}</p>
          </div>
          <div>
            <p className="text-sm font-bold">{win ? "Sieg" : "Niederlage"}</p>
            <p className="text-xs">{gameDuration}</p>
          </div>
        </div>

        <div className="flex flex-row flex-nowrap w-full gap-2">
          <div className="flex flex-col justify-between shrink-0">
            <div className="flex flex-row flex-nowrap gap-2 shrink-0">
              <div className="shrink-0">
                <ChampionIcon
                  champions={champions}
                  participant={participant}
                  patchVersion={patchVersion}
                />
              </div>
              <div className="grid grid-cols-2 gap-1 shrink-0">
                <SummonerSpellIcon
                  participant={participant}
                  spell={1}
                  patchVersion={patchVersion}
                  summonerSpells={summonerSpells}
                />
                <SummonerSpellIcon
                  participant={participant}
                  spell={2}
                  patchVersion={patchVersion}
                  summonerSpells={summonerSpells}
                />
              </div>
              <div className="flex flex-col flex-nowrap">
                <p className="font-bold whitespace-nowrap">
                  {participant.kills} /{" "}
                  <span className="text-red-600">{participant.deaths}</span> /{" "}
                  {participant.assists}
                </p>
                <p className="text-xs text-muted-foreground">KDA: {kda}</p>
              </div>
            </div>
            <div className="flex flex-row flex-nowrap">
              <ItemIcon
                patchVersion={patchVersion}
                itemId={participant.item0}
              />
              <ItemIcon
                patchVersion={patchVersion}
                itemId={participant.item1}
              />
              <ItemIcon
                patchVersion={patchVersion}
                itemId={participant.item2}
              />
              <ItemIcon
                patchVersion={patchVersion}
                itemId={participant.item3}
              />
              <ItemIcon
                patchVersion={patchVersion}
                itemId={participant.item4}
              />
              <ItemIcon
                patchVersion={patchVersion}
                itemId={participant.item5}
              />
              <ItemIcon
                patchVersion={patchVersion}
                itemId={participant.item6}
              />
            </div>
          </div>
          <div className="w-full flex justify-end text-xs flex-row flex-nowrap gap-2">
            <ul className="w-[100px] flex flex-col justify-between">
              {match.info?.participants
                ?.filter((p) => p.teamId === 100)
                .map((p) => (
                  <ParticipantListItem
                    key={p.puuid}
                    participant={p}
                    puuid={puuid}
                    patchVersion={patchVersion}
                  />
                ))}
            </ul>
            <ul className="w-[100px] flex flex-col justify-between">
              {match.info?.participants
                ?.filter((p) => p.teamId === 200)
                .map((p) => (
                  <ParticipantListItem
                    key={p.puuid}
                    participant={p}
                    puuid={puuid}
                    patchVersion={patchVersion}
                  />
                ))}
            </ul>
          </div>
        </div>
      </div>
      <button
        type="button"
        className={cn(
          win ? "bg-win-accent" : "bg-lose-accent",
          "px-2 flex-1 h-full flex flex-col justify-end pb-4 items-end",
        )}
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
