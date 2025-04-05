import { Card, CardContent } from "@/components/ui/card";
import { useAuth } from "@/hooks/use-auth";
import { Helmet } from "react-helmet-async";
import { VerifyCodeForm } from "./verify-code-form";
import { useNavigate } from "react-router-dom";
import { isAxiosError } from "axios";
import { useEffect, useRef } from "react";
import { useMutation } from "@tanstack/react-query";
import { checkCode } from "@/api/check-code";
import { ResendCodeButton } from "./resend-code-button";

export function VerifyCode() {
  const alreadyCalled = useRef(false);
  const navigate = useNavigate();
  const { email, setEmail } = useAuth();

  const { mutateAsync: checkCodeFn } = useMutation({
    mutationFn: checkCode,
  });

  useEffect(() => {
    async function handleCheckCode() {
      if (!email) {
        navigate("/login");
        return;
      }

      if (alreadyCalled.current) return;
      alreadyCalled.current = true;

      try {
        await checkCodeFn();
      } catch (error) {
        if (isAxiosError(error)) {
          if (
            error.response?.status === 401 ||
            error.response?.status === 403
          ) {
            setEmail("");
            navigate("/login");
          }
        }
      }
    }

    handleCheckCode();
  }, [navigate, setEmail, checkCodeFn, email]);

  return (
    <>
      <Helmet title="verificar código" />
      <h1 className="text-2xl font-semibold max-w-[400px] sm:max-w-[450px] break-words">
        Digite o código enviado para {email} para verificar
      </h1>
      <Card className="w-full max-w-[350px] sm:max-w-[480px]">
        <CardContent>
          <VerifyCodeForm />
          <ResendCodeButton />
        </CardContent>
      </Card>
    </>
  );
}
