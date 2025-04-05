import { createContext, ReactNode, useState } from "react";

interface AuthContextType {
  email: string | null;
  setEmail: (email: string | null) => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined
);

interface AuthProviderProps {
  children: ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [email, setEmailState] = useState<string | null>(
    sessionStorage.getItem("userEmail") || null
  );

  function setEmail(email: string | null) {
    if (email) {
      sessionStorage.setItem("userEmail", email);
    } else {
      sessionStorage.removeItem("userEmail");
    }
    setEmailState(email);
  }

  return (
    <AuthContext.Provider value={{ email, setEmail }}>
      {children}
    </AuthContext.Provider>
  );
}
