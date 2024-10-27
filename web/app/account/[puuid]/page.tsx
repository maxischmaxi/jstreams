import { Match } from "./_components/match";
import {
  champions as championsApi,
  masteries as masteriesApi,
  matches as matchesApi,
  summoner as summonerApi,
} from "@/lib/api";
import { Match as MatchType } from "@maxischmaxi/jstreams-ts/matches/v1/matches_pb";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Champions } from "./_components/champions";
import { PlatformRoutingValues } from "@maxischmaxi/jstreams-ts/summoner/v1/summoner_pb";
import { RegionalRoutingValues } from "@maxischmaxi/jstreams-ts/account/v1/account_pb";

type Props = {
  params: {
    puuid: string;
  };
};

export default async function Page({ params: { puuid } }: Props) {
  const masteries = await masteriesApi.getChampionMasteriesByPuuid({
    puuid,
    region: PlatformRoutingValues.EUW1,
  });

  const matchIds = await matchesApi.getMatchIdsByPuuid({
    region: RegionalRoutingValues.EUROPE,
    puuid,
  });

  const matches: MatchType[] = [];

  for (let i = 0; i < matchIds.matchIds.length; i++) {
    const match = await matchesApi
      .getMatchByMatchId({
        matchId: matchIds.matchIds[i],
        region: RegionalRoutingValues.EUROPE,
      })
      .then((res) => res.match);

    if (match) {
      matches.push(match);
    }
  }

  const { data: champions } = await championsApi.getChampions({
    patchVersion: "14.20.1",
  });

  const summonerSpells = await summonerApi.getSummonerSpells({
    patchVersion: "14.20.1",
  });

  return (
    <Tabs defaultValue="champions">
      <TabsList>
        <TabsTrigger value="champions">Champions</TabsTrigger>
        <TabsTrigger value="matchhistory">Match History</TabsTrigger>
      </TabsList>
      <TabsContent value="champions">
        <Champions masteries={masteries} />
      </TabsContent>
      <TabsContent value="matchhistory">
        <ul className="w-full flex flex-col gap-4">
          {matches.map((match, key) => (
            <Match
              key={key}
              summonerSpells={summonerSpells.data}
              match={match}
              puuid={puuid}
              champions={champions}
            />
          ))}
        </ul>
      </TabsContent>
    </Tabs>
  );
}
