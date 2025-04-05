import { Heading } from "@/components/heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Plus } from "lucide-react";
import { Helmet } from "react-helmet-async";
import { Link, useParams, useSearchParams } from "react-router-dom";
import { z } from "zod";
import { SizesTableFilter } from "./sizes-table-filter";
import { useQuery } from "@tanstack/react-query";
import { getSizesPagedList } from "@/api/get-sizes-paged-list";
import { SizeColumn } from "./sizes-table-column";
import { format } from "date-fns";
import { DataTable } from "@/components/ui/data-table";
import { TablePaginator } from "@/components/table-paginator";
import { ApiList } from "@/components/api-list";

export function Sizes() {
  const { storeId } = useParams();
  const [searchParams, setSearchParams] = useSearchParams();

  const pageIndex = z.number().parse(Number(searchParams.get("page")) || 1);
  const name = searchParams.get("name");

  if (!storeId) {
    throw new Error("O ID da loja é obrigatório.");
  }

  const { data: sizes } = useQuery({
    queryKey: ["sizes", name, pageIndex, storeId],
    queryFn: () => getSizesPagedList({ page: pageIndex, storeId, name }),
  });

  function handlePaginate(newPageIndex: number) {
    setSearchParams((prev) => {
      prev.set("page", newPageIndex.toString());
      return prev;
    });
  }

  const formattedSizes: SizeColumn[] | undefined = sizes?.data.map((size) => ({
    id: size.id,
    name: size.name,
    value: size.value,
    createdAt: format(new Date(size.createdAt), "dd/MM/yyyy"),
  }));
  return (
    <>
      <Helmet title="Tamanhos" />
      <div className="flex items-center justify-between">
        <Heading
          title={`Tamanhos (${sizes?.total ?? 0})`}
          description="Gerencie os tamanhos disponíveis para os produtos da sua loja."
        />
        <Button asChild>
          <Link to={`/${storeId}/sizes/new`}>
            <Plus className="mr-2 h-4 w-4" />
            Adicionar tamanho
          </Link>
        </Button>
      </div>
      <Separator />
      <SizesTableFilter />
      <DataTable columns={SizeColumn} data={formattedSizes || []} />
      <TablePaginator
        pageIndex={pageIndex}
        totalPages={sizes?.totalPages ?? 0}
        onPageChange={handlePaginate}
      />
      <Heading title="API" description="Chamadas de API para as tamanhos" />
      <Separator />
      <ApiList entityName="sizes" entityIdName="sizeId" />
    </>
  );
}
