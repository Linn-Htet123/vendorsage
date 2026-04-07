import {
  Table, TableBody, TableCell,
  TableHead, TableHeader, TableRow,
} from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import type { SaaSSubscription } from "@/api/saas"

interface Props {
  subs: SaaSSubscription[]
}

const signalVariant: Record<string, "default" | "secondary" | "destructive" | "outline"> = {
  invoice: "default",
  renewal: "destructive",
  trial_expiring: "destructive",
  welcome: "secondary",
  usage_report: "outline",
}

export function SaasTable({ subs }: Props) {
  if (subs.length === 0) {
    return (
      <div className="py-12 text-center">
        <p className="text-muted-foreground text-sm">No SaaS subscriptions found.</p>
        <p className="text-muted-foreground text-xs mt-1">Upload emails to detect tools and subscriptions.</p>
      </div>
    )
  }

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Product</TableHead>
          <TableHead>Signal</TableHead>
          <TableHead>Billing</TableHead>
          <TableHead>Cost</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {subs.map((s) => (
          <TableRow key={s.id}>
            <TableCell className="font-medium">
              {s.product_name}
              {s.uncertain && (
                <Badge variant="outline" className="ml-2 text-xs">uncertain</Badge>
              )}
            </TableCell>
            <TableCell>
              <Badge variant={signalVariant[s.signal_type] ?? "secondary"}>
                {s.signal_type.replace("_", " ")}
              </Badge>
            </TableCell>
            <TableCell className="text-muted-foreground capitalize">
              {s.billing_cycle}
            </TableCell>
            <TableCell>
              {s.estimated_cost > 0 ? `$${s.estimated_cost.toFixed(2)}` : "—"}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}
