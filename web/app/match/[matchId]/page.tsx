import { RegionalRoutingValues } from "@maxischmaxi/jstreams-ts/account/v1/account_pb";
import { Match } from "./_components/match";
import { matches } from "@/lib/api";

type Props = {
  params: {
    matchId: string;
  };
};

export default async function Page({ params: { matchId } }: Props) {
  const match = await matches.getMatchByMatchId({
    matchId,
    region: RegionalRoutingValues.EUROPE,
  });

  const timeline = await matches.getMatchTimeline({
    matchId,
    region: RegionalRoutingValues.EUROPE,
  });

  return (
    <div>
      <p>{match.match?.metadata?.matchId}</p>
      <Match match={match} timeline={timeline} />
    </div>
  );
}
