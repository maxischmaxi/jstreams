import Image from "next/image";
import { SummonerSpell } from "@/summoner/v1/summoner_pb";
import "server-only";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "../ui/hover-card";
import { MatchParticipant } from "@/matches/v1/matches_pb";
import { assets } from "@/lib/api";
import { cn } from "@/lib/utils";

type Props = {
  participant: MatchParticipant;
  spell: 1 | 2;
  patchVersion: string;
  summonerSpells: { [key: string]: SummonerSpell };
  className?: string;
};

export async function SummonerSpellIcon(props: Props) {
  const { className, patchVersion, participant, spell, summonerSpells } = props;

  const selectedSpell = Object.values(summonerSpells).find((s) => {
    return (
      s.key ===
      (spell === 1
        ? participant.summoner1Id.toString()
        : participant.summoner2Id.toString())
    );
  });

  const spellIcon = await assets.getSummonerSpellIcon({
    patchVersion,
    image: selectedSpell?.image,
  });

  return (
    <HoverCard>
      <HoverCardTrigger asChild>
        <Image
          alt={selectedSpell?.name ?? ""}
          width={selectedSpell?.image?.w ?? 0}
          height={selectedSpell?.image?.h ?? 0}
          src={spellIcon.url ?? ""}
          title={selectedSpell?.name ?? ""}
          className={cn(
            "w-6 h-6 rounded-full object-center object-cover col-span-1",
            className,
          )}
        />
      </HoverCardTrigger>
      <HoverCardContent>
        <p className="text-xs font-bold">{selectedSpell?.name}</p>
      </HoverCardContent>
    </HoverCard>
  );
}
