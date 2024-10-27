import { getChampionMasteriesByPuuidByChampion } from "@/lib/riot/champion_mastery_v4";
import { PLATFORM_ROUTING_VALUES } from "@/lib/riot/lib";

type Props = {
  params: {
    puuid: string;
    championId: number;
  };
};

export default async function Page({ params: { puuid, championId } }: Props) {
  const mastery = await getChampionMasteriesByPuuidByChampion({
    puuid,
    championId,
    region: PLATFORM_ROUTING_VALUES.EUW1,
  });

  return (
    <div>
      <p>{mastery.lastPlayTime}</p>
    </div>
  );
}
