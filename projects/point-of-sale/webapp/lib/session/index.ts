"use server";
import { cookies } from "next/headers";

export async function createSession(jwtToken: string) {
  const expiresAt = new Date(Date.now() + 7 * 24 * 60 * 60 * 1000);

  cookies().set("session", jwtToken, {
    httpOnly: true,
    secure: true,
    expires: expiresAt,
    sameSite: "lax",
    path: "/",
  });
}

export async function getSession() {
  return cookies().get("session")?.value;
}

export async function deleteSession() {
  cookies().delete("session");
}
