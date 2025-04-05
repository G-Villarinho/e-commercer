import { api } from "@/lib/axios";

export interface DeleteStoreParams {
    storeId: string;
}

export async function deleteStore({ storeId }: DeleteStoreParams) {
    await api.delete(`/stores/${storeId}`);
}