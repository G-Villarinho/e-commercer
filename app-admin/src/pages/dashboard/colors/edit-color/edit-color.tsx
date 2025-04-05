import { getColorById } from "@/api/get-color-by-id";
import { Separator } from "@/components/ui/separator";
import { useQuery } from "@tanstack/react-query";
import { Helmet } from "react-helmet-async";
import { useParams } from "react-router-dom";
import { EditColorForm } from "./edit-color-form";
import { EditColorFormSkeleton } from "./edit-color-form-skeleton";
import { Heading } from "@/components/heading";

export function EditColor() {
  const { storeId, colorId } = useParams();

  if (!storeId || !colorId) {
    throw new Error("storeId e colorId são obrigatórios");
  }

  const { data: color, isLoading } = useQuery({
    queryKey: ["color", storeId, colorId],
    queryFn: () => getColorById({ storeId, colorId }),
  });

  return (
    <>
      <Helmet title="Editar cor" />
      <Heading title="Editar cor" description="Altere os dados de uma cor." />
      <Separator />

      {isLoading ? (
        <EditColorFormSkeleton />
      ) : (
        <EditColorForm storeId={storeId} item={color!} />
      )}
    </>
  );
}
