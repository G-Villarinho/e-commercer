import "@/index.css";

import { Helmet, HelmetProvider } from "react-helmet-async";
import { RouterProvider } from "react-router-dom";

import { router } from "@/routes";
import { QueryClientProvider } from "@tanstack/react-query";
import { queryClient } from "@/lib/react-query";
import { ThemeProvider } from "@/contexts/theme-context";
import { Toaster } from "react-hot-toast";

export function App() {
  return (
    <HelmetProvider>
      <Helmet titleTemplate="%s | flash.buy" />
      <ThemeProvider defaultTheme="light" storageKey="flash-buy-theme">
        <QueryClientProvider client={queryClient}>
          <RouterProvider router={router} />
          <Toaster />
        </QueryClientProvider>
      </ThemeProvider>
    </HelmetProvider>
  );
}
