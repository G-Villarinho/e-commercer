import { Helmet } from "react-helmet-async";
import { Link, useParams, useSearchParams } from "react-router-dom";
import { z } from "zod";
import { useQuery } from "@tanstack/react-query";
import { getBillboardsPagedList } from "@/api/get-billboards-paged-list";
import { BillboardColumn } from "./billboards-table-column";
import { format } from "date-fns";
import { Button } from "@/components/ui/button";
import { Heading } from "@/components/heading";
import { Plus } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { BillboardsTableFilter } from "./billboards-table-filter";
import { DataTable } from "@/components/ui/data-table";
import { TablePaginator } from "@/components/table-paginator";
import { ApiList } from "@/components/api-list";

const BILLBOARD_LIMIT = 10;

export function Billboards() {
  const { storeId } = useParams() as { storeId: string };
  const [searchParams, setSearchParams] = useSearchParams();

  const pageIndex = z.number().parse(Number(searchParams.get("page")) || 1);
  const label = searchParams.get("label");

  const { data: billboards } = useQuery({
    queryKey: ["billboards", label, pageIndex],
    queryFn: () =>
      getBillboardsPagedList({
        page: pageIndex,
        limit: BILLBOARD_LIMIT,
        label,
        storeId: storeId || "",
      }),
  });

  function handlePaginate(newPageIndex: number) {
    setSearchParams((prev) => {
      prev.set("page", newPageIndex.toString());
      return prev;
    });
  }

  const formattedBillboards: BillboardColumn[] | undefined =
    billboards?.data.map((billboard) => ({
      id: billboard.id,
      label: billboard.label,
      createdAt: format(new Date(billboard.createdAt), "dd/MM/yyyy"),
    }));

  return (
    <>
      <Helmet title="Paines" />
      <div className="flex items-center justify-between">
        <Heading
          title={`Painéis (${billboards?.total ?? 0})`}
          description="Gerencie os painés de sua loja"
        />
        <Button asChild>
          <Link to={`/${storeId}/billboards/new`}>
            <Plus className="mr-2 h-4 w-4" />
            Adicionar painel
          </Link>
        </Button>
      </div>
      <Separator />
      <BillboardsTableFilter />
      <DataTable columns={BillboardColumn} data={formattedBillboards || []} />
      <TablePaginator
        pageIndex={pageIndex}
        totalPages={billboards?.totalPages ?? 0}
        onPageChange={handlePaginate}
      />
      <Heading title="API" description="Chamadas de API para os Painés" />
      <Separator />
      <ApiList entityName="billboards" entityIdName="billboardId" />
    </>
  );
}
