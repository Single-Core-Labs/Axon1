// Copyright (c) 2026 1Core Labs. MIT License.
"use client"
import { Card } from "@/components/ui"
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Legend } from "recharts"

const latencyData = [
  { model: "gpt-4o", p50: 120, p90: 240, p99: 450 },
  { model: "claude-3.5", p50: 210, p90: 420, p99: 800 },
  { model: "gemini-1.5", p50: 180, p90: 380, p99: 650 },
  { model: "mistral-large", p50: 150, p90: 320, p99: 550 },
  { model: "gpt-4o-mini", p50: 80, p90: 150, p99: 280 },
]

export default function LatencyPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Latency</h1>
        <p className="text-muted-foreground">Model response time performance percentiles</p>
      </div>

      <Card className="p-6 h-[500px]">
        <h3 className="font-semibold mb-6 text-center">Latency Performance by Model (ms)</h3>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={latencyData}>
            <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#334155" />
            <XAxis dataKey="model" stroke="#94a3b8" fontSize={12} tickLine={false} axisLine={false} />
            <YAxis stroke="#94a3b8" fontSize={12} tickLine={false} axisLine={false} tickFormatter={(v) => `${v}ms`} />
            <Tooltip 
              contentStyle={{ backgroundColor: "#1e293b", border: "none", borderRadius: "8px", color: "#f8fafc" }} 
            />
            <Legend />
            <Bar dataKey="p50" fill="#3b82f6" radius={[4, 4, 0, 0]} name="p50 (Median)" />
            <Bar dataKey="p90" fill="#6366f1" radius={[4, 4, 0, 0]} name="p90" />
            <Bar dataKey="p99" fill="#10b981" radius={[4, 4, 0, 0]} name="p99" />
          </BarChart>
        </ResponsiveContainer>
      </Card>
    </div>
  )
}
