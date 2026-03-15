// Copyright (c) 2026 1Core Labs. MIT License.
"use client"
import { Card } from "@/components/ui"
import { cn } from "@/lib/utils"
import { Activity, CreditCard, Clock, CheckCircle } from "lucide-react"
import { 
  LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer,
  BarChart, Bar
} from "recharts"

const data = [
  { name: "Mon", cost: 4.2, latency: 450 },
  { name: "Tue", cost: 3.8, latency: 420 },
  { name: "Wed", cost: 5.1, latency: 480 },
  { name: "Thu", cost: 4.5, latency: 440 },
  { name: "Fri", cost: 6.2, latency: 510 },
  { name: "Sat", cost: 2.1, latency: 380 },
  { name: "Sun", cost: 1.8, latency: 360 },
]

export default function OverviewPage() {
  return (
    <div className="space-y-8 animate-in fade-in duration-500">
      <div className="flex items-end justify-between">
        <div>
          <h1 className="text-4xl font-bold tracking-tight">Overview</h1>
          <p className="text-muted-foreground mt-2">Monitor your AI gateway performance and costs</p>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <StatCard title="Total Traces" value="12,482" icon={Activity} increment="+12%" />
        <StatCard title="Total Cost" value="$245.82" icon={CreditCard} increment="+5%" />
        <StatCard title="Avg Latency" value="452ms" icon={Clock} increment="-2%" />
        <StatCard title="Success Rate" value="99.2%" icon={CheckCircle} increment="+0.1%" />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <Card className="p-6 space-y-4">
          <h3 className="font-semibold text-lg">Cost Overview (Last 7 Days)</h3>
          <div className="h-[300px] w-full">
            <ResponsiveContainer width="100%" height="100%">
              <BarChart data={data}>
                <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#334155" />
                <XAxis dataKey="name" stroke="#94a3b8" fontSize={12} tickLine={false} axisLine={false} />
                <YAxis stroke="#94a3b8" fontSize={12} tickLine={false} axisLine={false} tickFormatter={(v) => `$${v}`} />
                <Tooltip 
                  contentStyle={{ backgroundColor: "#1e293b", border: "none", borderRadius: "8px", color: "#f8fafc" }} 
                  itemStyle={{ color: "#3b82f6" }}
                />
                <Bar dataKey="cost" fill="#3b82f6" radius={[4, 4, 0, 0]} />
              </BarChart>
            </ResponsiveContainer>
          </div>
        </Card>

        <Card className="p-6 space-y-4">
          <h3 className="font-semibold text-lg">Latency Trend (p90)</h3>
          <div className="h-[300px] w-full">
            <ResponsiveContainer width="100%" height="100%">
              <LineChart data={data}>
                <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#334155" />
                <XAxis dataKey="name" stroke="#94a3b8" fontSize={12} tickLine={false} axisLine={false} />
                <YAxis stroke="#94a3b8" fontSize={12} tickLine={false} axisLine={false} tickFormatter={(v) => `${v}ms`} />
                <Tooltip 
                  contentStyle={{ backgroundColor: "#1e293b", border: "none", borderRadius: "8px", color: "#f8fafc" }} 
                />
                <Line type="monotone" dataKey="latency" stroke="#3b82f6" strokeWidth={3} dot={{ fill: "#3b82f6" }} />
              </LineChart>
            </ResponsiveContainer>
          </div>
        </Card>
      </div>
    </div>
  )
}

function StatCard({ title, value, icon: Icon, increment }: any) {
  return (
    <Card className="p-6 space-y-2">
      <div className="flex items-center justify-between">
        <span className="text-sm font-medium text-muted-foreground">{title}</span>
        <Icon className="w-4 h-4 text-primary" />
      </div>
      <div className="flex items-end justify-between">
        <span className="text-2xl font-bold">{value}</span>
        <span className={cn(
          "text-xs font-medium px-2 py-1 rounded-full",
          increment.startsWith('+') ? "bg-green-500/10 text-green-500" : "bg-blue-500/10 text-blue-500"
        )}>
          {increment}
        </span>
      </div>
    </Card>
  )
}
