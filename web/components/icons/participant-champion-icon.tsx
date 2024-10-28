import "server-only";
import { assets, champions } from "@/lib/api";
import { MatchParticipant } from "@/matches/v1/matches_pb";
import Image from "next/image";
import { cn } from "@/lib/utils";

type Props = {
  participant: MatchParticipant;
  patchVersion: string;
  className?: string;
};

export async function ParticipantChampionIcon(props: Props) {
  const { className, patchVersion, participant } = props;

  const champ = await champions.getChampByChampId({
    patchVersion,
    champId: participant.championId,
  });

  const icon = await assets.getChampionSquareAssetUrl({
    patchVersion,
    championName: champ?.champion?.image?.full ?? "",
  });

  return (
    <Image
      width={16}
      height={16}
      src={icon?.url ?? ""}
      alt={participant.championName}
      className={cn(
        "w-4 h-4 object-center rounded-full object-cover",
        className,
      )}
    />
  );
}
