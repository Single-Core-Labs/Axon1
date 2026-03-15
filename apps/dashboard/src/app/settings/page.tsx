// Copyright (c) 2026 1Core Labs. MIT License.
"use client"
import { Card, Button, Input } from "@/components/ui"
import { Key, Shield, Bell, User } from "lucide-react"

export default function SettingsPage() {
  return (
    <div className="max-w-4xl mx-auto space-y-8 pb-12">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Settings</h1>
        <p className="text-muted-foreground">Manage your organization and project settings</p>
      </div>

      <Section icon={User} title="Profile" description="Update your personal information">
        <div className="grid grid-cols-2 gap-4">
          <Input placeholder="Full Name" defaultValue="Lucas" />
          <Input placeholder="Email Address" defaultValue="lucas@example.com" disabled />
        </div>
        <Button className="mt-4">Save Changes</Button>
      </Section>

      <Section icon={Key} title="API Keys" description="Keys used for authentication in the Axon SDK">
        <div className="space-y-4">
          <Card className="p-4 flex items-center justify-between border-dashed">
            <div>
              <p className="text-sm font-bold">ax_78d...9f32</p>
              <p className="text-xs text-muted-foreground">Created 2 days ago • Active</p>
            </div>
            <Button variant="ghost" className="text-red-500">Revoke</Button>
          </Card>
          <Button variant="outline" className="w-full">Generate New Key</Button>
        </div>
      </Section>

      <Section icon={Shield} title="Security" description="Password and authentication settings">
        <Button variant="outline">Change Password</Button>
        <div className="mt-4 flex items-center justify-between">
          <p className="text-sm">Two-Factor Authentication</p>
          <Button variant="ghost" className="text-primary h-auto p-0">Enable</Button>
        </div>
      </Section>
    </div>
  )
}

function Section({ icon: Icon, title, description, children }: any) {
  return (
    <Card className="p-6">
      <div className="flex items-start gap-4 mb-6">
        <div className="p-2 rounded-lg bg-primary/10 text-primary">
          <Icon className="w-5 h-5" />
        </div>
        <div>
          <h3 className="font-semibold">{title}</h3>
          <p className="text-sm text-muted-foreground">{description}</p>
        </div>
      </div>
      <div className="ml-11">{children}</div>
    </Card>
  )
}
