import { api } from "@/lib/axios";
import { StoreResponse } from "./responses/store-response";

export async function getUserStores() {
  const response = await api.get<StoreResponse[]>("/me/stores");
  return response.data;
}
