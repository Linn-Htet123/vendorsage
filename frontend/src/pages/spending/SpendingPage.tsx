import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { fetchExpenses, fetchSpendingSummary, type SpendingFilter } from "@/api/spending"
import { SpendingFilters } from "./SpendingFilters"
import { SpendingTable } from "./SpendingTable"
import { SpendingSummaryCard } from "./SpendingSummaryCard"
import { Card, CardContent } from "@/components/ui/card"

export function SpendingPage() {
  const [filter, setFilter] = useState<SpendingFilter>({})

  const expensesQuery = useQuery({
    queryKey: ["expenses", filter],
    queryFn: () => fetchExpenses(filter),
  })

  const summaryQuery = useQuery({
    queryKey: ["spending-summary"],
    queryFn: fetchSpendingSummary,
  })

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold tracking-tight">Spending</h1>
        <p className="text-sm text-muted-foreground mt-1">Extracted expenses from your emails</p>
      </div>

      {summaryQuery.data && <SpendingSummaryCard summary={summaryQuery.data} />}

      <div className="flex flex-col gap-4">
        <SpendingFilters filter={filter} onChange={setFilter} />

        <Card>
          <CardContent className="p-0">
            {expensesQuery.isLoading && (
              <p className="text-muted-foreground py-10 text-center text-sm">Loading...</p>
            )}
            {expensesQuery.isError && (
              <p className="text-destructive py-10 text-center text-sm">Failed to load expenses.</p>
            )}
            {expensesQuery.data && <SpendingTable expenses={expensesQuery.data} />}
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
