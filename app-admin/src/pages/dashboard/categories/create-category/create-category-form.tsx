import { InputError } from "@/components/input-error";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import { z } from "zod";

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button"; // Adicionado para o botão
import { useNavigate, useParams } from "react-router-dom";
import { LoadingButton } from "@/components/loading-button";
import { useMutation } from "@tanstack/react-query";
import { createCategory } from "@/api/create-category";
import toast from "react-hot-toast";

const createCategorySchema = z.object({
  name: z.string().nonempty("O nome da categoria não pode estar vazio."),
  billboardId: z.string().nonempty("Selecione um painel."),
});

type CreateCategorySchema = z.infer<typeof createCategorySchema>;

export interface CreateCategoryFormProps {
  billboards: {
    id: string;
    label: string;
  }[];
}

export function CreateCategoryForm({ billboards }: CreateCategoryFormProps) {
  const { storeId } = useParams();
  const navigate = useNavigate();

  const {
    register,
    control,
    handleSubmit,
    formState: { isSubmitting, errors },
  } = useForm<CreateCategorySchema>({
    resolver: zodResolver(createCategorySchema),
    defaultValues: {
      name: "",
      billboardId: "",
    },
  });

  function handleCreateBillboard() {
    if (!storeId) {
      throw new Error("O ID da loja é obrigatório.");
    }

    navigate(
      `/${storeId}/billboards/new?redirectUrl=/${storeId}/categories/new`
    );
  }

  const { mutateAsync: createCategoryFn } = useMutation({
    mutationFn: createCategory,
  });

  async function handleCreateCategory(data: CreateCategorySchema) {
    try {
      if (!storeId) {
        throw new Error("O ID da loja é obrigatório.");
      }

      await createCategoryFn({
        storeId: storeId,
        ...data,
      });
      navigate(`/${storeId}/categories`);
    } catch (error) {
      toast.error("Erro ao criar categoria");
    }
  }

  return (
    <form
      onSubmit={handleSubmit(handleCreateCategory)}
      className="space-y-8 w-full"
    >
      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-1">
          <Label htmlFor="name">Nome da categoria</Label>
          <Input
            id="name"
            placeholder="Ex: Eletrônicos"
            disabled={isSubmitting}
            autoFocus
            {...register("name")}
          />
          {errors.name && <InputError error={errors.name.message} />}
        </div>

        <div className="space-y-1">
          <Label htmlFor="billboardId">Painel</Label>
          <Controller
            name="billboardId"
            control={control}
            render={({ field }) => (
              <Select
                onValueChange={field.onChange}
                defaultValue={field.value}
                disabled={isSubmitting}
              >
                <SelectTrigger id="billboardId" className="w-full">
                  <SelectValue placeholder="Selecione um painel" />
                </SelectTrigger>
                <SelectContent>
                  {billboards.length > 0 ? (
                    billboards.map((billboard) => (
                      <SelectItem key={billboard.id} value={billboard.id}>
                        {billboard.label}
                      </SelectItem>
                    ))
                  ) : (
                    <div className="p-2 text-center">
                      <p className="text-muted-foreground mb-2">
                        Nenhum painel disponível
                      </p>
                      <Button
                        variant="outline"
                        className="w-full"
                        size="xs"
                        onClick={handleCreateBillboard}
                      >
                        Criar novo painel
                      </Button>
                    </div>
                  )}
                </SelectContent>
              </Select>
            )}
          />
          {errors.billboardId && (
            <InputError error={errors.billboardId.message} />
          )}
        </div>
      </div>
      <LoadingButton isLoading={isSubmitting} type="submit">
        Criar categoria
      </LoadingButton>
    </form>
  );
}
