import { masteries } from "@/lib/api";
import { getChampionMateriesByPuuidForChampionValidation } from "@/lib/riot/validation";
import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest): Promise<NextResponse> {
  const url = new URL(request.url);
  const puuid = url.searchParams.get("puuid");
  const region = url.searchParams.get("region");
  const championId = url.searchParams.get("championId");

  const validRequest =
    getChampionMateriesByPuuidForChampionValidation.safeParse({
      puuid,
      region,
      championId,
    });

  if (!validRequest.success) {
    return new NextResponse(JSON.stringify(validRequest.error), {
      status: 400,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  try {
    const { championMastery } =
      await masteries.getChampionMasteriesByPuuidByChampion({
        puuid: validRequest.data.puuid,
        region: validRequest.data.region,
        championId: validRequest.data.championId,
      });

    return NextResponse.json(championMastery);
  } catch (error) {
    return new NextResponse(JSON.stringify(error), {
      status: 500,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
}
