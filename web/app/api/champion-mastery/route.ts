import { masteries } from "@/lib/api";
import { getChampionMasteriesByPuuidValidation } from "@/lib/riot/validation";
import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest): Promise<NextResponse> {
  const url = new URL(request.url);
  const puuid = url.searchParams.get("puuid");
  const region = url.searchParams.get("region");

  const validRequest = getChampionMasteriesByPuuidValidation.safeParse({
    puuid,
    region,
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
    const { championMasteries } = await masteries.getChampionMasteriesByPuuid({
      puuid: validRequest.data.puuid,
      region: validRequest.data.region,
    });

    return NextResponse.json(championMasteries);
  } catch (error) {
    return new NextResponse(JSON.stringify(error), {
      status: 500,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
}
