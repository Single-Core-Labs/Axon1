// Copyright (c) 2026 1Core Labs. MIT License.
"use client"
import { useState } from "react"
import { Card, Button, Input } from "@/components/ui"
import { Search, Filter, ChevronRight, ExternalLink } from "lucide-react"
import { cn } from "@/lib/utils"

const mockTraces = [
  { id: "tr_1", model: "gpt-4o", provider: "openai", latency: "450ms", cost: "$0.024", status: "success", time: "2 mins ago" },
  { id: "tr_2", model: "claude-3-opus", provider: "anthropic", latency: "1.2s", cost: "$0.142", status: "success", time: "5 mins ago" },
  { id: "tr_3", model: "gemini-1.5-pro", provider: "google", latency: "850ms", cost: "$0.008", status: "error", time: "12 mins ago" },
  { id: "tr_4", model: "mistral-large", provider: "mistral", latency: "620ms", cost: "$0.012", status: "success", time: "1 hour ago" },
  { id: "tr_5", model: "gpt-4o-mini", provider: "openai", latency: "210ms", cost: "$0.002", status: "success", time: "2 hours ago" },
]

export default function TracesPage() {
  const [searchTerm, setSearchTerm] = useState("")

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Traces</h1>
          <p className="text-muted-foreground">Deep dive into every request passing through Axon</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline"><Filter className="w-4 h-4 mr-2" /> Filters</Button>
          <Button variant="primary">Export Data</Button>
        </div>
      </div>

      <div className="relative">
        <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-muted-foreground" />
        <Input 
          className="pl-10" 
          placeholder="Search by trace ID, model, or prompt content..." 
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>

      <Card className="overflow-hidden glass">
        <table className="w-full text-sm text-left">
          <thead className="text-xs uppercase bg-muted/50 text-muted-foreground font-semibold">
            <tr>
              <th className="px-6 py-4">Status</th>
              <th className="px-6 py-4">Trace ID</th>
              <th className="px-6 py-4">Model & Provider</th>
              <th className="px-6 py-4">Latency</th>
              <th className="px-6 py-4">Cost</th>
              <th className="px-6 py-4">Time</th>
              <th className="px-6 py-4"></th>
            </tr>
          </thead>
          <tbody className="divide-y divide-border">
            {mockTraces.map((trace) => (
              <tr key={trace.id} className="hover:bg-accent/50 transition-colors cursor-pointer group">
                <td className="px-6 py-4">
                  <span className={cn(
                    "w-2 h-2 rounded-full inline-block mr-2",
                    trace.status === "success" ? "bg-green-500 shadow-[0_0_8px_rgba(34,197,94,0.5)]" : "bg-red-500 shadow-[0_0_8px_rgba(239,68,68,0.5)]"
                  )} />
                </td>
                <td className="px-6 py-4 font-mono text-xs">{trace.id}</td>
                <td className="px-6 py-4">
                  <div className="flex flex-col">
                    <span className="font-semibold">{trace.model}</span>
                    <span className="text-[10px] text-muted-foreground uppercase">{trace.provider}</span>
                  </div>
                </td>
                <td className="px-6 py-4">{trace.latency}</td>
                <td className="px-6 py-4 text-primary font-medium">{trace.cost}</td>
                <td className="px-6 py-4 text-muted-foreground">{trace.time}</td>
                <td className="px-6 py-4 text-right opacity-0 group-hover:opacity-100 transition-opacity">
                  <ChevronRight className="w-4 h-4 text-muted-foreground" />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </Card>
      
      <div className="flex justify-center py-4">
        <Button variant="ghost" className="text-xs">Load more traces</Button>
      </div>
    </div>
  )
}
