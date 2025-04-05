import logo from "@/assets/logo.svg";
import { AuthProvider } from "@/contexts/auth-context";

import { Outlet } from "react-router-dom";

export function AuthLayout() {
  return (
    <AuthProvider>
      <div className="w-full">
        <header className="w-full border-b shadow-sm py-2 px-2 lg:px-80">
          <nav>
            <img src={logo} alt="Logo" />
          </nav>
        </header>

        <main className="w-full p-3 pt-4 sm:pt-10 flex flex-col sm:flex-row justify-center gap-6">
          <Outlet />
        </main>
      </div>
    </AuthProvider>
  );
}
