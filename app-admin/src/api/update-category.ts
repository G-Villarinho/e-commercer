import { api } from "@/lib/axios";

export interface CreateCategoryRequest {
    name: string;
    billboardId: string;
    storeId: string;
    categoryId: string;
}

export async function updateCategory({ name, billboardId, storeId, categoryId }: CreateCategoryRequest) {
    await api.put(`stores/${storeId}/categories/${categoryId}`, { name, billboardId });
}