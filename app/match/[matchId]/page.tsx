import { REGIONAL_ROUTING_VALUES } from "@/lib/riot/lib";
import { getMatchByMatchId, getMatchTimeline } from "@/lib/riot/match_v5";
import { Match } from "./_components/match";

type Props = {
  params: {
    matchId: string;
  };
};

export default async function Page({ params: { matchId } }: Props) {
  const match = await getMatchByMatchId({
    matchId,
    region: REGIONAL_ROUTING_VALUES.EUROPE,
  });

  const timeline = await getMatchTimeline({
    matchId,
    region: REGIONAL_ROUTING_VALUES.EUROPE,
  });

  return (
    <div>
      <p>{match.metadata.matchId}</p>
      <Match match={match} timeline={timeline} />
    </div>
  );
}
