import { SizeResponse } from "@/api/responses/size-response";
import { updateSize } from "@/api/update-size";
import { InputError } from "@/components/input-error";
import { LoadingButton } from "@/components/loading-button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { useNavigate, useParams } from "react-router-dom";
import { z } from "zod";

const editSizeSchema = z.object({
  name: z.string().nonempty("O nome do tamanho não pode estar vazio."),
  value: z.string().nonempty("O valor do tamanho não pode estar vazio."),
});

type EditSizeSchema = z.infer<typeof editSizeSchema>;

export interface EditSizeFormProps {
  size: SizeResponse;
}

export function EditSizeForm({ size }: EditSizeFormProps) {
  const { storeId } = useParams();
  const navigate = useNavigate();

  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm<EditSizeSchema>({
    resolver: zodResolver(editSizeSchema),
    defaultValues: {
      name: size.name,
      value: size.value,
    },
  });

  const { mutateAsync: updateSizeFn } = useMutation({
    mutationFn: updateSize,
  });

  async function handleUpdateSize(data: EditSizeSchema) {
    try {
      if (!storeId) {
        throw new Error("O ID da loja é obrigatório.");
      }

      await updateSizeFn({
        name: data.name,
        value: data.value,
        storeId,
        sizeId: size.id,
      });
      toast.success("Tamanho criado com sucesso.");
      navigate(`/${storeId}/sizes`);
    } catch (error) {
      toast.error("Não foi possível criar o tamanho.");
    }
  }

  return (
    <form onSubmit={handleSubmit(handleUpdateSize)}>
      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-1">
          <Label htmlFor="name">Nome da tamanho</Label>
          <Input
            id="name"
            placeholder="Ex: Pequeno, Médio, Grande"
            disabled={isSubmitting}
            {...register("name")}
          />
          {errors.name && <InputError error={errors.name.message} />}
        </div>
        <div className="space-y-1">
          <Label htmlFor="name">Valor do tamanho</Label>
          <Input
            id="value"
            placeholder="Ex: P, M, G"
            disabled={isSubmitting}
            {...register("value")}
          />
          {errors.value && <InputError error={errors.value.message} />}
        </div>
      </div>
      <LoadingButton isLoading={isSubmitting} type="submit" className="mt-6">
        Criar tamanho
      </LoadingButton>
    </form>
  );
}
