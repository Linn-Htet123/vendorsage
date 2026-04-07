import api from "./client"

export interface Expense {
  id: number
  email_id: number
  merchant_name: string
  amount: number
  currency: string
  category: string
  transaction_date: string
  uncertain: boolean
  source_subject: string
  source_date: string
}

export interface SpendingSummary {
  total: number
  breakdown: { category: string; total: number; count: number }[]
}

export interface SpendingFilter {
  category?: string
  date_start?: string
  date_end?: string
}

export async function fetchExpenses(filter: SpendingFilter = {}): Promise<Expense[]> {
  const params = Object.fromEntries(
    Object.entries(filter).filter(([, v]) => v !== "" && v != null)
  )
  const { data } = await api.get("/api/spending", { params })
  return data.data ?? []
}

export async function fetchSpendingSummary(): Promise<SpendingSummary> {
  const { data } = await api.get("/api/spending/summary")
  return data.data
}
