import { Helmet } from "react-helmet-async";
import { Link, useParams, useSearchParams } from "react-router-dom";
import { z } from "zod";
import { useQuery } from "@tanstack/react-query";
import { CategoryColumn } from "./categories-table-column";
import { format } from "date-fns";
import { Button } from "@/components/ui/button";
import { Heading } from "@/components/heading";
import { Plus } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { CategoriesTableFilter } from "./categories-table-filter";
import { DataTable } from "@/components/ui/data-table";
import { TablePaginator } from "@/components/table-paginator";
import { ApiList } from "@/components/api-list";
import { getCategoriesPagedList } from "@/api/get-categories-paged-list";
import { getAllBillboards } from "@/api/get-all-billboards";

const CATEGORIES_LIMIT = 10;

export function Categories() {
  const { storeId } = useParams();
  const [searchParams, setSearchParams] = useSearchParams();

  const pageIndex = z.number().parse(Number(searchParams.get("page")) || 1);
  const name = searchParams.get("name");
  const billboardId = searchParams.get("billboardId");

  if (!storeId) {
    throw new Error("O ID da loja é obrigatório.");
  }

  const { data: billboards } = useQuery({
    queryKey: ["billboards", storeId],
    queryFn: () =>
      getAllBillboards({
        storeId: storeId,
      }),
    staleTime: 1000 * 60 * 5,
  });

  const { data: categories } = useQuery({
    queryKey: ["categories", name, billboardId, pageIndex, storeId],
    queryFn: () =>
      getCategoriesPagedList({
        page: pageIndex,
        limit: CATEGORIES_LIMIT,
        name,
        billboardId,
        storeId: storeId,
      }),
  });

  function handlePaginate(newPageIndex: number) {
    setSearchParams((prev) => {
      prev.set("page", newPageIndex.toString());
      return prev;
    });
  }

  const formattedCategories: CategoryColumn[] | undefined =
    categories?.data.map((category) => ({
      id: category.id,
      name: category.name,
      billboardName: category.billboard?.label ?? "Sem painel",
      createdAt: format(new Date(category.createdAt), "dd/MM/yyyy"),
    }));

  return (
    <>
      <Helmet title="Categorias" />
      <div className="flex items-center justify-between">
        <Heading
          title={`Categorias (${categories?.total ?? 0})`}
          description="Gerencie as categorias de sua loja"
        />
        <Button asChild>
          <Link to={`/${storeId}/categories/new`} state={{ billboards }}>
            <Plus className="mr-2 h-4 w-4" />
            Adicionar categoria
          </Link>
        </Button>
      </div>
      <Separator />
      <CategoriesTableFilter billboards={billboards || []} />
      <DataTable columns={CategoryColumn} data={formattedCategories || []} />
      <TablePaginator
        pageIndex={pageIndex}
        totalPages={categories?.totalPages ?? 0}
        onPageChange={handlePaginate}
      />
      <Heading title="API" description="Chamadas de API para as Categorias" />
      <Separator />
      <ApiList entityName="categories" entityIdName="categoryId" />
    </>
  );
}
