import { useRef, useState } from "react"
import { toast } from "sonner"
import { Button } from "@/components/ui/button"
import { uploadEmails } from "@/api/saas"
import { useQueryClient } from "@tanstack/react-query"

function parseCSV(text: string): object[] {
  const lines = text.trim().split("\n")
  const headers = lines[0].split(",").map((h) => h.trim().replace(/^"|"$/g, ""))
  return lines.slice(1).map((line) => {
    const values = line.match(/(".*?"|[^,]+)/g) ?? []
    return Object.fromEntries(
      headers.map((h, i) => [h, (values[i] ?? "").replace(/^"|"$/g, "").trim()])
    )
  })
}

export function UploadEmails() {
  const inputRef = useRef<HTMLInputElement>(null)
  const [loading, setLoading] = useState(false)
  const queryClient = useQueryClient()

  async function handleFile(e: React.ChangeEvent<HTMLInputElement>) {
    const file = e.target.files?.[0]
    if (!file) return

    setLoading(true)
    const toastId = toast.loading(`Processing ${file.name}…`)
    try {
      const text = await file.text()
      const isCSV = file.name.endsWith(".csv")
      const parsed = isCSV ? parseCSV(text) : JSON.parse(text)
      const list = Array.isArray(parsed) ? parsed : parsed.emails
      const result = await uploadEmails(list)
      await queryClient.invalidateQueries()

      const breakdown = [
        result.expense > 0 && `${result.expense} expense`,
        result.saas > 0 && `${result.saas} saas`,
        result.both > 0 && `${result.both} both`,
        result.non_financial > 0 && `${result.non_financial} non-financial`,
        result.uncertain > 0 && `${result.uncertain} uncertain`,
      ].filter(Boolean).join(" · ")

      const description = [
        breakdown && `Classified: ${breakdown}`,
        result.skipped > 0 && `${result.skipped} duplicates skipped`,
        result.failed > 0 && `${result.failed} failed`,
      ].filter(Boolean).join(" · ")

      toast.success(`${result.processed} emails processed`, {
        id: toastId,
        description,
      })
    } catch {
      toast.error("Upload failed", { id: toastId, description: "Check file format (JSON or CSV)." })
    }
    setLoading(false)
    e.target.value = ""
  }

  return (
    <>
      <Button
        variant="outline"
        size="sm"
        disabled={loading}
        onClick={() => inputRef.current?.click()}
      >
        {loading ? "Processing…" : "Import JSON / CSV"}
      </Button>
      <input ref={inputRef} type="file" accept=".json,.csv" className="hidden" onChange={handleFile} />
    </>
  )
}
