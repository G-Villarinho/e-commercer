import { Heading } from "@/components/heading";
import { Helmet } from "react-helmet-async";
import { EditSizeForm } from "./edit-size-form";
import { Navigate, useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { getSizeById } from "@/api/get-size-by-id";
import { EditSizeFormSkeleton } from "./edit-form-skeleton";
import { Separator } from "@/components/ui/separator";

export function EditSize() {
  const { storeId, sizeId } = useParams();

  if (!storeId || !sizeId) {
    throw new Error("storeId e sizeId são obrigatórios");
  }

  const { data: size, isPending } = useQuery({
    queryKey: ["size", storeId, sizeId],
    queryFn: () => getSizeById({ storeId, sizeId }),
  });

  if (!size) {
    return <Navigate to={`/${storeId}/sizes`} />;
  }

  return (
    <>
      <Helmet title="Editar tamanho" />
      <Heading
        title="Editar tamanho"
        description="Altere os dados de um tamanho."
      />

      <Separator />

      {isPending ? <EditSizeFormSkeleton /> : <EditSizeForm size={size} />}
    </>
  );
}
