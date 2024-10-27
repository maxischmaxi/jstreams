import { masteries } from "@/lib/api";
import { PlatformRoutingValues } from "@/summoner/v1/summoner_pb";

type Props = {
  params: {
    puuid: string;
    championId: number;
  };
};

export default async function Page({ params: { puuid, championId } }: Props) {
  const mastery = await masteries.getChampionMasteriesByPuuidByChampion({
    puuid,
    championId,
    region: PlatformRoutingValues.EUW1,
  });

  return (
    <div>
      <p>{mastery.championMastery?.lastPlayTime}</p>
    </div>
  );
}
