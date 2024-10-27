import {
  champions as championsApi,
  masteries as masteriesApi,
  matches as matchesApi,
  summoner as summonerApi,
} from "@/lib/api";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Champions } from "./_components/champions";
import { PlatformRoutingValues } from "@/summoner/v1/summoner_pb";
import { RegionalRoutingValues } from "@/account/v1/account_pb";
import { PlainMessage } from "@bufbuild/protobuf";
import { ChampionMastery } from "@/masteries/v1/masteries_pb";
import { headers } from "next/headers";
import { ServerMatch } from "./_components/servermatch";

type Props = {
  params: {
    puuid: string;
  };
};

export default async function Page({ params: { puuid } }: Props) {
  const h = headers();
  const version = h.get("x-version") ?? "";
  const masteries = await masteriesApi.getChampionMasteriesByPuuid({
    puuid,
    region: PlatformRoutingValues.EUW1,
  });

  const matchIds = await matchesApi.getMatchIdsByPuuid({
    region: RegionalRoutingValues.EUROPE,
    puuid,
  });

  const { data: champions } = await championsApi.getChampions({
    patchVersion: version,
  });

  const summonerSpells = await summonerApi.getSummonerSpells({
    patchVersion: version,
  });

  const championMasteries = masteries.championMasteries.map((m) => m.toJson() as unknown as PlainMessage<ChampionMastery>)

  return (
    <Tabs defaultValue="champions">
      <TabsList>
        <TabsTrigger value="champions">Champions</TabsTrigger>
        <TabsTrigger value="matchhistory">Match History</TabsTrigger>
      </TabsList>
      <TabsContent value="champions">
        <Champions masteries={championMasteries} />
      </TabsContent>
      <TabsContent value="matchhistory">
        <ul className="w-full flex flex-col gap-4">
          {matchIds.matchIds.map((match, key) => (
            <ServerMatch
              key={key}
              matchId={match}
              summonerSpells={summonerSpells.data}
              puuid={puuid}
              champions={champions}
            />
          ))}
        </ul>
      </TabsContent>
    </Tabs>
  );
}
