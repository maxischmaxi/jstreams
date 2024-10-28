import "server-only";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "../ui/hover-card";
import Image from "next/image";
import { items } from "@/lib/api";
import { cn } from "@/lib/utils";

type Props = {
  itemId: number;
  patchVersion: string;
  className?: string;
};

export async function ItemIcon(props: Props) {
  const { itemId, patchVersion, className } = props;

  const item = await items.getItemImageById({
    patchVersion,
    itemId,
  });

  const info = Boolean(item.url)
    ? await items.getItemInformationById({
        patchVersion,
        itemId,
      })
    : null;

  return (
    <>
      {Boolean(item.url) ? (
        <HoverCard>
          <HoverCardTrigger asChild>
            <Image
              src={item.url}
              width={48}
              height={48}
              alt={info?.item?.name ?? "Item Image"}
              className={cn("w-6 h-6 object-center object-cover", className)}
            />
          </HoverCardTrigger>
          <HoverCardContent>
            <p className="text-xs font-bold">{info?.item?.name}</p>
            <ul className="text-2xs">
              <li>{info?.item?.gold?.total} Gold</li>
              <li>{info?.item?.plaintext}</li>
            </ul>
          </HoverCardContent>
        </HoverCard>
      ) : (
        <div className="w-6 h-6 bg-secondary" />
      )}
    </>
  );
}
