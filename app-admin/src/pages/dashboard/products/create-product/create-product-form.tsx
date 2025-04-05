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
import { LoadingButton } from "@/components/loading-button";
import toast from "react-hot-toast";
import { useNavigate, useParams } from "react-router-dom";
import { GetAllCategoriesResponse } from "@/api/get-all-categories";
import { GetAllSizesResponse } from "@/api/get-all-sizes";
import { GetAllColorsResponse } from "@/api/get-all-colors";
import { Checkbox } from "@/components/ui/checkbox";
import { useMutation } from "@tanstack/react-query";
import { createProduct } from "@/api/create-product";
import { ImageUpload } from "@/components/image-upload";
import { useState } from "react";

const createProductSchema = z
  .object({
    name: z.string().min(1, "O nome é obrigatório."),
    price: z.coerce
      .number({
        invalid_type_error: "Por favor, insira um número válido",
      })
      .positive("O preço deve ser maior que zero."),
    categoryId: z.string().nonempty("A categoria é obrigatória."),
    sizeId: z.string().nonempty("O tamanho é obrigatório."),
    colorId: z.string().nonempty("A cor é obrigatória."),
    isFeatured: z.boolean().default(false),
    isArchived: z.boolean().default(false),
    images: z
      .array(z.instanceof(File), { message: "Adicione uma imagem válida." })
      .min(1, "Adicione pelo menos uma imagem."),
  })
  .superRefine((data, ctx) => {
    if (!data.isFeatured && !data.isArchived) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Selecione pelo menos uma opção: Destaque ou Arquivado",
        path: ["isFeatured"],
      });
    }

    if (data.isFeatured && data.isArchived) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message:
          "Não é possível selecionar Destaque e Arquivado simultaneamente",
        path: ["isArchived"],
      });
    }
  });

type CreateProductSchema = z.infer<typeof createProductSchema>;

interface CreateProductFormProps {
  categories: GetAllCategoriesResponse[];
  sizes: GetAllSizesResponse[];
  colors: GetAllColorsResponse[];
}

