import { getAllCategories } from "@/api/get-all-categories";
import { getAllColors } from "@/api/get-all-colors";
import { getAllSizes } from "@/api/get-all-sizes";
import { useQueries } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { CreateProductFormSkeleton } from "./create-product-form-skeleton";
import { CreateProductForm } from "./create-product-form";
import { Separator } from "@/components/ui/separator";
import { Heading } from "@/components/heading";
import { Helmet } from "react-helmet-async";

export function CreateProduct() {
  const { storeId } = useParams();

  if (!storeId) {
    throw new Error("O ID da loja é obrigatório.");
  }

  const [categoriesQuery, sizesQuery, colorsQuery] = useQueries({
    queries: [
      {
        queryKey: ["categories", storeId],
        queryFn: () => getAllCategories({ storeId }),
      },
      {
        queryKey: ["sizes", storeId],
        queryFn: () => getAllSizes({ storeId }),
      },
      {
        queryKey: ["colors", storeId],
        queryFn: () => getAllColors({ storeId }),
      },
    ],
  });

  const isLoading =
    categoriesQuery.isLoading || sizesQuery.isLoading || colorsQuery.isLoading;

  return (
    <>
      <Helmet title="Adicionar produto" />
      <Heading
        title="Adicionar produto"
        description="Adicione um novo produto à sua loja."
      />

      <Separator />
      {isLoading ? (
        <CreateProductFormSkeleton />
      ) : (
        <CreateProductForm
          categories={categoriesQuery.data || []}
          sizes={sizesQuery.data || []}
          colors={colorsQuery.data || []}
        />
      )}
    </>
  );
}
