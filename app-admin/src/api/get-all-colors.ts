import { api } from "@/lib/axios";
import { PaginatedResponse } from "./responses/paginated-response";

export interface GetAllColorsParams {
  storeId: string;
}

export interface GetAllColorsResponse {
  id: string;
  name: string;
  hex: string;
}

export async function getAllColors({ storeId }: GetAllColorsParams) {
  const response = await api.get<PaginatedResponse<GetAllColorsResponse>>(
    `stores/${storeId}/colors`,
    {
      params: {
        page: 1,
        limit: 1000,
      },
    }
  );

  return response.data.data;
}
