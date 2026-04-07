import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import type { SpendingSummary } from "@/api/spending"

interface Props {
  summary: SpendingSummary
}

export function SpendingSummaryCard({ summary }: Props) {
  const breakdown = summary.breakdown ?? []
  const max = breakdown[0]?.total ?? 1

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm text-muted-foreground">Total Spend</CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-3xl font-bold">${summary.total.toFixed(2)}</p>
          <p className="text-xs text-muted-foreground mt-1">
            across {breakdown.length} categor{breakdown.length === 1 ? "y" : "ies"}
          </p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm text-muted-foreground">By Category</CardTitle>
        </CardHeader>
        <CardContent className="space-y-2">
          {breakdown.slice(0, 4).map((b) => (
            <div key={b.category}>
              <div className="flex justify-between text-xs mb-0.5">
                <span className="capitalize text-muted-foreground">{b.category.replace("_", " ")}</span>
                <span className="font-medium">${b.total.toFixed(2)}</span>
              </div>
              <div className="h-1.5 bg-muted rounded-full overflow-hidden">
                <div
                  className="h-full bg-primary rounded-full"
                  style={{ width: `${(b.total / max) * 100}%` }}
                />
              </div>
            </div>
          ))}
        </CardContent>
      </Card>
    </div>
  )
}
