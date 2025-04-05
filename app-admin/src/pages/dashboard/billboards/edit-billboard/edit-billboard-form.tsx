import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { InputError } from "@/components/input-error";
import { LoadingButton } from "@/components/loading-button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ImageUpload } from "@/components/image-upload";
import toast from "react-hot-toast";
import { useNavigate, useParams } from "react-router-dom";
import { isAxiosError } from "axios";
import { BillboardResponse } from "@/api/responses/billboard-response";
import { useEffect } from "react";
import { updateBillboard } from "@/api/update-billboard";
import { useMutation } from "@tanstack/react-query";

const editBillboardSchema = z.object({
  label: z.string().nonempty("O rótulo da loja não pode estar vazio."),
  image: z
    .instanceof(File, { message: "Adicione uma imagem válida." })
    .optional(),
});

type EditBillboardSchema = z.infer<typeof editBillboardSchema>;

interface EditBillboardFormProps {
  data: BillboardResponse;
}

export function EditBillboardForm({ data }: EditBillboardFormProps) {
  const { storeId } = useParams();
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isSubmitting },
  } = useForm<EditBillboardSchema>({
    resolver: zodResolver(editBillboardSchema),
    defaultValues: {
      label: data.label,
      image: undefined as unknown as File,
    },
  });

  useEffect(() => {
    setValue("label", data.label);
  }, [data, setValue]);

  const image = watch("image");

  function handleSetImage(file: File) {
    setValue("image", file);
  }

  function handleRemoveImage() {
    setValue("image", undefined as unknown as File);
  }

  const { mutateAsync: updateBillboardFn } = useMutation({
    mutationFn: updateBillboard,
  });

  async function handleEditBillboard(values: EditBillboardSchema) {
    try {
      if (!storeId) {
        toast.error("O ID da loja é obrigatório.");
        return;
      }

      await updateBillboardFn({
        label: values.label,
        image: values.image as File,
        storeId,
        billboardId: data.id,
      });

      toast.success("Painel atualizado com sucesso.");
      navigate(`/${storeId}/billboards`);
    } catch (errors) {
      if (isAxiosError(errors)) {
        toast.error(errors.response?.data.message);
      }
    }
  }

  return (
    <form onSubmit={handleSubmit(handleEditBillboard)} className="space-y-6">
      <div>
        <Label>Imagem do Painel</Label>
        <ImageUpload
          value={image ? [URL.createObjectURL(image)] : [String(data.imageUrl)]}
          onChange={handleSetImage}
          onRemove={handleRemoveImage}
          disabled={isSubmitting}
          maxImages={1}
        />
        {errors.image && <InputError error={errors.image.message} />}
      </div>

      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-1">
          <Label htmlFor="label">Rótulo do painel</Label>
          <Input id="label" disabled={isSubmitting} {...register("label")} />
          {errors.label && <InputError error={errors.label.message} />}
        </div>
      </div>

      <LoadingButton isLoading={isSubmitting} type="submit">
        Atualizar painel
      </LoadingButton>
    </form>
  );
}
