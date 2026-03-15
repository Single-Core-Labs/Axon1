// Copyright (c) 2026 1Core Labs. MIT License.
"use client"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { cn } from "@/lib/utils"
import { LayoutDashboard, ListTree, PieChart, LineChart, Settings, Boxes } from "lucide-react"

const menuItems = [
  { name: "Overview", href: "/", icon: LayoutDashboard },
  { name: "Traces", href: "/traces", icon: ListTree },
  { name: "Costs", href: "/costs", icon: PieChart },
  { name: "Latency", href: "/latency", icon: LineChart },
  { name: "Projects", href: "/projects", icon: Boxes },
  { name: "Settings", href: "/settings", icon: Settings },
]

export function Sidebar() {
  const pathname = usePathname()

  return (
    <aside className="w-64 border-r bg-card/50 backdrop-blur-md flex flex-col h-screen sticky top-0">
      <div className="p-6 border-b">
        <h1 className="text-xl font-bold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent italic">
          AXON
        </h1>
      </div>
      <nav className="flex-1 p-4 space-y-1">
        {menuItems.map((item) => {
          const Icon = item.icon
          const active = pathname === item.href || (item.href !== "/" && pathname?.startsWith(item.href))
          return (
            <Link
              key={item.name}
              href={item.href}
              className={cn(
                "flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200",
                active 
                  ? "bg-primary/10 text-primary shadow-sm" 
                  : "text-muted-foreground hover:bg-accent hover:text-accent-foreground"
              )}
            >
              <Icon className="w-4 h-4" />
              {item.name}
            </Link>
          )
        })}
      </nav>
      <div className="p-4 border-t text-xs text-muted-foreground text-center">
        &copy; 2026 1Core Labs
      </div>
    </aside>
  )
}
