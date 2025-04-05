import { api } from "@/lib/axios";

export interface UpdateColorRequest {
  name: string;
  hex: string;
  storeId: string;
  colorId: string;
}

export async function updateColor({
  name,
  hex,
  storeId,
  colorId,
}: UpdateColorRequest) {
  await api.put(`/stores/${storeId}/colors/${colorId}`, {
    name,
    hex,
  });
}