export function CreateProductForm({
  categories,
  sizes,
  colors,
}: CreateProductFormProps) {
  const { storeId } = useParams();
  const navigate = useNavigate();
  const [files, setFiles] = useState<File[]>([]);

  const {
    register,
    control,
    handleSubmit,
    setValue,
    formState: { isSubmitting, errors },
  } = useForm<CreateProductSchema>({
    resolver: zodResolver(createProductSchema),
    defaultValues: {
      categoryId: "",
      sizeId: "",
      colorId: "",
    },
  });

  const { mutateAsync: createProductFn } = useMutation({
    mutationFn: createProduct,
    onError: () => toast.error("Erro ao criar produto"),
    onSuccess: () => {
      toast.success("Produto criado com sucesso!");
      navigate(`/${storeId}/products`);
    },
  });

  function handleAddImage(newFile: File) {
    const updatedFiles = [...files, newFile].slice(0, 5);
    setFiles(updatedFiles);
    if (updatedFiles.length > 0) {
      setValue("images", updatedFiles as [File, ...File[]], {
        shouldValidate: true,
      });
    }
  }

  function handleRemoveImage(index: number) {
    const updatedFiles = files.filter((_, i) => i !== index);
    setFiles(updatedFiles);
    setValue("images", updatedFiles, { shouldValidate: true });
  }

  async function handleCreateProduct(data: CreateProductSchema) {
    if (!storeId) {
      return;
    }

    try {
      await createProductFn({
        storeId,
        ...data,
        images: files,
      });
      toast(
        () => (
          <div className="flex items-center gap-2">
            ⏳ As imagens estão sendo processadas em segundo plano...
          </div>
        ),
        {
          duration: 4000,
        }
      );
    } catch (error) {
      console.error("Erro ao criar produto:", error);
    }
  }

  return (
    <form
      onSubmit={handleSubmit(handleCreateProduct)}
      className="space-y-8 w-full"
    >
      <div>
        <Label>Imagens do produto</Label>
        <ImageUpload
          value={files.map((file) => URL.createObjectURL(file))}
          onChange={handleAddImage}
          onRemove={handleRemoveImage}
          disabled={isSubmitting || files.length >= 5}
        />
        {errors.images && <InputError error={errors.images.message} />}

        <p className="text-sm text-muted-foreground mt-2">
          {files.length} de 5 imagens selecionadas
        </p>
      </div>

      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-1">
          <Label htmlFor="name">Nome do Produto</Label>
          <Input
            id="name"
            placeholder="Ex: Camiseta Branca"
            disabled={isSubmitting}
            autoFocus
            {...register("name")}
          />
          {errors.name && <InputError error={errors.name.message} />}
        </div>

        <div className="space-y-1">
          <Label htmlFor="price">Preço</Label>
          <Input
            id="price"
            type="number"
            step="0.01"
            placeholder="Ex: 99.90"
            disabled={isSubmitting}
            {...register("price")}
          />
          {errors.price && <InputError error={errors.price.message} />}
        </div>

        <div className="space-y-1">
          <Label htmlFor="categoryId">Categoria</Label>
          <Controller
            name="categoryId"
            control={control}
            render={({ field }) => (
              <>
                <Select
                  onValueChange={field.onChange}
                  value={field.value}
                  disabled={isSubmitting}
                >
                  <SelectTrigger id="categoryId" className="w-full">
                    <SelectValue placeholder="Selecione uma categoria" />
                  </SelectTrigger>
                  <SelectContent>
                    {categories.map((category) => (
                      <SelectItem key={category.id} value={category.id}>
                        {category.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                {errors.categoryId && (
                  <InputError error={errors.categoryId.message} />
                )}
              </>
            )}
          />
        </div>

        <div className="space-y-1">
          <Label htmlFor="sizeId">Tamanho</Label>
          <Controller
            name="sizeId"
            control={control}
            render={({ field }) => (
              <>
                <Select
                  onValueChange={field.onChange}
                  value={field.value}
                  disabled={isSubmitting}
                >
                  <SelectTrigger id="sizeId" className="w-full">
                    <SelectValue placeholder="Selecione um tamanho" />
                  </SelectTrigger>
                  <SelectContent>
                    {sizes.map((size) => (
                      <SelectItem key={size.id} value={size.id}>
                        {size.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                {errors.sizeId && <InputError error={errors.sizeId.message} />}
              </>
            )}
          />
        </div>

        <div className="space-y-1">
          <Label htmlFor="colorId">Cor</Label>
          <Controller
            name="colorId"
            control={control}
            render={({ field }) => (
              <>
                <Select
                  onValueChange={field.onChange}
                  value={field.value}
                  disabled={isSubmitting}
                >
                  <SelectTrigger id="colorId" className="w-full">
                    <SelectValue placeholder="Selecione uma cor" />
                  </SelectTrigger>
                  <SelectContent>
                    {colors.map((color) => (
                      <SelectItem key={color.id} value={color.id}>
                        <div className="flex items-center gap-2">
                          <div
                            className="w-4 h-4 rounded-full border"
                            style={{ backgroundColor: color.hex }}
                          />
                          {color.name}
                        </div>
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                {errors.colorId && (
                  <InputError error={errors.colorId.message} />
                )}
              </>
            )}
          />
        </div>
      </div>

      <div className="flex gap-4">
        {/* Checkbox de Destaque */}
        <div className="rounded-md border p-4 w-fit min-w-[200px]">
          <div className="flex items-center gap-3">
            <Controller
              name="isFeatured"
              control={control}
              render={({ field, fieldState }) => (
                <>
                  <Checkbox
                    id="isFeatured"
                    checked={field.value}
                    onCheckedChange={field.onChange}
                    className="h-5 w-5"
                  />
                  <div className="space-y-1 leading-none">
                    <Label
                      htmlFor="isFeatured"
                      className="font-normal cursor-pointer"
                    >
                      Destaque
                    </Label>
                    <p className="text-sm text-muted-foreground">
                      Aparecerá na página inicial
                    </p>
                    {fieldState.error && (
                      <InputError error={fieldState.error.message} />
                    )}
                  </div>
                </>
              )}
            />
          </div>
        </div>

        {/* Checkbox de Arquivado */}
        <div className="rounded-md border p-4 w-fit min-w-[200px]">
          <div className="flex items-center gap-3">
            <Controller
              name="isArchived"
              control={control}
              render={({ field, fieldState }) => (
                <>
                  <Checkbox
                    id="isArchived"
                    checked={field.value}
                    onCheckedChange={field.onChange}
                    className="h-5 w-5"
                  />
                  <div className="space-y-1 leading-none">
                    <Label
                      htmlFor="isArchived"
                      className="font-normal cursor-pointer"
                    >
                      Arquivado
                    </Label>
                    <p className="text-sm text-muted-foreground">
                      Não visível para clientes
                    </p>
                    {fieldState.error && (
                      <InputError error={fieldState.error.message} />
                    )}
                  </div>
                </>
              )}
            />
          </div>
        </div>
      </div>

      <LoadingButton isLoading={isSubmitting} type="submit">
        Criar Produto
      </LoadingButton>
    </form>
  );
}
