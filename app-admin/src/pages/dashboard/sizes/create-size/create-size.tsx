import { Heading } from "@/components/heading";
import { Helmet } from "react-helmet-async";
import { CreateSizeForm } from "./create-size-form";
import { Separator } from "@/components/ui/separator";

export function CreateSize() {
  return (
    <>
      <Helmet title="Criar tamanho" />
      <Heading
        title="Criar tamanho"
        description="Adicione um novo tamanho para os produtos da sua loja."
      />
      <Separator />
      <CreateSizeForm />
    </>
  );
}
