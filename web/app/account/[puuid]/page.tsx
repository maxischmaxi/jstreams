import {
  assets,
  champions as championsApi,
  entries,
  masteries as masteriesApi,
  matches as matchesApi,
  summoner as summonerApi,
} from "@/lib/api";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { PlatformRoutingValues } from "@/summoner/v1/summoner_pb";
import { RegionalRoutingValues } from "@/account/v1/account_pb";
import { headers } from "next/headers";
import { MatchItem } from "./_components/match-item";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import dynamic from "next/dynamic";

const Champions = dynamic(
  () => import("./_components/champions").then((mod) => mod.Champions),
  {
    ssr: false,
  },
);

const RolesChart = dynamic(
  () => import("./_components/roles-chart").then((mod) => mod.RolesChart),
  {
    ssr: false,
  },
);

const Chart = dynamic(
  () =>
    import("./_components/recently-played-chart").then(
      (mod) => mod.RecentlyPlayerChart,
    ),
  {
    ssr: false,
  },
);

type Props = {
  params: {
    puuid: string;
  };
};

export type ChampionChartEntry = {
  icon: string;
  name: string;
  championLevel: number;
  championId: number;
  championPoints: number;
};

export default async function Page({ params: { puuid } }: Props) {
  const h = headers();
  const patchVersion = h.get("x-version") ?? "";
  const masteries = await masteriesApi.getChampionMasteriesByPuuid({
    puuid,
    region: PlatformRoutingValues.EUW1,
  });

  const matchIds = await matchesApi.getMatchIdsByPuuid({
    region: RegionalRoutingValues.EUROPE,
    puuid,
  });

  const { data: champions } = await championsApi.getChampions({
    patchVersion: patchVersion,
  });

  const summonerSpells = await summonerApi.getSummonerSpells({
    patchVersion: patchVersion,
  });

  const summoner = await summonerApi.getSummonerByPuuid({
    puuid,
    region: PlatformRoutingValues.EUW1,
  });

  const { entries: summonerEntries } = await entries.getEntriesBySummoner({
    region: PlatformRoutingValues.EUW1,
    summonerId: summoner.summoner?.id,
  });

  const matches = [];

  for (let i = 0; i < matchIds.matchIds.length; i++) {
    const match = await matchesApi.getMatchByMatchId({
      matchId: matchIds.matchIds[i],
      region: RegionalRoutingValues.EUROPE,
    });

    if (!match.match) {
      continue;
    }

    matches.push(match.match);
  }

  const winrate =
    Math.round(
      (summonerEntries[0].wins /
        (summonerEntries[0].wins + summonerEntries[0].losses)) *
        100,
    ) || 0;

  const numberOfWins = matches.filter(
    (m) => m.info?.participants.find((p) => p.puuid === puuid)?.win,
  ).length;

  const numberOfLosses = matches.filter(
    (m) => !m.info?.participants.find((p) => p.puuid === puuid)?.win,
  ).length;

  const { topGames, jungleGames, midGames, bottomGames, supporGames } =
    matches.reduce(
      (acc, match) => {
        const participant = match.info?.participants.find(
          (p) => p.puuid === puuid,
        );

        if (!participant) {
          return acc;
        }

        switch (participant.lane) {
          case "TOP":
            acc.topGames++;
            break;
          case "JUNGLE":
            acc.jungleGames++;
            break;
          case "MID":
            acc.midGames++;
            break;
          case "BOTTOM":
            acc.bottomGames++;
            break;
          case "SUPPORT":
            acc.supporGames++;
            break;
        }

        return acc;
      },
      {
        topGames: 0,
        jungleGames: 0,
        midGames: 0,
        bottomGames: 0,
        supporGames: 0,
      },
    );

  const data: ChampionChartEntry[] = [];

  for (let i = 0; i < masteries.championMasteries.length; i++) {
    const entry = masteries.championMasteries[i];

    const champ = await championsApi.getChampByChampId({
      patchVersion,
      champId: entry.championId,
    });

    const icon = await assets.getChampionSquareAssetUrl({
      patchVersion,
      championName: champ?.champion?.image?.full ?? "",
    });

    data.push({
      icon: icon?.url ?? "",
      name: champ?.champion?.name ?? "",
      championLevel: entry.championLevel,
      championId: entry.championId,
      championPoints: entry.championPoints,
    });
  }

  return (
    <Tabs defaultValue="overview">
      <TabsList className="w-full justify-start">
        <TabsTrigger value="overview">Overview</TabsTrigger>
        <TabsTrigger value="champions">Champions</TabsTrigger>
        <TabsTrigger value="matchhistory">Match History</TabsTrigger>
      </TabsList>
      <TabsContent value="overview">
        <div className="flex flex-row flex-nowrap w-full gap-2">
          <div className="w-full min-w-[300px] max-w-[300px]">
            <Card>
              <CardHeader className="py-4 border-b border-secondary">
                <CardTitle>Solo/Duo (Rangliste)</CardTitle>
              </CardHeader>
              <CardContent className="pt-4">
                <div className="flex flex-row flex-nowrap justify-between w-full items-center">
                  <div>
                    <p>
                      {summonerEntries[0].tier.charAt(0) +
                        summonerEntries[0].tier.slice(1).toLowerCase()}{" "}
                      {summonerEntries[0].rank}
                    </p>
                    <p className="text-xs text-muted-foreground">
                      {summonerEntries[0].leaguePoints} LP
                    </p>
                  </div>

                  <div className="text-xs text-muted-foreground">
                    <p>
                      {summonerEntries[0].wins}W / {summonerEntries[0].losses}L
                    </p>
                    <p>Winrate {winrate} %</p>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
          <div className="w-full">
            <Tabs defaultValue="all">
              <TabsList>
                <TabsTrigger value="all">All</TabsTrigger>
                <TabsTrigger value="soloq">Solo/Duo (Rangliste)</TabsTrigger>
                <TabsTrigger value="flexq">Flexi (Rangliste)</TabsTrigger>
                <TabsTrigger value="aram">ARAM</TabsTrigger>
              </TabsList>
              <TabsContent className="space-y-2" value="all">
                <Card>
                  <CardHeader className="py-4 border-b border-secondary">
                    <CardTitle>
                      Kürzlich gespielt{" "}
                      <span className="text-sm text-muted-foreground">
                        ({matches.length} Games)
                      </span>
                    </CardTitle>
                  </CardHeader>
                  <CardContent className="pt-4 max-h-[200px]">
                    <div className="flex flex-row flex-nowrap justify-between items-center gap-4">
                      <div className="w-[150px] h-[150px] shrink-0">
                        <p className="text-xs text-muted-foreground text-center w-full">
                          {matches.length} G / {numberOfWins} W /{" "}
                          {numberOfLosses} L
                        </p>
                        <Chart wins={numberOfWins} losses={numberOfLosses} />
                      </div>
                      <div className="max-w-[300px] w-full h-full">
                        <RolesChart
                          data={[
                            { role: "top", games: topGames },
                            { role: "jungle", games: jungleGames },
                            { role: "mid", games: midGames },
                            { role: "bottom", games: bottomGames },
                            { role: "support", games: supporGames },
                          ]}
                        />
                      </div>
                    </div>
                  </CardContent>
                </Card>
                <ul className="w-full flex flex-col gap-2">
                  {matches.map((match, key) => (
                    <MatchItem
                      patchVersion={patchVersion}
                      key={key}
                      match={match}
                      summonerSpells={summonerSpells.data}
                      puuid={puuid}
                      champions={champions}
                    />
                  ))}
                </ul>
              </TabsContent>
              <TabsContent value="soloq">
                <ul className="w-full flex flex-col gap-2">
                  {matches.map((match, key) => (
                    <MatchItem
                      patchVersion={patchVersion}
                      key={key}
                      match={match}
                      summonerSpells={summonerSpells.data}
                      puuid={puuid}
                      champions={champions}
                    />
                  ))}
                </ul>
              </TabsContent>
              <TabsContent value="flexa"></TabsContent>
              <TabsContent value="aram"></TabsContent>
            </Tabs>
          </div>
        </div>
      </TabsContent>
      <TabsContent value="champions">
        <Champions data={data} />
      </TabsContent>
      <TabsContent value="matchhistory"></TabsContent>
    </Tabs>
  );
}
