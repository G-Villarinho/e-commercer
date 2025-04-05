import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { BillboardColumn } from "./billboards-table-column";
import { Button } from "@/components/ui/button";
import { Copy, Edit, MoreHorizontal, Trash } from "lucide-react";
import toast from "react-hot-toast";
import { useNavigate, useParams } from "react-router-dom";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { deleteBillboard } from "@/api/delete-billboard";
import { useState } from "react";
import { AlertModal } from "@/components/modals/alert-modal";

export interface BillboardTableCellActionProps {
  data: BillboardColumn;
}

export function BillboardsTabelCellAction({
  data,
}: BillboardTableCellActionProps) {
  const { storeId } = useParams();
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const [open, setOpen] = useState(false);

  function handleCopyId() {
    navigator.clipboard.writeText(data.id);
    toast.success("ID do painel copiado para a área de transferência");
  }

  function handleEdit() {
    navigate(`/${storeId}/billboards/${data.id}`);
  }

  const { mutateAsync: deleteBillboardFn, isPending } = useMutation({
    mutationFn: deleteBillboard,
  });

  async function handleDelete() {
    try {
      await deleteBillboardFn({ storeId: storeId || "", billboardId: data.id });
      queryClient.invalidateQueries({ queryKey: ["billboards"] });
      toast.success("Painel removido com sucesso");
    } catch (error) {
      toast.error("Erro ao remover painel");
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
