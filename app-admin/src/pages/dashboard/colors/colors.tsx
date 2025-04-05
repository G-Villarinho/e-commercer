import { getColorsPagedList } from "@/api/get-colors-paged-list";
import { Heading } from "@/components/heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { useQuery } from "@tanstack/react-query";
import { Plus } from "lucide-react";
import { Helmet } from "react-helmet-async";
import { Link, useParams, useSearchParams } from "react-router-dom";
import { z } from "zod";
import { ColorsTableFilter } from "./colors-table-filter";
import { ColorColumn } from "./colors-table-column";
import { TablePaginator } from "@/components/table-paginator";
import { DataTable } from "@/components/ui/data-table";
import { format } from "date-fns/format";
import { ApiList } from "@/components/api-list";

export function Colors() {
  const { storeId } = useParams();
  const [searchParams, setSearchParams] = useSearchParams();

  if (!storeId) {
    throw new Error("O ID da loja é obrigatório.");
  }

  const page = z.number().parse(Number(searchParams.get("page")) || 1);
  const name = searchParams.get("name");

  const { data: colors } = useQuery({
    queryKey: ["colors", name, page, storeId],
    queryFn: () =>
      getColorsPagedList({
        page,
        name,
        storeId,
      }),
  });

  function handlePaginate(newPageIndex: number) {
    setSearchParams((prev) => {
      prev.set("page", newPageIndex.toString());
      return prev;
    });
  }

  const formattedColors: ColorColumn[] | undefined = colors?.data.map(
    (category) => ({
      id: category.id,
      name: category.name,
      hex: category.hex,
      createdAt: format(new Date(category.createdAt), "dd/MM/yyyy"),
    })
  );

  return (
    <>
      <Helmet title="Cores" />
      <div className="flex items-center justify-between">
        <Heading
          title={`Cores (${colors?.total ?? 0})`}
          description="Gerencie as cores disponíveis para os produtos da sua loja."
        />
        <Button asChild>
          <Link to={`/${storeId}/colors/new`}>
            <Plus className="mr-2 h-4 w-4" />
            Adicionar cor
          </Link>
        </Button>
      </div>
      <Separator />
      <ColorsTableFilter />
      <DataTable columns={ColorColumn} data={formattedColors || []} />
      <TablePaginator
        pageIndex={page}
        totalPages={colors?.totalPages ?? 0}
        onPageChange={handlePaginate}
      />
      <Separator />
      <ApiList entityIdName="colorId" entityName="colors" />
    </>
  );
}
