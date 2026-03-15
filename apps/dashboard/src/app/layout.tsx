// Copyright (c) 2026 1Core Labs. MIT License.
import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { Sidebar } from "@/components/Sidebar";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Axon | AI Gateway & Observability",
  description: "Scale your AI applications with confidence.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased flex`}>
        <Sidebar />
        <main className="flex-1 min-h-screen overflow-y-auto bg-background p-8">
          {children}
        </main>
      </body>
    </html>
  );
}
