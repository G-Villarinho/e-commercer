import { z } from "zod";
import { StoreResponse } from "@/api/responses/store-response";
import { zodResolver } from "@hookform/resolvers/zod";
import { Trash } from "lucide-react";
import { useForm } from "react-hook-form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { updateStore } from "@/api/update-store";
import { isAxiosError } from "axios";
import toast from "react-hot-toast";
import { useState } from "react";
import { deleteStore } from "@/api/delete-store";
import { useNavigate, useParams } from "react-router-dom";

import { Heading } from "@/components/heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { InputError } from "@/components/input-error";
import { LoadingButton } from "@/components/loading-button";
import { AlertModal } from "@/components/modals/alert-modal";
import { ApiAlert } from "@/components/api-alert";
import { useOrigin } from "@/hooks/use-origin";

const settingsSchema = z.object({
  name: z.string().nonempty("O nome da loja não pode estar vazio."),
});

type SettingsSchema = z.infer<typeof settingsSchema>;

interface SettingsFormProps {
  initialData: StoreResponse;
}

export function SettingsForm({ initialData }: SettingsFormProps) {
  const [isAlertOpen, setIsAlertOpen] = useState(false);
  const queryClient = useQueryClient();
  const naviagate = useNavigate();
  const { storeId } = useParams();
  const origin = useOrigin();

  const {
    handleSubmit,
    register,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<SettingsSchema>({
    resolver: zodResolver(settingsSchema),
    defaultValues: {
      name: initialData.name,
    },
  });

  function updateStoreNameOnCache(newName: string) {
    const storesCache = queryClient.getQueriesData<StoreResponse[]>({
      queryKey: ["stores"],
    });

    storesCache.forEach(([cacheKey, cachedStores]) => {
      if (!cachedStores) {
        return;
      }

      queryClient.setQueryData<StoreResponse[]>(
        cacheKey,
        cachedStores.map((store) =>
          store.id === initialData.id ? { ...store, name: newName } : store
        )
      );
    });

    queryClient.setQueryData<StoreResponse>(
      ["store", initialData.id],
      (cachedStore) => {
        if (!cachedStore) return cachedStore;
        return { ...cachedStore, name: newName };
      }
    );
  }

  const { mutateAsync: updateStoreFn } = useMutation({
    mutationFn: updateStore,
  });

  async function handleUpdateStore(data: SettingsSchema) {
    try {
      await updateStoreFn({
        storeId: initialData.id,
        name: data.name,
      });

      updateStoreNameOnCache(data.name);

      toast.success("Loja atualizada com sucesso.");
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response?.status === 409) {
          setError("name", {
            type: "manual",
            message: "Já existe uma loja com esse nome.",
          });
        }
      }
    }
  }

  const { mutateAsync: deleteStoreFn, isPending } = useMutation({
    mutationFn: deleteStore,
  });

  async function handleDeleteStore() {
    try {
      await deleteStoreFn({ storeId: initialData.id });
      queryClient.removeQueries({ queryKey: ["userFirstStore"] });
      queryClient.removeQueries({ queryKey: ["stores"] });
      queryClient.removeQueries({ queryKey: ["store", initialData.id] });
      toast.success("Loja excluída.");
      naviagate("/");
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response?.status === 409) {
          toast.error(
            "Tenha certeza de que você removeu todos os produtos e categorias antes."
          );
        }
      }
    }
  }

  return (
    <>
      <AlertModal
        isOpen={isAlertOpen}
        onClose={() => setIsAlertOpen(false)}
        onConfirm={handleDeleteStore}
        loading={isPending}
      />
      <div className="flex items-center justify-between">
        <Heading
          title="Configurações"
          description="Gerenciar preferências da loja"
        />

        <Button
          variant="destructive"
          size="sm"
          onClick={() => setIsAlertOpen(true)}
        >
          <Trash className="h-4 w-4 text-white" />
        </Button>
      </div>
      <Separator />
      <form onSubmit={handleSubmit(handleUpdateStore)}>
        <div className="grid grid-cols-3 gap-8">
          <div className="space-y-1">
            <Label htmlFor="name">Nome da loja</Label>
            <Input
              id="name"
              placeholder="Nome da loja"
              disabled={isSubmitting}
              {...register("name")}
            />
            {errors.name && <InputError error={errors.name.message} />}
          </div>
        </div>
        <LoadingButton isLoading={isSubmitting} type="submit" className="mt-6">
          Salvar alterações
        </LoadingButton>
      </form>
      <Separator />

      <ApiAlert
        title="FAST_BUY_PUBLIC_API_URL"
        description={`${origin}/api/stores/${storeId}`}
      />
    </>
  );
}
