import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import type { SaaSSummary } from "@/api/saas"

interface Props {
  summary: SaaSSummary
}

const signalColors: Record<string, string> = {
  invoice:        "bg-blue-100 text-blue-700",
  renewal:        "bg-orange-100 text-orange-700",
  trial_expiring: "bg-red-100 text-red-700",
  welcome:        "bg-green-100 text-green-700",
  usage_report:   "bg-purple-100 text-purple-700",
}

export function SaasSummaryCard({ summary }: Props) {
  const signals = summary.signals ?? []

  return (
    <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm text-muted-foreground">Monthly Spend</CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-3xl font-bold">${summary.total_monthly.toFixed(2)}</p>
          <p className="text-xs text-muted-foreground mt-1">estimated / month</p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm text-muted-foreground">Tools Tracked</CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-3xl font-bold">{summary.tool_count}</p>
          <p className="text-xs text-muted-foreground mt-1">unique products</p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm text-muted-foreground">Signal Types</CardTitle>
        </CardHeader>
        <CardContent className="flex flex-wrap gap-1.5">
          {signals.map((s) => (
            <span
              key={s.signal_type}
              className={`text-xs px-2 py-0.5 rounded-full font-medium ${signalColors[s.signal_type] ?? "bg-muted text-muted-foreground"}`}
            >
              {s.signal_type.replace("_", " ")} · {s.count}
            </span>
          ))}
        </CardContent>
      </Card>
    </div>
  )
}
