import { useQuery, UseQueryResult } from "@tanstack/react-query";
import account from "@/lib/api/account";
import {
  GetAccountByGamenameAndTaglineRequest,
  GetAccountByGamenameAndTaglineResponse,
} from "@maxischmaxi/jstreams-ts/gen/account_pb";

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
