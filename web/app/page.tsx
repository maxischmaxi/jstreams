"use client";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import { useSummonerSearchHistory } from "@/hooks/useSummonerSearchHistory";
import { useRouter } from "next/navigation";
import dayjs from "dayjs";
import { useForm, useWatch } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { UserLoading } from "@/components/layout/user-loading";
import { getAccountByGamenameAndTaglineValidation } from "@/lib/riot/validation";
import { ChevronDown } from "lucide-react";
import {
  Account,
  RegionalRoutingValues,
} from "@maxischmaxi/jstreams-ts/account_pb";
import { accountRegionToString, accountRegionToTagline } from "@/lib/riot/lib";

export default function Page() {
  const router = useRouter();
  const history = useSummonerSearchHistory();

  const form = useForm<
    z.infer<typeof getAccountByGamenameAndTaglineValidation>
  >({
    resolver: zodResolver(getAccountByGamenameAndTaglineValidation),
    defaultValues: {
      gameName: useSummonerSearchHistory.getState().latest?.gameName ?? "",
      region:
        useSummonerSearchHistory.getState().latest?.region ??
        RegionalRoutingValues.EUROPE,
    },
  });

  const gameName = useWatch({ control: form.control, name: "gameName" });
  const region = useWatch({ control: form.control, name: "region" });

  async function handleSubmit(
    data: z.infer<typeof getAccountByGamenameAndTaglineValidation>,
  ) {
    const response = await fetch(
      "/api/account?gameName=" +
        data.gameName +
        "&region=" +
        RegionalRoutingValues.EUROPE +
        "&tagLine=" +
        accountRegionToTagline(RegionalRoutingValues.EUROPE),
    );
    const account = (await response.json()) as Account;
    history.addItem({ gameName, region, timestamp: dayjs().unix() });
    router.push(`/account/${account.puuid}`);
  }

  return (
    <UserLoading>
      <div className="flex flex-row justify-center items-center">
        <form
          onSubmit={form.handleSubmit(handleSubmit)}
          className="flex flex-row flex-nowrap gap-4 bg-card rounded-full py-4 px-12"
        >
          <Input
            {...form.register("gameName")}
            placeholder="Summoner Name"
            className="w-full max-w-[12rem] border-none rounded-none outline-none shadow-none focus-visible:ring-0 px-0"
          />
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                size="icon"
                className="px-4 w-auto border-none hover:bg-background"
              >
                {accountRegionToString(region)}
                <ChevronDown className="w-3 h-3" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuItem
                onClick={() =>
                  form.setValue("region", RegionalRoutingValues.EUROPE)
                }
              >
                {accountRegionToString(RegionalRoutingValues.EUROPE)}
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
          <Button type="submit">Search</Button>
        </form>
      </div>
    </UserLoading>
  );
}
