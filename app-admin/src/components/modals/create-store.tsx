import { Modal } from "@/components/ui/modal";
import { useStoreModal } from "@/hooks/use-store-modal";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import { InputError } from "../input-error";
import { useMutation } from "@tanstack/react-query";
import { createStore } from "@/api/create-store";
import { isAxiosError } from "axios";
import { useNavigate } from "react-router-dom";

const createStoreSchema = z.object({
  name: z.string().nonempty("O nome da loja não pode estar vazio."),
});

type CreateStoreSchema = z.infer<typeof createStoreSchema>;

export function CreateStoreModal() {
  const navigate = useNavigate();
  const storeModal = useStoreModal();

  const {
    handleSubmit,
    register,
    setError,
    formState: { errors },
  } = useForm<CreateStoreSchema>({
    resolver: zodResolver(createStoreSchema),
  });

  const { mutate: createStoreFn, isPending } = useMutation({
    mutationFn: createStore,
    onSuccess: (response) => {
      storeModal.onClose();
      navigate(`/${response.storeId}`);
    },
    onError: (error) => {
      if (isAxiosError(error) && error.response?.status === 409) {
        setError("name", {
          type: "manual",
          message: "Já existe uma loja com esse nome.",
        });
      }
    },
  });

  function handleCreateStore({ name }: CreateStoreSchema) {
    createStoreFn({ name });
  }

  return (
    <Modal
      title="Crie uma nova loja"
      description="Adicione uma nova loja para gerenciar produtos e categorias"
      isOpen={storeModal.isOpen}
      onClose={storeModal.onClose}
    >
      <div className="space-y-4 py-2 pb-4">
        <form onSubmit={handleSubmit(handleCreateStore)}>
          <div className="space-y-2">
            <Label className="">Name</Label>
            <Input placeholder="E-Commercer..." {...register("name")} />
            {errors.name && <InputError error={errors.name.message} />}
          </div>
          <div className="pt-6 space-x-2 flex items-center justify-end">
            <Button
              type="button"
              variant="secondary"
              onClick={() => storeModal.onClose(true)}
            >
              Cancelar
            </Button>
            <Button type="submit" disabled={isPending}>
              {isPending ? "Criando..." : "Criar loja"}
            </Button>
          </div>
        </form>
      </div>
    </Modal>
  );
}
