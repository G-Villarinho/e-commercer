import { api } from "@/lib/axios";
import { PaginatedResponse } from "./responses/paginated-response";

export interface GetAllSizesParams {
  storeId: string;
}

export interface GetAllSizesResponse {
  id: string;
  name: string;
}

export async function getAllSizes({ storeId }: GetAllSizesParams) {
  const response = await api.get<PaginatedResponse<GetAllSizesResponse>>(
    `stores/${storeId}/sizes`,
    {
      params: {
        page: 1,
        limit: 1000,
      },
    }
  );

  return response.data.data;
}
