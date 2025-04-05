import { Heading } from "@/components/heading";
import { Helmet } from "react-helmet-async";
import { CreateCategoryForm } from "./create-category-form";
import { useQuery } from "@tanstack/react-query";
import { useLocation, useParams } from "react-router-dom";
import { getAllBillboards } from "@/api/get-all-billboards";
import { Separator } from "@/components/ui/separator";

export function CreateCategory() {
  const { storeId } = useParams();
  const location = useLocation();
  const billboards = location.state?.billboards;

  if (!storeId) {
    throw new Error("storeId is required");
  }

  const { data: fetchedBillboards } = useQuery({
    queryKey: ["billboards", storeId],
    queryFn: () =>
      getAllBillboards({
        storeId: storeId,
      }),
    enabled: !billboards,
  });

  return (
    <>
      <Helmet title="Criar categoria" />
      <Heading
        title="Criar categoria"
        description="Crie uma nova categoria para sua loja"
      />

      <Separator />
      <CreateCategoryForm billboards={billboards || fetchedBillboards || []} />
    </>
  );
}
