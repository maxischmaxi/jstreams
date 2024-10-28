import { RegionalRoutingValues } from "@/account/v1/account_pb";
import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

type SearchHistoryItem = {
  gameName: string;
  region: RegionalRoutingValues;
  timestamp: number;
};

type Store = {
  items: SearchHistoryItem[];
  latest: SearchHistoryItem | null;
  addItem: (item: SearchHistoryItem) => void;
  removeItem: (index: number) => void;
};

export const useSummonerSearchHistory = create(
  persist<Store>(
    (set, get) => ({
      items: [],
      latest: null,
      addItem: (item) => {
        const { items } = get();
        const index = items.findIndex(
          (i) => i.gameName === item.gameName && i.region === item.region,
        );
        if (index === -1) {
          set((state) => ({ items: [item, ...state.items], latest: item }));
          return;
        }

        const newItems = structuredClone(items);
        newItems.splice(index, 1);
        newItems.unshift(item);
        set({ items: newItems, latest: item });
      },
      removeItem: (index) => {
        const { items } = get();
        const newItems = structuredClone(items);
        newItems.splice(index, 1);
        set({
          items: newItems,
          latest: get().items.reduce((a, b) =>
            a.timestamp > b.timestamp ? a : b,
          ),
        });
      },
    }),
    {
      name: "summoner-search-history",
      storage: createJSONStorage(() => localStorage),
    },
  ),
);
