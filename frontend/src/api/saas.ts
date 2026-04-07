import api from "./client"

export interface SaaSSubscription {
  id: number
  email_id: number
  product_name: string
  signal_type: string
  billing_cycle: string
  estimated_cost: number
  uncertain: boolean
}

export interface SaaSSummary {
  total_monthly: number
  tool_count: number
  signals: { signal_type: string; count: number }[]
}

export async function fetchSaas(): Promise<SaaSSubscription[]> {
  const { data } = await api.get("/api/saas")
  return data.data ?? []
}

export async function fetchSaasSummary(): Promise<SaaSSummary> {
  const { data } = await api.get("/api/saas/summary")
  return data.data
}

export interface UploadResult {
  processed: number
  skipped: number
  failed: number
  total: number
  expense: number
  saas: number
  both: number
  non_financial: number
  uncertain: number
}

export async function uploadEmails(emails: unknown[]): Promise<UploadResult> {
  const { data } = await api.post("/api/emails/upload", { emails })
  return data.data
}
