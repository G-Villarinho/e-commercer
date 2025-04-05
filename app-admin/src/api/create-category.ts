import { api } from "@/lib/axios";

export interface CreateCategoryRequest {
    name: string;
    billboardId: string;
    storeId: string;
}

export async function createCategory({ name, billboardId, storeId }: CreateCategoryRequest) {
    await api.post(`stores/${storeId}/categories`, { name, billboardId });
}