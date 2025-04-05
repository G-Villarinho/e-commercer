import { GetAllBillboardsResponse } from "@/api/get-all-billboards";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { zodResolver } from "@hookform/resolvers/zod";
import { Search, X } from "lucide-react";
import { useForm, Controller } from "react-hook-form";
import { useSearchParams } from "react-router-dom";
import { z } from "zod";

const categoriesFilterSchema = z.object({
  name: z.string().optional(),
  billboardId: z.string().optional(),
});

type CategoriesFilter = z.infer<typeof categoriesFilterSchema>;

interface CategoriesTableFilterProps {
  billboards: GetAllBillboardsResponse[];
}

export function CategoriesTableFilter({
  billboards,
}: CategoriesTableFilterProps) {
  const [searchParams, setSearchParams] = useSearchParams();

  const name = searchParams.get("name") || "";
  const billboardId = searchParams.get("billboardId") || "";

  const { register, handleSubmit, reset, control } = useForm<CategoriesFilter>({
    resolver: zodResolver(categoriesFilterSchema),
    defaultValues: {
      name,
      billboardId,
    },
  });

  function handleFilter(data: CategoriesFilter) {
    setSearchParams((prev) => {
      if (data.name) {
        prev.set("name", data.name);
      } else {
        prev.delete("name");
      }

      if (data.billboardId) {
        prev.set("billboardId", data.billboardId);
      } else {
        prev.delete("billboardId");
      }

      prev.set("page", "1");

      return prev;
    });
  }

  function handleClearFilters() {
    setSearchParams((prev) => {
      prev.delete("name");
      prev.delete("billboardId");
      prev.set("page", "1");
      return prev;
    });
    reset();
  }

  const hasAnyFilter = !!name || !!billboardId;

  return (
    <form
      onSubmit={handleSubmit(handleFilter)}
      className="flex items-center gap-2"
    >
      <span className="text-sm font-semibold">Filtros:</span>

      <Input placeholder="Nome" className=" w-auto" {...register("name")} />

      <Controller
        name="billboardId"
        control={control}
        render={({ field }) => (
          <Select
            onValueChange={(value) =>
              field.onChange(value === "all" ? undefined : value)
            }
            defaultValue={field.value || "all"}
          >
            <SelectTrigger className="w-72">
              <SelectValue placeholder="Filtrar por painel" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">Todos os painéis</SelectItem>
              {billboards.map((billboard) => (
                <SelectItem key={billboard.id} value={billboard.id}>
                  {billboard.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        )}
      />

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
