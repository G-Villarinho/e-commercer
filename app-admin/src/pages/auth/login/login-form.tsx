import { Link, useNavigate } from "react-router-dom";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useMutation } from "@tanstack/react-query";
import { login } from "@/api/login";
import { isAxiosError } from "axios";
import { InputError } from "@/components/input-error";
import { useAuth } from "@/hooks/use-auth";
import { LoadingButton } from "@/components/loading-button";

const signInSchema = z.object({
  email: z
    .string()
    .nonempty("O e-mail é obrigatório.")
    .email("Por favor, insira um e-mail válido."),
});

type SignInSchema = z.infer<typeof signInSchema>;

export function LoginForm() {
  const { setEmail } = useAuth();
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { isSubmitting: isLoading, errors },
  } = useForm<SignInSchema>({
    resolver: zodResolver(signInSchema),
  });

  const { mutateAsync: loginFn } = useMutation({
    mutationFn: login,
  });

  async function handleLogin(data: SignInSchema) {
    try {
      await loginFn(data);
      setEmail(data.email);
      navigate("/verify-code");
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.status === 404) {
          setEmail(data.email);
          navigate("/account-not-found");
        }
      }
    }
  }

  return (
    <form onSubmit={handleSubmit(handleLogin)}>
      <div className="grid w-full items-center gap-4">
        <div className="flex flex-col space-y-1.5">
          <Label htmlFor="email">E-mail</Label>
          <Input
            id="email"
            className="h-12"
            {...register("email")}
            disabled={isLoading}
          />
          {errors.email && <InputError error={errors.email.message} />}
        </div>
      </div>
      <div className="w-full flex flex-col space-y-4 mt-4">
        <LoadingButton type="submit" isLoading={isLoading} size="lg">
          Continuar
        </LoadingButton>
        <Link
          to="/register"
          className="w-full text-center hover:bg-accent hover:text-accent-foreground rounded-lg py-3 text-sm font-medium transition"
        >
          Criar conta
        </Link>
      </div>
    </form>
  );
}
