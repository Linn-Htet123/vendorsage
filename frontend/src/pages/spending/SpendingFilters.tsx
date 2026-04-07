import { Input } from "@/components/ui/input"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import type { SpendingFilter } from "@/api/spending"

const CATEGORIES = [
  "food_delivery", "travel", "software",
  "shopping", "utilities", "entertainment", "other",
]

interface Props {
  filter: SpendingFilter
  onChange: (f: SpendingFilter) => void
}

export function SpendingFilters({ filter, onChange }: Props) {
  return (
    <div className="flex flex-wrap items-center gap-3">
      <span className="text-sm font-medium text-muted-foreground">Filter:</span>
      <Select
        value={filter.category ?? ""}
        onValueChange={(v) => onChange({ ...filter, category: !v || v === "all" ? undefined : v })}
      >
        <SelectTrigger className="w-44 h-9">
          <SelectValue placeholder="All categories" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="all">All categories</SelectItem>
          {CATEGORIES.map((c) => (
            <SelectItem key={c} value={c}>{c.replace("_", " ")}</SelectItem>
          ))}
        </SelectContent>
      </Select>

      <div className="flex items-center gap-2">
        <Input
          type="date"
          className="w-40 h-9"
          value={filter.date_start ?? ""}
          onChange={(e) => onChange({ ...filter, date_start: e.target.value || undefined })}
        />
        <span className="text-muted-foreground text-sm">–</span>
        <Input
          type="date"
          className="w-40 h-9"
          value={filter.date_end ?? ""}
          onChange={(e) => onChange({ ...filter, date_end: e.target.value || undefined })}
        />
      </div>
    </div>
  )
}
