import { getBillboard } from "@/api/get-billboard";
import { Heading } from "@/components/heading";
import { Separator } from "@/components/ui/separator";
import { useQuery } from "@tanstack/react-query";
import { Helmet } from "react-helmet-async";
import { useParams } from "react-router-dom";
import { EditBillboardForm } from "./edit-billboard-form";

export function EditBillboard() {
  const { storeId, billboardId } = useParams();

  const { data } = useQuery({
    queryKey: ["billboard", { storeId, billboardId }],
    queryFn: () => {
      if (!storeId || !billboardId) {
        throw new Error("storeId and billboardId must be defined");
      }
      return getBillboard({ storeId, billboardId });
    },
  });

  return (
    <>
      <Helmet title="Editar painel" />
      <div className="flex items-center justify-between">
        <Heading
          title="Editar Painel"
          description="Atualize os detalhes do painel"
        />
      </div>
      <Separator />
      {data && <EditBillboardForm data={data} />}
    </>
  );
}
