import { createColor } from "@/api/create-color";
import { InputError } from "@/components/input-error";
import { LoadingButton } from "@/components/loading-button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { z } from "zod";

const createStoreSchema = z.object({
  name: z.string().nonempty("O nome da categoria não pode estar vazio."),
  hex: z
    .string()
    .nonempty("O valor do tamanho não pode estar vazio.")
    .regex(
      /^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$/,
      "O valor deve ser um código hexadecimal válido."
    ),
});

type CreateColorSchema = z.infer<typeof createStoreSchema>;

export interface CreateColorFormProps {
  storeId: string;
}

export function CreateColorForm({ storeId }: CreateColorFormProps) {
  const navigate = useNavigate();

  const {
    register,
    watch,
    setError,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<CreateColorSchema>({
    resolver: zodResolver(createStoreSchema),
  });

  const { mutateAsync: createColorFn } = useMutation({
    mutationFn: createColor,
  });

  async function handleCreateColor(data: CreateColorSchema) {
    try {
      await createColorFn({ storeId, ...data });
      toast.success("Cor criada com sucesso.");
      navigate(`/${storeId}/colors`);
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.status === 409) {
          setError("hex", {
            type: "manual",
            message: "Já existe uma cor com este valor nessa loja.",
          });
        }
      }
    }
  }

  const hexValue = watch("hex", "#000000");

  return (
    <form
      onSubmit={handleSubmit(handleCreateColor)}
      className="space-y-8 w-full"
    >
      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-1">
          <Label htmlFor="name">Nome da cor</Label>
          <Input
            id="name"
            placeholder="Ex: Azul, Vermelho, Amarelo"
            disabled={isSubmitting}
            autoFocus
            {...register("name")}
          />
          {errors.name && <InputError error={errors.name.message} />}
        </div>
        <div className="space-y-1">
          <Label htmlFor="hex">Valor da cor</Label>
          <Input
            id="hex"
            placeholder="Ex: #000000, #FFFFFF, #FF0000"
            disabled={isSubmitting}
            {...register("hex")}
          />
          {errors.hex && <InputError error={errors.hex.message} />}
        </div>
        <div
          className="border p-4 rounded-full w-8 h-8 mt-5"
          style={{ backgroundColor: hexValue }}
        />
      </div>
      <LoadingButton isLoading={isSubmitting} type="submit">
        Criar cor
      </LoadingButton>
    </form>
  );
}
