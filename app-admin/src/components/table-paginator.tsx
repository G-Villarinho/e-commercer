import { Button } from "@/components/ui/button";

interface TablePaginatorProps {
  pageIndex: number;
  totalPages: number;
  onPageChange: (newPageIndex: number) => void;
}

export function TablePaginator({
  pageIndex,
  totalPages,
  onPageChange,
}: TablePaginatorProps) {
  return (
    <div className="flex justify-end gap-4 mt-4">
      <Button
        onClick={() => onPageChange(pageIndex - 1)}
        disabled={pageIndex <= 1}
        variant="outline"
      >
        Anterior
      </Button>
      <Button
        onClick={() => onPageChange(pageIndex + 1)}
        disabled={totalPages <= pageIndex}
        variant="outline"
      >
        Próximo
      </Button>
    </div>
  );
}
