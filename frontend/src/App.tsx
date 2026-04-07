import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Toaster } from "@/components/ui/sonner";
import { Navbar } from "@/components/layout/Navbar";
import { TotalsBar } from "@/components/layout/TotalsBar";
import { SpendingPage } from "@/pages/spending/SpendingPage";
import { SaasPage } from "@/pages/saas/SaasPage";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 30_000,
      retry: 1,
    },
  },
});

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <div className="min-h-screen bg-background">
          <Navbar />
          <TotalsBar />
          <main className="mx-auto px-6 py-8">
            <Routes>
              <Route path="/" element={<Navigate to="/spending" replace />} />
              <Route path="/spending" element={<SpendingPage />} />
              <Route path="/saas" element={<SaasPage />} />
            </Routes>
          </main>
        </div>
        <Toaster richColors position="bottom-right" />
      </BrowserRouter>
    </QueryClientProvider>
  );
}
