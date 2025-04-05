import { ColumnDef } from "@tanstack/react-table";
import { ColorsTabelCellAction } from "./colors-table-cell-action";

export interface ColorColumn {
  id: string;
  name: string;
  hex: string;
  createdAt: string;
}

export const ColorColumn: ColumnDef<ColorColumn>[] = [
  {
    accessorKey: "name",
    header: "Nome",
  },
  {
    accessorKey: "hex",
    header: "Valor",
    cell: ({ row }) => (
      <div className="flex items-center gap-x-2">
        {row.original.hex}
        <div
          className="w-5 h-5 rounded-full border"
          style={{ backgroundColor: row.original.hex }}
        />
      </div>
    ),
  },
  {
    accessorKey: "createdAt",
    header: "Data de criação",
  },
  {
    id: "actions",
    cell: ({ row }) => <ColorsTabelCellAction data={row.original} />,
  },
];
