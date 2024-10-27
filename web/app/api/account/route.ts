import { account as accountApi } from "@/lib/api";
import { accountRegionToTagline } from "@/lib/riot/lib";
import {
  getAccountByGamenameAndTaglineResponseValidation,
  getAccountByGamenameAndTaglineValidation,
} from "@/lib/riot/validation";
import { RegionalRoutingValues } from "@maxischmaxi/jstreams-ts/account_pb";
import { NextRequest, NextResponse } from "next/server";

async function getAccountByGamenameAndTaglineHandler({
  gameName,
  tagLine,
  region,
}: {
  gameName: string;
  tagLine: string;
  region: RegionalRoutingValues;
}): Promise<NextResponse> {
  try {
    const { account: accountData } =
      await accountApi.getAccountByGamenameAndTagline({
        region,
        gamename: gameName,
        tagline: tagLine,
      });

    const responseValid =
      getAccountByGamenameAndTaglineResponseValidation.safeParse(accountData);

    if (!responseValid.success) {
      return new NextResponse(JSON.stringify(responseValid.error), {
        status: 500,
        headers: {
          "Content-Type": "application/json",
        },
      });
    }

    return NextResponse.json(accountData);
  } catch (error) {
    return new NextResponse(JSON.stringify(error), {
      status: 500,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
}

async function getAccountByPuuidHandler(puuid: string): Promise<NextResponse> {
  try {
    const account = await accountApi.getAccountByPuuid({
      puuid,
      region: RegionalRoutingValues.EUROPE,
    });

    const responseValid =
      getAccountByGamenameAndTaglineResponseValidation.safeParse(account);

    if (!responseValid.success) {
      return new NextResponse(JSON.stringify(responseValid.error), {
        status: 500,
        headers: {
          "Content-Type": "application/json",
        },
      });
    }

    return NextResponse.json(account);
  } catch (error) {
    return new NextResponse(JSON.stringify(error), {
      status: 500,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
}

export async function GET(request: NextRequest): Promise<NextResponse> {
  const url = new URL(request.url);
  const gameName = url.searchParams.get("gameName");
  const tagLine = url.searchParams.get("tagLine");
  const puuid = url.searchParams.get("puuid");
  const region = url.searchParams.get("region");
  const intRegion = region ? parseInt(region) : null;

  const valid = getAccountByGamenameAndTaglineValidation.safeParse({
    gameName,
    tagLine,
    region: intRegion,
  });

  if (valid.success) {
    return getAccountByGamenameAndTaglineHandler({
      gameName: valid.data.gameName,
      region: valid.data.region,
      tagLine: accountRegionToTagline(valid.data.region),
    });
  }

  if (puuid) {
    return getAccountByPuuidHandler(puuid);
  }

  return new NextResponse(
    JSON.stringify({
      error: "Invalid request",
    }),
    {
      status: 400,
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
}
