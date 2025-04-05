import { getAllBillboards } from "@/api/get-all-billboards";
import { getCategoryById } from "@/api/get-category-by-id";
import { Separator } from "@/components/ui/separator";
import { useQuery } from "@tanstack/react-query";
import { Helmet } from "react-helmet-async";
import { useParams, useLocation } from "react-router-dom";
import { EditCategoryForm } from "./edit-category-form";
import { EditCategoryFormSkeleton } from "./edit-category-form-skeleton";

export function EditCategory() {
  const { storeId, categoryId } = useParams();
  const location = useLocation();
  const billboards = location.state?.billboards;

  if (!storeId || !categoryId) {
    throw new Error("storeId e categoryId são obrigatórios");
  }

  const { data: fetchedBillboards } = useQuery({
    queryKey: ["billboards", storeId],
    queryFn: () => getAllBillboards({ storeId }),
    enabled: !billboards,
  });

  const { data: category, isLoading } = useQuery({
    queryKey: ["category", storeId, categoryId],
    queryFn: () => getCategoryById({ storeId, categoryId }),
  });

  return (
    <>
      <Helmet title="Editar categoria" />
      <Separator />
      {isLoading ? (
        <EditCategoryFormSkeleton />
      ) : (
        <EditCategoryForm
          data={category!}
          billboards={billboards || fetchedBillboards || []}
        />
      )}
    </>
  );
}
