"use client";
import { LoaderCircle, Menu, Package2 } from "lucide-react";
import Link from "next/link";
import { Sheet, SheetContent, SheetTrigger } from "../ui/sheet";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { ModeToggle } from "../theme/mode-toggle";
import { useUser } from "@auth0/nextjs-auth0/client";
import { usePathname, useRouter } from "next/navigation";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { cn } from "@/lib/utils";
import { ReactNode } from "react";

type DesktopLinkProps = {
  href: string;
  children: ReactNode;
};

function DesktopLink(props: DesktopLinkProps) {
  const pathname = usePathname();
  return (
    <Link
      href={props.href}
      className={cn(
        pathname === props.href ? "text-foreground" : "text-muted-foreground",
        "transition-colors hover:text-foreground",
      )}
    >
      {props.children}
    </Link>
  );
}

export function Header() {
  const user = useUser();
  const router = useRouter();

  function handleSignOut() {
    router.push("/api/auth/logout");
  }

  return (
    <header className="z-10 sticky top-0 flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6">
      <nav className="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
        <Link
          href="/"
          className="flex items-center gap-2 text-lg font-semibold md:text-base"
        >
          <Package2 className="h-6 w-6" />
          <span className="sr-only">Acme Inc</span>
        </Link>
        <DesktopLink href="/">Dashboard</DesktopLink>
      </nav>
      <Sheet>
        <SheetTrigger asChild>
          <Button variant="outline" size="icon" className="shrink-0 md:hidden">
            <Menu className="h-5 w-5" />
            <span className="sr-only">Toggle navigation menu</span>
          </Button>
        </SheetTrigger>
        <SheetContent side="left">
          <nav className="grid gap-6 text-lg font-medium">
            <Link
              href="/"
              className="flex items-center gap-2 text-lg font-semibold"
            >
              <Package2 className="h-6 w-6" />
              <span className="sr-only">Acme Inc</span>
            </Link>
            <Link
              href="/"
              className="text-muted-foreground hover:text-foreground"
            >
              Dashboard
            </Link>
            <Link
              href="#"
              className="text-muted-foreground hover:text-foreground"
            >
              Orders
            </Link>
            <Link
              href="#"
              className="text-muted-foreground hover:text-foreground"
            >
              Products
            </Link>
            <Link
              href="#"
              className="text-muted-foreground hover:text-foreground"
            >
              Customers
            </Link>
          </nav>
        </SheetContent>
      </Sheet>
      <div className="flex w-full justify-end items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
        {user.isLoading && (
          <LoaderCircle className="h-7 w-7 text-muted-foreground animate-spin" />
        )}
        {!user.isLoading && user.user && (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="secondary" size="icon" className="rounded-full">
                <Avatar className="h-7 w-7">
                  <AvatarFallback />
                  <AvatarImage
                    src={user.user.picture ?? undefined}
                    alt="User avatar"
                  />
                </Avatar>
                <span className="sr-only">Toggle user menu</span>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>
                {user.user.name ?? user.user.email ?? "-"}
              </DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem>Settings</DropdownMenuItem>
              <DropdownMenuItem>Support</DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem onClick={handleSignOut}>
                Logout
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        )}
        {!user.isLoading && !user.user && (
          <div className="text-lg md:text-sm font-medium">
            <Link
              className="text-muted-foreground transition-colors hover:text-foreground"
              href="/api/auth/login"
            >
              Einloggen
            </Link>
          </div>
        )}
        <ModeToggle />
      </div>
    </header>
  );
}
