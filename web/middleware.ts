import { NextResponse } from "next/server";
import { version } from "./lib/api";

export const config = {
  matcher: [
    '/((?!api|_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)',
  ],
}

export async function middleware() {
  const v = await version.getCurrentVersion({});

  const h = new Headers();
  h.set("x-version", v.version);

  return NextResponse.next({
    headers: h,
    request: {
      headers: h
    }
  });
}
