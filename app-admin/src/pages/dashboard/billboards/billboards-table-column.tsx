import { ColumnDef } from "@tanstack/react-table";
import { BillboardsTabelCellAction } from "./billboards-table-cell-action";

export interface BillboardColumn {
  id: string;
  label: string;
  createdAt: string;
}

export const BillboardColumn: ColumnDef<BillboardColumn>[] = [
  {
    accessorKey: "label",
    header: "Rótulo",
  },

  {
    accessorKey: "createdAt",
    header: "Data de criação",
  },
  {
    id: "actions",
    cell: ({ row }) => <BillboardsTabelCellAction data={row.original} />,
  },
];
