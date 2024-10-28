"use client";

import { Match as MatchType } from "@/matches/v1/matches_pb";

type Props = {
  match: MatchType;
  timeline: object;
};

export function Match(props: Props) {
  return (
    <button
      onClick={() => console.log(JSON.stringify(props.timeline, null, 2))}
    >
      log
    </button>
  );
}
