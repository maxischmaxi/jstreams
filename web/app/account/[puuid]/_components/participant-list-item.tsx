import Link from "next/link";
import { ParticipantChampionIcon } from "@/components/icons/participant-champion-icon";
import { MatchParticipant } from "@/matches/v1/matches_pb";
import { cn } from "@/lib/utils";

type Props = {
  patchVersion: string;
  participant: MatchParticipant;
  puuid: string;
};

export function ParticipantListItem(props: Props) {
  const { patchVersion, puuid, participant } = props;
  return (
    <li className="flex flex-row flex-nowrap gap-1">
      <ParticipantChampionIcon
        participant={participant}
        patchVersion={patchVersion}
      />
      <Link
        className={cn(
          participant.puuid !== puuid && "text-muted-foreground",
          "hover:text-foreground whitespace-nowrap",
        )}
        href={`/account/${participant.puuid}`}
      >
        {participant.summonerName}
      </Link>
    </li>
  );
}
