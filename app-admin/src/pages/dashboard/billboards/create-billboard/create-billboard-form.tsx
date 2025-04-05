import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { InputError } from "@/components/input-error";
import { LoadingButton } from "@/components/loading-button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ImageUpload } from "@/components/image-upload";
import toast from "react-hot-toast";
import { useNavigate, useParams, useSearchParams } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import { createBillboard } from "@/api/create-billboard";
import { isAxiosError } from "axios";

const createBillboardSchema = z.object({
  label: z.string().nonempty("O rótulo da loja não pode estar vazio."),
  image: z.instanceof(File, { message: "Adicione uma imagem válida." }),
});

type CreateBillboardSchema = z.infer<typeof createBillboardSchema>;

export function CreateBillboardForm() {
  const { storeId } = useParams();
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  const redirectUrl = searchParams.get("redirectUrl");

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isSubmitting },
  } = useForm<CreateBillboardSchema>({
    resolver: zodResolver(createBillboardSchema),
    defaultValues: {
      label: "",
      image: undefined as unknown as File,
    },
  });

  const image = watch("image");

  function handleSetImage(file: File) {
    setValue("image", file);
  }

  function handleRemoveImage() {
    setValue("image", undefined as unknown as File);
  }

  const { mutateAsync: createBillboardFn } = useMutation({
    mutationFn: createBillboard,
  });

  async function handleCreateBillboard(data: CreateBillboardSchema) {
    try {
      if (!storeId) {
        toast.error("O ID da loja é obrigatório.");
        return;
      }

      await createBillboardFn({
        label: data.label,
        image: data.image,
        storeId,
      });

      toast.success("Painel criado com sucesso.");

      if (redirectUrl) {
        navigate(redirectUrl);
        return;
      }

      navigate(`/${storeId}/billboards`);
    } catch (errors) {
      if (isAxiosError(errors)) {
        toast.error(errors.response?.data.message);
      }
    }
  }

  return (
    <form onSubmit={handleSubmit(handleCreateBillboard)} className="space-y-6">
      <div>
        <Label>Imagem do Painel</Label>
        <ImageUpload
          value={image ? [URL.createObjectURL(image)] : []}
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
          <Input
            id="label"
            placeholder="Ex: Painel 1"
            disabled={isSubmitting}
            {...register("label")}
          />
          {errors.label && <InputError error={errors.label.message} />}
        </div>
      </div>

      <LoadingButton isLoading={isSubmitting} type="submit">
        Criar painel
      </LoadingButton>
    </form>
  );
}
