import { summoner as summonerApi, account as accountApi } from "@/lib/api";
import { RegionalRoutingValues } from "@/account/v1/account_pb";
import { PlatformRoutingValues } from "@/summoner/v1/summoner_pb";
import Image from "next/image";
import { ReactNode } from "react";
import { headers } from "next/headers";

type Props = {
  params: {
    puuid: string;
  };
  children: ReactNode;
};

export default async function Layout({ children, params: { puuid } }: Props) {
  const h = headers();
  const v = h.get("x-version") ?? "";

  const summoner = await summonerApi.getSummonerByPuuid({
    puuid,
    region: PlatformRoutingValues.EUW1,
  });

  const account = await accountApi.getAccountByPuuid({
    puuid,
    region: RegionalRoutingValues.EUROPE,
  });

  const icon = await accountApi.getAccountProfileIcon({
    patchVersion: v,
    profileIconId: summoner.summoner?.profileIconId,
  });

  return (
    <div className="container mx-auto flex flex-col flex-nowrap gap-8">
      <div className="flex flex-row flex-nowrap gap-4">
        <div className="relative">
          <Image
            src={icon.url}
            priority
            width={128}
            height={128}
            alt={`${summoner.summoner?.accountId}-profile-icon`}
            className="rounded-md object-center object-cover"
          />
          <span className="absolute left-1/2 bottom-0 -translate-x-1/2 translate-y-1/2 text-white bg-gray-800 text-xs px-2 rounded-full">
            {summoner.summoner?.summonerLevel}
          </span>
        </div>
        <div className="flex flex-col">
          <h1 className="text-xl font-bold">{account.account?.gameName}</h1>
          <h4 className="text-primary text-xs">{account.account?.tagLine}</h4>
        </div>
        <div className="flex flex-row flex-nowrap justify-start items-center"></div>
      </div>
      {children}
    </div>
  );
}
