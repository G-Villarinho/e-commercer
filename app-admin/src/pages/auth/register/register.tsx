import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Helmet } from "react-helmet-async";
import { RegisterForm } from "./register-form";

export function Register() {
  return (
    <>
      <Helmet title="Registrar" />
      <Card className="w-full max-w-[350px] sm:max-w-[480px]">
        <CardHeader>
          <CardTitle className="text-xl text-center">
            Preencha os dados para criar sua conta
          </CardTitle>
        </CardHeader>
        <CardContent>
          <RegisterForm />
        </CardContent>
      </Card>
    </>
  );
}
