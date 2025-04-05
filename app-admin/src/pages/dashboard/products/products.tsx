import { ApiList } from "@/components/api-list";
import { Heading } from "@/components/heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Plus } from "lucide-react";
import { Helmet } from "react-helmet-async";
import { Link, useParams } from "react-router-dom";

export function Products() {
  const { storeId } = useParams();

  if (!storeId) {
    throw new Error("O ID da loja é obrigatório.");
  }

  return (
    <>
      <Helmet title="Produtos" />
      <div className="flex items-center justify-between">
        <Heading
          title={`Produtos (0)`}
          description="Gerencie os produtos da sua loja."
        />
        <Button asChild>
          <Link to={`/${storeId}/products/new`}>
            <Plus className="mr-2 h-4 w-4" />
            Adicionar produto
          </Link>
        </Button>
      </div>
      {/* <Separator />
      <ColorsTableFilter />
      <DataTable columns={ColorColumn} data={formattedColors || []} />
      <TablePaginator
        pageIndex={page}
        totalPages={colors?.totalPages ?? 0}
        onPageChange={handlePaginate}
      /> */}
      <Separator />
      <ApiList entityName="products" entityIdName="productId" />
    </>
  );
}
