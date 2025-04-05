import { api } from "@/lib/axios";
import { StoreResponse } from "@/api/responses/store-response";

export async function getUserFirstStore() {
  const response = await api.get<StoreResponse>("/me/stores/first");

  return response.data;
}
