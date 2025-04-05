import { Card, CardContent } from "@/components/ui/card";
import { useAuth } from "@/hooks/use-auth";
import { Helmet } from "react-helmet-async";
import { AccountNotFoundForm } from "./account-not-found-form";

export function AccountNotFound() {
  const { email } = useAuth();

  return (
    <>
      <Helmet title="Conta não encontrada" />
      <h1 className="text-3xl font-semibold max-w-[400px] sm:max-w-[450px] break-words">
        Não encontramos a conta associada ao e-mail: <br /> {email}
      </h1>
      <Card className="w-full max-w-[350px] sm:max-w-[480px]">
        <CardContent>
          <AccountNotFoundForm />
        </CardContent>
      </Card>
    </>
  );
}
