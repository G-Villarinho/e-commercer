import { ColumnDef } from "@tanstack/react-table";
import { SizesTableCellAction } from "./sizes-table-cell-action";

export interface SizeColumn {
  id: string;
  name: string;
  value: string;
  createdAt: string;
}

export const SizeColumn: ColumnDef<SizeColumn>[] = [
  {
    accessorKey: "name",
    header: "Nome",
  },
  {
    accessorKey: "value",
    header: "Valor",
  },
  {
    accessorKey: "createdAt",
    header: "Data de criação",
  },
  {
    id: "actions",
    cell: ({ row }) => <SizesTableCellAction data={row.original} />,
  },
];
