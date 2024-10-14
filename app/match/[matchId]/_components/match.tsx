"use client";

import { Match as MatchType } from "@/lib/riot/types";

type Props = {
  match: MatchType;
  timeline: any;
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
