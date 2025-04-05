import { api } from "@/lib/axios";
import { CategoryResponse } from "./responses/category-response";

export interface GetCategoryByIdParams {
  storeId: string;
  categoryId: string;
}

export async function getCategoryById({
  storeId,
  categoryId,
}: GetCategoryByIdParams) {
  const response = await api.get<CategoryResponse>(
    `/stores/${storeId}/categories/${categoryId}`
  );

  return response.data;
}
