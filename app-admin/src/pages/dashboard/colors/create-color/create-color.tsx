import { Heading } from "@/components/heading";
import { Separator } from "@/components/ui/separator";
import { Helmet } from "react-helmet-async";
import { useParams } from "react-router-dom";
import { CreateColorForm } from "./create-color-form";

export function CreateColor() {
  const { storeId } = useParams();

  if (!storeId) {
    throw new Error("O ID da loja é obrigatório.");
  }

  return (
    <>
      <Helmet title="Criar cor" />
      <Heading
        title="Criar cor"
        description="Adicione uma nova cor para os produtos da sua loja."
      />
      <Separator />
      <CreateColorForm storeId={storeId} />
    </>
  );
}
