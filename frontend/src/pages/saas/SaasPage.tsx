import { useQuery } from "@tanstack/react-query"
import { fetchSaas, fetchSaasSummary } from "@/api/saas"
import { SaasTable } from "./SaasTable"
import { SaasSummaryCard } from "./SaasSummaryCard"
import { Card, CardContent } from "@/components/ui/card"

export function SaasPage() {
  const saasQuery = useQuery({
    queryKey: ["saas"],
    queryFn: fetchSaas,
  })

  const summaryQuery = useQuery({
    queryKey: ["saas-summary"],
    queryFn: fetchSaasSummary,
  })

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold tracking-tight">SaaS Subscriptions</h1>
        <p className="text-sm text-muted-foreground mt-1">Tools and services detected from your emails</p>
      </div>

      {summaryQuery.data && <SaasSummaryCard summary={summaryQuery.data} />}

      <Card>
        <CardContent className="p-0">
          {saasQuery.isLoading && (
            <p className="text-muted-foreground py-10 text-center text-sm">Loading...</p>
          )}
          {saasQuery.isError && (
            <p className="text-destructive py-10 text-center text-sm">Failed to load subscriptions.</p>
          )}
          {saasQuery.data && <SaasTable subs={saasQuery.data} />}
        </CardContent>
      </Card>
    </div>
  )
}
