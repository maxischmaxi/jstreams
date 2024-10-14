import { getChampionMateriesByPuuid } from "@/lib/riot/champion_mastery_v4";
import {
  REGIONAL_ROUTING_VALUES,
  PLATFORM_ROUTING_VALUES,
} from "@/lib/riot/lib";
import { getMatchByMatchId, getMatchIdsByPuuid } from "@/lib/riot/match_v5";
import { Match } from "./_components/match";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Champions } from "./_components/champions";
import { Match as MatchType } from "@/lib/riot/types/match/match";
import { getChampions } from "@/lib/riot/champion";
import { getSummonerSpells } from "@/lib/riot/summonerspell";

type Props = {
  params: {
    puuid: string;
  };
};

export default async function Page({ params: { puuid } }: Props) {
  const masteries = await getChampionMateriesByPuuid({
    puuid,
    region: PLATFORM_ROUTING_VALUES.EUW1,
  });

  const matchIds = await getMatchIdsByPuuid({
    region: REGIONAL_ROUTING_VALUES.EUROPE,
    puuid,
  });

  const matches: MatchType[] = [];

  for (let i = 0; i < matchIds.length; i++) {
    matches.push(
      await getMatchByMatchId({
        matchId: matchIds[i],
        region: REGIONAL_ROUTING_VALUES.EUROPE,
      }),
    );
  }

  const { data: champions } = await getChampions("14.20.1");

  const summonerSpells = await getSummonerSpells("14.20.1");

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
