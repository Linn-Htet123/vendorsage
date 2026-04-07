import { useQuery } from "@tanstack/react-query"
import { fetchSpendingSummary } from "@/api/spending"
import { fetchSaasSummary } from "@/api/saas"

function Stat({ label, value, highlight }: { label: string; value: string; highlight?: boolean }) {
  return (
    <div className={`flex flex-col items-center px-8 py-4 ${highlight ? "bg-primary text-primary-foreground rounded-xl" : ""}`}>
      <span className={`text-[11px] font-semibold uppercase tracking-widest ${highlight ? "opacity-70" : "text-muted-foreground"}`}>
        {label}
      </span>
      <span className={`text-2xl font-bold mt-1 ${highlight ? "" : "text-foreground"}`}>
        {value}
      </span>
    </div>
  )
}

export function TotalsBar() {
  const spending = useQuery({ queryKey: ["spending-summary"], queryFn: fetchSpendingSummary })
  const saas = useQuery({ queryKey: ["saas-summary"], queryFn: fetchSaasSummary })

  const totalSpend = spending.data?.total ?? 0
  const totalSaaS = saas.data?.total_monthly ?? 0
  const combined = totalSpend + totalSaaS

  return (
    <div className="border-b bg-muted/20 py-2 flex items-center justify-center gap-1">
      <Stat label="Total Spending" value={`$${totalSpend.toFixed(2)}`} />
      <span className="text-muted-foreground/50 text-xl font-light px-1">+</span>
      <Stat label="Monthly SaaS" value={`$${totalSaaS.toFixed(2)}`} />
      <span className="text-muted-foreground/50 text-xl font-light px-1">=</span>
      <Stat label="Combined Total" value={`$${combined.toFixed(2)}`} highlight />
    </div>
  )
}
