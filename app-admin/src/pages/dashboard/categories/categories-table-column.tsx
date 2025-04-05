import { ColumnDef } from "@tanstack/react-table";
import { CategoriesTabelCellAction } from "./categories-table-cell-action";

export interface CategoryColumn {
  id: string;
  name: string;
  billboardName: string;
  createdAt: string;
}

export const CategoryColumn: ColumnDef<CategoryColumn>[] = [
  {
    accessorKey: "name",
    header: "Nome",
  },
  {
    accessorKey: "billboardName",
    header: "Painel",
  },
  {
    accessorKey: "createdAt",
    header: "Data de criação",
  },
  {
    id: "actions",
    cell: ({ row }) => <CategoriesTabelCellAction data={row.original} />,
  },
];
