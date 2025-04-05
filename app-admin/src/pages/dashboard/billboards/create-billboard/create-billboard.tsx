import { Helmet } from "react-helmet-async";
import { CreateBillboardForm } from "./create-billboard-form";
import { Heading } from "@/components/heading";
import { Separator } from "@/components/ui/separator";

export function CreateBillboard() {
  return (
    <>
      <Helmet title="Criar Painel" />
      <Heading
        title="Criar Painel"
        description="Crie um novo painel para sua loja"
      />
      <Separator />
      <CreateBillboardForm />
      <Separator />
    </>
  );
}
