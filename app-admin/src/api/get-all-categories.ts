import { api } from "@/lib/axios";
import { PaginatedResponse } from "./responses/paginated-response";

export interface GetAllCategoriesParams {
  storeId: string;
}

export interface GetAllCategoriesResponse {
  id: string;
  name: string;
}

export async function getAllCategories({ storeId }: GetAllCategoriesParams) {
  const response = await api.get<PaginatedResponse<GetAllCategoriesResponse>>(
    `stores/${storeId}/categories`,
    {
      params: {
        page: 1,
        limit: 1000,
      },
    }
  );

  return response.data.data;
}
