import { api } from "@/lib/axios";

export interface DeleteCategoryParams {
    storeId: string;
    categoryId: string;
}

export async function deleteCategory({storeId, categoryId}: DeleteCategoryParams) {
    await api.delete(`/stores/${storeId}/categories/${categoryId}`);
}