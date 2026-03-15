// Copyright (c) 2026 1Core Labs. MIT License.
"use client"
import { useState } from "react"
import { Button, Card, Input } from "@/components/ui"
import Link from "next/link"

export default function SignupPage() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [orgName, setOrgName] = useState("")

  return (
    <div className="flex items-center justify-center min-h-[80vh]">
      <Card className="w-full max-w-md p-8 space-y-6 shadow-2xl glass">
        <div className="space-y-2 text-center">
          <h1 className="text-3xl font-bold tracking-tight">Create an account</h1>
          <p className="text-muted-foreground">Get started with Axon today</p>
        </div>
        <div className="space-y-4">
          <div className="space-y-2">
            <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">Organization Name</label>
            <Input 
              placeholder="Acme Inc." 
              value={orgName}
              onChange={(e) => setOrgName(e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">Email</label>
            <Input 
              placeholder="name@example.com" 
              type="email" 
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">Password</label>
            <Input 
              type="password" 
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <Button className="w-full h-11" variant="primary">
            Create Account
          </Button>
        </div>
        <div className="text-center text-sm">
          Already have an account?{" "}
          <Link href="/login" className="underline hover:text-primary">
            Sign in
          </Link>
        </div>
      </Card>
    </div>
  )
}
