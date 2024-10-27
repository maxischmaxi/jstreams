import { GetAccountByGamenameAndTaglineRequest, GetAccountByGamenameAndTaglineResponse } from "@/account/v1/account_pb";
import { account } from "@/lib/api";
import { useQuery, UseQueryResult } from "@tanstack/react-query";

export function useQueryAccountByGamenameAndTagline(
  props: GetAccountByGamenameAndTaglineRequest,
): UseQueryResult<GetAccountByGamenameAndTaglineResponse, TypeError> {
  return useQuery({
    queryKey: ["useAccountByGamenameAndTagline", props.gamename, props.region],
    enabled: false,
    async queryFn() {
      return account.getAccountByGamenameAndTagline(props);
    },
  });
}
