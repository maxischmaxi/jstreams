import { matches } from "@/lib/api";
import { getMatchByMatchIdValidation } from "@/lib/riot/validation";
import { RegionalRoutingValues } from "@maxischmaxi/jstreams-ts/gen/account_pb";
import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest): Promise<NextResponse> {
  const url = new URL(request.url);
  const matchId = url.searchParams.get("matchId");
  const region = url.searchParams.get("region");

  const isValid = getMatchByMatchIdValidation.safeParse({
    matchId,
    region,
  });

  if (!isValid.success) {
    return new NextResponse(JSON.stringify(isValid.error), {
      status: 400,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  try {
    const match = await matches.getMatchByMatchId({
      region: RegionalRoutingValues.EUROPE,
      matchId: isValid.data.matchId,
    });

    return NextResponse.json(match);
  } catch (error) {
    return new NextResponse(JSON.stringify(error), {
      status: 500,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
}
