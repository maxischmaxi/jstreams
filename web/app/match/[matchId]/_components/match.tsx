"use client";

import { Match as MatchType } from "@/lib/riot/types/match/match";

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
