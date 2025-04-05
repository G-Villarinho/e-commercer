import { useState, useEffect } from "react";
import { useMutation } from "@tanstack/react-query";
import { resendCode } from "@/api/resend-code";
import { isAxiosError } from "axios";
import { useAuth } from "@/hooks/use-auth";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";
import { LoadingButton } from "@/components/loading-button";

interface ResendCodeButtonProps {
  initialCountdown?: number;
}

export function ResendCodeButton({
  initialCountdown = 120,
}: ResendCodeButtonProps) {
  const [countdown, setCountdown] = useState<number>(0);
  const navigate = useNavigate();
  const { setEmail } = useAuth();

  const { mutateAsync: resendCodeFn, isPending } = useMutation({
    mutationFn: resendCode,
  });

  useEffect(() => {
    let timer: NodeJS.Timeout;
    if (countdown > 0) {
      timer = setInterval(() => {
        setCountdown((prev) => prev - 1);
      }, 1000);
    }
    return () => clearInterval(timer);
  }, [countdown]);

  async function handleResendCode() {
    try {
      await resendCodeFn();
      toast.success("Código reenviado com sucesso");
      setCountdown(initialCountdown);
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response?.status === 401 || error.response?.status === 403) {
          setEmail("");
          navigate("/login");
        }
      }
    }
  }

  return (
    <LoadingButton
      type="submit"
      variant="outline"
      className="mt-3 w-full text-black"
      isLoading={isPending}
      onSubmit={handleResendCode}
      size="lg"
    >
      Reenviar código
    </LoadingButton>
  );
}
