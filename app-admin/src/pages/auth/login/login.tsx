import { Card, CardContent } from "@/components/ui/card";
import { Helmet } from "react-helmet-async";
import { LoginForm } from "./login-form";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

export function Login() {
  const navigate = useNavigate();

  useEffect(() => {
    window.history.pushState(null, "", window.location.href);
    const handlePopState = () => {
      navigate(0);
    };

    window.addEventListener("popstate", handlePopState);

    return () => {
      window.removeEventListener("popstate", handlePopState);
    };
  }, [navigate]);

  return (
    <>
      <Helmet title="login" />
      <h1 className="text-3xl font-semibold max-w-[400px] sm:max-w-[450px] break-words">
        Digite seu e-mail para iniciar sessão
      </h1>
      <Card className="w-full max-w-[350px] sm:max-w-[480px]">
        <CardContent>
          <LoginForm />
        </CardContent>
      </Card>
    </>
  );
}
