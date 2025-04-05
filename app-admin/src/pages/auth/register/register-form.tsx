import { Link, useNavigate } from "react-router-dom";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useMutation } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { InputError } from "@/components/input-error";
import { registerUser } from "@/api/register-user";
import { useAuth } from "@/hooks/use-auth";
import { LoadingButton } from "@/components/loading-button";

const registerUserSchema = z.object({
  name: z.string().nonempty("O nome é obrigatório."),
  email: z
    .string()
    .nonempty("O e-mail é obrigatório.")
    .email("Por favor, insira um e-mail válido."),
});

type RegisterUserSchema = z.infer<typeof registerUserSchema>;

export function RegisterForm() {
  const { setEmail } = useAuth();
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    setError,
    formState: { isSubmitting, errors },
  } = useForm<RegisterUserSchema>({
    resolver: zodResolver(registerUserSchema),
  });

  const { mutateAsync: registerUserFn } = useMutation({
    mutationFn: registerUser,
  });

  async function handleRegisterUser({ name, email }: RegisterUserSchema) {
    try {
      await registerUserFn({ name, email });
      setEmail(email);
      navigate("/verify-code");
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.status === 409) {
          setError("email", {
            type: "manual",
            message: "Este e-mail já está em uso por outro usuário.",
          });
        }
      }
    }
  }

  return (
    <form onSubmit={handleSubmit(handleRegisterUser)}>
      <div className="grid w-full items-center gap-4">
        <div className="flex flex-col space-y-1.5">
          <Label htmlFor="name">Name</Label>
          <Input
            id="name"
            className="h-12"
            {...register("name")}
            disabled={isSubmitting}
          />
          {errors.name && <InputError error={errors.name.message} />}
        </div>
        <div className="flex flex-col space-y-1.5">
          <Label htmlFor="email">E-mail</Label>
          <Input
            id="email"
            className="h-12"
            {...register("email")}
            disabled={isSubmitting}
          />
          {errors.email && <InputError error={errors.email.message} />}
        </div>
      </div>
      <div className="w-full flex flex-col space-y-4 mt-4">
        <LoadingButton type="submit" isLoading={isSubmitting} size="lg">
          Criar conta
        </LoadingButton>
        <Link
          to="/login"
          className="w-full text-center hover:bg-accent hover:text-accent-foreground rounded-lg py-3 text-sm font-medium transition"
        >
          Entrar agora
        </Link>
      </div>
    </form>
  );
}
