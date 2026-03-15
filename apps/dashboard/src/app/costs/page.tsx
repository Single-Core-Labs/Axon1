// Copyright (c) 2026 1Core Labs. MIT License.
"use client"
import { Card } from "@/components/ui"
import { PieChart, Pie, Cell, ResponsiveContainer, Tooltip, Legend } from "recharts"

const modelData = [
  { name: "GPT-4o", value: 125.40, color: "#3b82f6" },
  { name: "Claude 3.5 Sonnet", value: 84.20, color: "#6366f1" },
  { name: "Gemini 1.5 Pro", value: 15.60, color: "#10b981" },
  { name: "Mistral Large", value: 20.62, color: "#f59e0b" },
]

export default function CostsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Costs</h1>
        <p className="text-muted-foreground">Detailed breakdown of your AI spending</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <Card className="lg:col-span-2 p-6 flex flex-col justify-center items-center h-[450px]">
          <h3 className="font-semibold mb-6 flex-none">Spending by Model</h3>
          <div className="w-full flex-1">
            <ResponsiveContainer width="100%" height="100%">
              <PieChart>
                <Pie
                  data={modelData}
                  cx="50%"
                  cy="50%"
                  innerRadius={80}
                  outerRadius={120}
                  paddingAngle={5}
                  dataKey="value"
                >
                  {modelData.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.color} />
                  ))}
                </Pie>
                <Tooltip 
                  contentStyle={{ backgroundColor: "#1e293b", border: "none", borderRadius: "8px", color: "#f8fafc" }} 
                  formatter={(value) => [`$${value}`, "Cost"]}
                />
                <Legend verticalAlign="bottom" height={36}/>
              </PieChart>
            </ResponsiveContainer>
          </div>
        </Card>

        <Card className="p-6 space-y-6">
          <h3 className="font-semibold">Top Expense Projects</h3>
          <div className="space-y-4">
            <ProjectCostItem name="Production Gateway" cost="$142.40" percentage={60} color="bg-blue-500" />
            <ProjectCostItem name="Development Sandbox" cost="$45.20" percentage={20} color="bg-indigo-500" />
            <ProjectCostItem name="Marketing Experiments" cost="$32.90" percentage={15} color="bg-green-500" />
            <ProjectCostItem name="Internal Tools" cost="$25.32" percentage={5} color="bg-amber-500" />
          </div>
        </Card>
      </div>
    </div>
  )
}

function ProjectCostItem({ name, cost, percentage, color }: any) {
  return (
    <div className="space-y-2">
      <div className="flex justify-between text-sm">
        <span className="font-medium">{name}</span>
        <span className="text-primary font-bold">{cost}</span>
      </div>
      <div className="w-full bg-muted rounded-full h-2">
        <div className={`${color} h-2 rounded-full`} style={{ width: `${percentage}%` }}></div>
      </div>
    </div>
  )
}
