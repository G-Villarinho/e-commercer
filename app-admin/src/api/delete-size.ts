import { api } from "@/lib/axios";

export interface DeleteSizeParams {
    storeId: string;
    sizeId: string;
}

export async function deleteSize({ storeId, sizeId }: DeleteSizeParams) {
    await api.delete(`stores/${storeId}/sizes/${sizeId}`);
}