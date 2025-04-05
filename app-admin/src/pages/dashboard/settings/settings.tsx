import { StoreResponse } from "@/api/responses/store-response";
import { useQueryClient } from "@tanstack/react-query";
import { Navigate, useParams } from "react-router-dom";
import { SettingsForm } from "./settings-form";
import { Helmet } from "react-helmet-async";

export function Settings() {
  const { storeId } = useParams();
  const queryClient = useQueryClient();

  const store = queryClient.getQueryData<StoreResponse>(["store", storeId]);

  if (!store || store === undefined) {
    return <Navigate to="/" />;
  }

  return (
    <>
      <Helmet title="Configurações" />
      {store && <SettingsForm initialData={store} />}
    </>
  );
}
