import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { CategoryColumn } from "./categories-table-column";
import { Button } from "@/components/ui/button";
import { Copy, Edit, MoreHorizontal, Trash } from "lucide-react";
import toast from "react-hot-toast";
import { useNavigate, useParams } from "react-router-dom";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AlertModal } from "@/components/modals/alert-modal";
import { useState } from "react";
import { deleteCategory } from "@/api/delete-category";

export interface CategoriesTableCellActionProps {
  data: CategoryColumn;
}

export function CategoriesTabelCellAction({
  data,
}: CategoriesTableCellActionProps) {
  const [open, setOpen] = useState(false);
  const { storeId } = useParams();
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  function handleCopyId() {
    navigator.clipboard.writeText(data.id);
    toast.success("ID copiado para a área de transferência");
  }

  function handleEdit() {
    navigate(`/${storeId}/categories/${data.id}`);
  }

  const { mutateAsync: deleteCategoryFn, isPending } = useMutation({
    mutationFn: deleteCategory,
  });

  async function handleDelete() {
    try {
      if (!storeId) {
        throw new Error("O ID da loja é obrigatório.");
      }

      await deleteCategoryFn({ storeId, categoryId: data.id });
      await queryClient.invalidateQueries({
        queryKey: ["categories"],
        exact: false,
      });

      toast.success("Categoria removido com sucesso");
    } catch (error) {
      toast.error(
        "Erro ao remover categoria verifique se a categoria não está sendo usada em algum produto."
      );
    }
  }

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={handleDelete}
        loading={isPending}
      />

      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="outline" size="xxs">
            <span className="sr-only">Abrir menu</span>
            <MoreHorizontal className="w-4 h-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Ações</DropdownMenuLabel>
          <DropdownMenuItem onClick={handleCopyId}>
            <Copy className="w-4 h-4 mr-2" />
            Copiar ID
          </DropdownMenuItem>
          <DropdownMenuItem onClick={handleEdit}>
            <Edit className="w-4 h-4 mr-2" />
            Editar
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setOpen(true)}>
            <Trash className="w-4 h-4 mr-2" />
            Remover
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
}
