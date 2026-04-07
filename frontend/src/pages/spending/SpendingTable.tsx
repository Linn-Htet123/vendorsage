import {
  Table, TableBody, TableCell,
  TableHead, TableHeader, TableRow,
} from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import type { Expense } from "@/api/spending"

interface Props {
  expenses: Expense[]
}

export function SpendingTable({ expenses }: Props) {
  if (expenses.length === 0) {
    return (
      <div className="py-12 text-center">
        <p className="text-muted-foreground text-sm">No expenses found.</p>
        <p className="text-muted-foreground text-xs mt-1">Try uploading emails or clearing your filters.</p>
      </div>
    )
  }

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Merchant</TableHead>
          <TableHead>Amount</TableHead>
          <TableHead>Category</TableHead>
          <TableHead>Date</TableHead>
          <TableHead>Source Email</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {expenses.map((e) => {
          const isRefund = e.amount < 0
          return (
            <TableRow key={e.id}>
              <TableCell className="font-medium">
                {e.merchant_name}
                {isRefund && (
                  <Badge variant="outline" className="ml-2 text-xs text-green-600 border-green-400">
                    Refund
                  </Badge>
                )}
                {e.uncertain && (
                  <Badge variant="outline" className="ml-2 text-xs">uncertain</Badge>
                )}
              </TableCell>
              <TableCell className={isRefund ? "text-green-600 font-medium" : ""}>
                {e.currency} {e.amount.toFixed(2)}
              </TableCell>
              <TableCell>
                <Badge variant="secondary">{e.category.replace("_", " ")}</Badge>
              </TableCell>
              <TableCell className="text-muted-foreground">
                {e.transaction_date.slice(0, 10)}
              </TableCell>
              <TableCell className="text-muted-foreground max-w-[220px]">
                <p className="text-xs truncate" title={e.source_subject}>{e.source_subject}</p>
                <p className="text-xs opacity-60">{e.source_date}</p>
              </TableCell>
            </TableRow>
          )
        })}
      </TableBody>
    </Table>
  )
}
