import { api } from "@/lib/axios";

export interface CreateSizeRequest {
    name: string;
    value: string;
    storeId: string;
}

export async function createSize({ name, value, storeId }: CreateSizeRequest) {
    await api.post(`stores/${storeId}/sizes`, { name, value });
}