import { api } from "@/lib/axios";

export interface UpdateSizeRequest {
    name: string;
    value: string;
    storeId: string;
    sizeId: string;
}

export async function updateSize({ name, value, storeId, sizeId }: UpdateSizeRequest) {
    await api.put(`stores/${storeId}/sizes/${sizeId}`, { name, value });
}