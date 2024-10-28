import "server-only";
import { Champion } from "@/champions/v1/champions_pb";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "../ui/hover-card";
import Image from "next/image";
import { MatchParticipant } from "@/matches/v1/matches_pb";
import { assets } from "@/lib/api";
import { cn } from "@/lib/utils";

type Props = {
  champions: { [key: string]: Champion };
  participant: MatchParticipant;
  patchVersion: string;
  className?: string;
};

export async function ChampionIcon(props: Props) {
  const { className, champions, participant, patchVersion } = props;

  const champ = Object.values(champions).find((c) => {
    return c.key == participant.championId.toString();
  });

  const squareAsset = await assets.getChampionSquareAssetUrl({
    patchVersion,
    championName: champ?.image?.full ?? "",
  });

  return (
    <HoverCard>
      <HoverCardTrigger>
        <Image
          alt={champ?.name ?? ""}
          width={champ?.image?.w ?? 0}
          height={champ?.image?.h ?? 0}
          src={squareAsset.url ?? ""}
          className={cn(
            "w-12 h-12 shrink-0 rounded-full object-center object-cover",
            className,
          )}
        />
      </HoverCardTrigger>
      <HoverCardContent>
        <p className="text-xs font-bold">{champ?.name}</p>
      </HoverCardContent>
    </HoverCard>
  );
}
