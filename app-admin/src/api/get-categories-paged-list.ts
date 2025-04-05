import { api } from "@/lib/axios";
import { CategoryResponse } from "./responses/category-response";
import { PaginatedResponse } from "./responses/paginated-response";

export interface GetCategoriesPagedListParams {
  storeId: string;
  page?: number | null;
  limit?: number | null;
  name?: string | null;
  billboardId?: string | null;
}

export async function getCategoriesPagedList({
  page,
  limit,
  name,
  storeId,
  billboardId,
}: GetCategoriesPagedListParams) {
  const response = await api.get<PaginatedResponse<CategoryResponse>>(
    `/stores/${storeId}/categories`,
    {
      params: {
        page,
        limit,
        name,
        billboardId,
      },
    }
  );

  return response.data;
}
