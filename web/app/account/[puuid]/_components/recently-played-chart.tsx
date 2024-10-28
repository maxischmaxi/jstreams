"use client";
import { ChartConfig, ChartContainer } from "@/components/ui/chart";
import { Cell, Label, Pie, PieChart } from "recharts";

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
  wins: number;
  losses: number;
};

export function RecentlyPlayerChart(props: Props) {
  const data = [
    {
      name: "Wins",
      value: props.wins,
    },
    {
      name: "Losses",
      value: props.losses,
    },
  ];

  const winrate = Math.round((props.wins / (props.wins + props.losses)) * 100);

  return (
    <ChartContainer className="h-full w-[150px]" config={chartConfig}>
      <PieChart accessibilityLayer data={data}>
        <Pie
          data={data}
          dataKey="value"
          fill="var(--color-desktop)"
          cy="50%"
          cx="50%"
          innerRadius={40}
          outerRadius={60}
        >
          {data.map((entry, index) => (
            <Cell
              key={`cell-${index}`}
              fill={
                entry.name === "Wins"
                  ? "var(--win-foreground)"
                  : "var(--lose-foreground)"
              }
            />
          ))}
          <Label
            content={({ x, y, value }) => (
              <g>
                <text
                  x={x}
                  y={y}
                  textAnchor="middle"
                  dominantBaseline="central"
                  alignmentBaseline="middle"
                  fontSize="14"
                  color="var(--win-foreground)"
                  fontWeight="bold"
                >
                  {value}
                </text>
              </g>
            )}
            value={`${winrate}%`}
            position="center"
          />
        </Pie>
      </PieChart>
    </ChartContainer>
  );
}
