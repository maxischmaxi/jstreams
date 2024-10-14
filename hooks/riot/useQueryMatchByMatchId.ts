import { REGIONAL_ROUTING_VALUES } from "@/lib/riot/lib";
import { GetMatchByMatchIdResponse } from "@/lib/riot/types";
import { useQuery, UseQueryResult } from "@tanstack/react-query";

type Props = {
  matchId: string;
  region: REGIONAL_ROUTING_VALUES;
};

export function useQueryMatchByMatchId(
  props: Props,
): UseQueryResult<GetMatchByMatchIdResponse, TypeError> {
  const { matchId, region } = props;

  return useQuery({
    queryKey: ["useMatchByMatchId", matchId, region],
    enabled: true,
    async queryFn() {
      const url = new URL("/api/match");
      url.searchParams.set("matchId", matchId);
      url.searchParams.set("region", region);

      const res = await fetch(url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });

      return await res.json();
    },
  });
}
