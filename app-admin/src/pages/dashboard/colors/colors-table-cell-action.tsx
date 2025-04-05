import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";
import { Copy, Edit, MoreHorizontal, Trash } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { AlertModal } from "@/components/modals/alert-modal";

import { ColorColumn } from "../colors/colors-table-column";
import { deleteColor } from "@/api/delete-color";

export interface ColorsTableCellActionProps {
  data: ColorColumn;
}

export function ColorsTabelCellAction({ data }: ColorsTableCellActionProps) {
  const [isAlertOpen, setIsAlertOpen] = useState(false);
  const { storeId } = useParams();
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const { mutateAsync: deleteColorFn, isPending } = useMutation({
    mutationFn: deleteColor,
  });

  function handleCopyId() {
    navigator.clipboard.writeText(data.id);
    toast.success("ID copiado para a área de transferência");
  }

  function handleEdit() {
    navigate(`/${storeId}/colors/${data.id}`);
  }

  async function handleDeleteColor() {
    if (!storeId) {
      toast.error("Erro: O ID da loja é obrigatório.");
      return;
    }

    try {
      await deleteColorFn({ storeId, colorId: data.id });
      await queryClient.invalidateQueries({
        queryKey: ["colors"],
        exact: false,
      });

      toast.success("Cor removida com sucesso!");
    } catch (error) {
      toast.error(
        "Erro ao remover a cor. Verifique se ela não está sendo usada."
      );
    }
  }

  return (
    <>
      <AlertModal
        isOpen={isAlertOpen}
        onClose={() => setIsAlertOpen(false)}
        onConfirm={handleDeleteColor}
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
          <DropdownMenuItem onClick={() => setIsAlertOpen(true)}>
            <Trash className="w-4 h-4 mr-2" />
            Remover
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
}
