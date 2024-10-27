"use client";

import { useUser } from "@auth0/nextjs-auth0/client";
import { LoaderCircle } from "lucide-react";
import { ReactNode } from "react";

type Props = {
  children?: ReactNode;
};

export function UserLoading({ children }: Props) {
  const user = useUser();

  if (user.isLoading) {
    return (
      <div className="flex items-center justify-center h-[theme(spacing.16)]">
        <LoaderCircle className="w-8 h-8 animate-spin" />
      </div>
    );
  }

  return children;
}
