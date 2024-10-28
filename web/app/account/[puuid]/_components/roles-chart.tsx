"use client";

import { ChartConfig, ChartContainer } from "@/components/ui/chart";
import { Bar, BarChart, LabelList, Tooltip } from "recharts";
import top from "@/assets/roles/icon-position-top.svg";
import jungle from "@/assets/roles/icon-position-jungle.svg";
import mid from "@/assets/roles/icon-position-mid.svg";
import bottom from "@/assets/roles/icon-position-adc.svg";
import support from "@/assets/roles/icon-position-support.svg";

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#2563eb",
  },
  mobile: {
    label: "Mobile",
    color: "#60a5fa",
  },
} satisfies ChartConfig;

type Props = {
  data: {
    role: string;
    games: number;
  }[];
};

export function RolesChart(props: Props) {
  return (
    <ChartContainer
      config={chartConfig}
      className="w-full h-full max-h-[180px]"
    >
      <BarChart accessibilityLayer data={props.data} margin={{ top: 25 }}>
        <Bar
          fill="var(--win-foreground)"
          dataKey="games"
          background={{ fill: "hsl(var(--muted))" }}
        >
          <Tooltip />
          <LabelList
            dataKey="role"
            position="top"
            content={({ value, width, x }) => {
              let newX: number;
              const newY = 0;
              let newWidth: number;

              if (typeof x === "string") {
                newX = parseInt(x);
              } else if (typeof x === "number") {
                newX = x;
              } else {
                newX = 0;
              }

              if (typeof width === "string") {
                newWidth = parseInt(width);
              } else if (typeof width === "number") {
                newWidth = width;
              } else {
                newWidth = 0;
              }

              newX = newWidth / 2 + newX - 10;

              switch (value) {
                case "top":
                  return (
                    <g>
                      <image
                        href={top.src}
                        x={newX}
                        y={newY}
                        width={20}
                        height={20}
                      />
                    </g>
                  );
                case "jungle":
                  return (
                    <g>
                      <image
                        href={jungle.src}
                        x={newX}
                        y={newY}
                        width={20}
                        height={20}
                      />
                    </g>
                  );
                case "mid":
                  return (
                    <g>
                      <image
                        href={mid.src}
                        x={newX}
                        y={newY}
                        width={20}
                        height={20}
                      />
                    </g>
                  );
                case "bottom":
                  return (
                    <g>
                      <image
                        href={bottom.src}
                        x={newX}
                        y={newY}
                        width={20}
                        height={20}
                      />
                    </g>
                  );
                case "support":
                  return (
                    <g>
                      <image
                        href={support.src}
                        x={newX}
                        y={newY}
                        width={20}
                        height={20}
                      />
                    </g>
                  );
                default:
                  return (
                    <g>
                      <text fill="#fff" x={x} y={y}>
                        {value}
                      </text>
                    </g>
                  );
              }
            }}
          />
        </Bar>
      </BarChart>
    </ChartContainer>
  );
}
