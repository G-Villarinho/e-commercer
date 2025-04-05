import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { Search, X } from "lucide-react";
import { useForm } from "react-hook-form";
import { useSearchParams } from "react-router-dom";
import { z } from "zod";

const sizesFilterSchema = z.object({
  name: z.string().optional(),
});

type SizesFilterSchema = z.infer<typeof sizesFilterSchema>;

export function SizesTableFilter() {
  const [searchParams, setSearchParams] = useSearchParams();

  const label = searchParams.get("label");

  const { register, handleSubmit, reset } = useForm<SizesFilterSchema>({
    resolver: zodResolver(sizesFilterSchema),
    defaultValues: {
      name: label || "",
    },
  });

  function handleFilter(data: SizesFilterSchema) {
    const label = data.name?.toString();

    setSearchParams((prev) => {
      if (label) {
        prev.set("label", label);
      } else {
        prev.delete("label");
      }

      prev.set("page", "1");

      return prev;
    });
  }

  function handleClearFilters() {
    setSearchParams((prev) => {
      prev.delete("label");
      prev.set("page", "1");

      return prev;
    });
    reset();
  }

  const hasAnyFilter = !!label;

  return (
    <form
      onSubmit={handleSubmit(handleFilter)}
      className="flex items-center gap-2"
    >
      <span className="text-sm font-semibold">Filtros:</span>
      <Input placeholder="Nome" className="w-auto" {...register("name")} />
      <Button type="submit" variant="secondary" size="sm">
        <Search className="mr-2 h-4 w-4" />
        Filtrar resultados
      </Button>
      <Button
        type="button"
        variant="outline"
        size="sm"
        disabled={!hasAnyFilter}
        onClick={handleClearFilters}
      >
        <X className="mr-2 h-4 w-4" />
        Remover filtros
      </Button>
    </form>
  );
}
