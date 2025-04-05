import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from "@/components/ui/input-otp";
import { zodResolver } from "@hookform/resolvers/zod";
import { REGEXP_ONLY_DIGITS_AND_CHARS } from "input-otp";
import { useForm } from "react-hook-form";
import { Fragment } from "react";
import { z } from "zod";
import { InputError } from "@/components/input-error";
import { VerifyOtp } from "@/api/verify-otp";
import { useMutation } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/use-auth";
import toast from "react-hot-toast";
import { LoadingButton } from "@/components/loading-button";

const OTP_SIZE = 6;

const verifyCodeSchema = z.object({
  code: z
    .string()
    .nonempty("Por favor, insira um código de 6 caracteres")
    .length(OTP_SIZE, {
      message: "O código deve ter exatamente 6 caracteres.",
    }),
});

type VerifyCodeData = z.infer<typeof verifyCodeSchema>;

export function VerifyCodeForm() {
  const navigate = useNavigate();
  const { setEmail } = useAuth();

  const {
    setValue,
    handleSubmit,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<VerifyCodeData>({
    resolver: zodResolver(verifyCodeSchema),
    defaultValues: {
      code: "",
    },
  });

  const { mutateAsync: verifyCodeFn } = useMutation({
    mutationFn: VerifyOtp,
  });

  async function handleVerifyCode(data: VerifyCodeData) {
    try {
      await verifyCodeFn(data);
      setEmail(null);
      navigate("/");
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response?.status === 401 || error.response?.status === 403) {
          setEmail("");
          navigate("/login");
          toast.error("Sessão expirada. Por favor, faça login novamente.");
        }

        if (error.response?.status === 404) {
          setError("code", {
            type: "manual",
            message: "Código inválido. Por favor, tente novamente.",
          });
        }
      }
    }
  }

  return (
    <form onSubmit={handleSubmit(handleVerifyCode)}>
      <div className="flex flex-col items-center gap-4">
        <div className="flex justify-center gap-2">
          <InputOTP
            maxLength={OTP_SIZE}
            pattern={REGEXP_ONLY_DIGITS_AND_CHARS}
            className="flex justify-center gap-2"
            onChange={(value) =>
              setValue("code", value, { shouldValidate: true })
            }
          >
            <InputOTPGroup className="flex justify-center gap-2">
              {[...Array(OTP_SIZE)].map((_, index) => (
                <Fragment key={index}>
                  <InputOTPSlot
                    index={index}
                    className="sm:h-12 sm:w-12 uppercase text-lg text-center border-2 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
                  />
                </Fragment>
              ))}
            </InputOTPGroup>
          </InputOTP>
        </div>

        {errors.code && <InputError error={errors.code.message} />}
      </div>
      <LoadingButton
        type="submit"
        className="mt-8 w-full"
        isLoading={isSubmitting}
        size="lg"
      >
        Verificar código
      </LoadingButton>
    </form>
  );
}
