import { api } from "@/lib/axios";

export interface DeleteColorParams {
  storeId: string;
  colorId: string;
}

export async function deleteColor({ storeId, colorId }: DeleteColorParams) {
  await api.delete(`/stores/${storeId}/colors/${colorId}`);
}
