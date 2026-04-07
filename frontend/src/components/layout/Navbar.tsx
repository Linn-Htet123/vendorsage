import { NavLink } from "react-router-dom"
import { useQuery } from "@tanstack/react-query"
import { UploadEmails } from "./UploadEmails"
import { fetchSpendingSummary } from "@/api/spending"
import { fetchSaasSummary } from "@/api/saas"

export function Navbar() {
  const spending = useQuery({ queryKey: ["spending-summary"], queryFn: fetchSpendingSummary })
  const saas = useQuery({ queryKey: ["saas-summary"], queryFn: fetchSaasSummary })

  return (
    <nav className="border-b bg-white px-6 py-3 flex items-center gap-6">
      <span className="font-semibold text-lg tracking-tight">VendorSage</span>

      <NavLink
        to="/spending"
        className={({ isActive }) =>
          "flex items-center gap-1.5 text-sm " +
          (isActive ? "text-primary font-medium" : "text-muted-foreground hover:text-primary")
        }
      >
        Spending
        {spending.data && (
          <span className="text-xs bg-muted px-1.5 py-0.5 rounded-full">
            ${spending.data.total.toFixed(0)}
          </span>
        )}
      </NavLink>

      <NavLink
        to="/saas"
        className={({ isActive }) =>
          "flex items-center gap-1.5 text-sm " +
          (isActive ? "text-primary font-medium" : "text-muted-foreground hover:text-primary")
        }
      >
        SaaS
        {saas.data && (
          <span className="text-xs bg-muted px-1.5 py-0.5 rounded-full">
            {saas.data.tool_count} tools
          </span>
        )}
      </NavLink>

      <div className="ml-auto">
        <UploadEmails />
      </div>
    </nav>
  )
}
