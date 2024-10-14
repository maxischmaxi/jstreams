import { GetAccountByGamenameAndTaglineResponse } from "@/lib/riot/types";
import {
  REGIONAL_ROUTING_VALUES,
  accountRegionToTagline,
} from "@/lib/riot/lib";
import { getAccountByGamenameAndTaglineResponseValidation } from "@/lib/riot/validation";
import { useQuery, UseQueryResult } from "@tanstack/react-query";

type UseAccountByGamenameAndTaglineProps = {
  gameName: string;
  region: REGIONAL_ROUTING_VALUES;
};

export async function getAccountByGamenameAndTagline(
  props: UseAccountByGamenameAndTaglineProps,
): Promise<GetAccountByGamenameAndTaglineResponse> {
  const { gameName, region } = props;
  const tagLine = accountRegionToTagline(region);
  const url = `/api/account?gameName=${gameName}&tagLine=${tagLine}&region=${region}`;
  const response = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const data = await response.json();
  const isValid =
    getAccountByGamenameAndTaglineResponseValidation.safeParse(data);

  if (!isValid.success) {
    throw new TypeError(isValid.error.message);
  }

  return data;
}

export function useQueryAccountByGamenameAndTagline(
  props: UseAccountByGamenameAndTaglineProps,
): UseQueryResult<GetAccountByGamenameAndTaglineResponse, TypeError> {
  const { gameName, region } = props;

  return useQuery({
    queryKey: ["useAccountByGamenameAndTagline", gameName, region],
    enabled: false,
    async queryFn() {
      return getAccountByGamenameAndTagline(props);
    },
  });
}
