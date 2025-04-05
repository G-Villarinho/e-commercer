import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { Search, X } from "lucide-react";
import { useForm } from "react-hook-form";
import { useSearchParams } from "react-router-dom";
import { z } from "zod";

const colorsFilterSchema = z.object({
  name: z.string().optional(),
});

type ColorsFilterSchema = z.infer<typeof colorsFilterSchema>;

export function ColorsTableFilter() {
  const [searchParams, setSearchParams] = useSearchParams();

  const name = searchParams.get("name") || "";

  const { register, handleSubmit, reset } = useForm<ColorsFilterSchema>({
    resolver: zodResolver(colorsFilterSchema),
    defaultValues: {
      name,
    },
  });

  function handleFilter(data: ColorsFilterSchema) {
    setSearchParams((prev) => {
      if (data.name) {
        prev.set("name", data.name);
      } else {
        prev.delete("name");
      }

      prev.set("page", "1");

      return prev;
    });
  }

  function handleClearFilters() {
    setSearchParams((prev) => {
      prev.delete("name");
      prev.set("page", "1");
      return prev;
    });
    reset();
  }

  const hasAnyFilter = !!name;

  return (
    <form
      onSubmit={handleSubmit(handleFilter)}
      className="flex items-center gap-2"
    >
      <span className="text-sm font-semibold">Filtros:</span>

      <Input placeholder="Nome" className=" w-auto" {...register("name")} />

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
