import { api } from "@/lib/axios";
import { PaginatedResponse } from "./responses/paginated-response";
import { ColorResponse } from "./responses/colors-response";

export interface GetColorsPagedListParams {
  storeId: string;
  page?: number | null;
  limit?: number | null;
  name?: string | null;
}

export async function getColorsPagedList({
  storeId,
  page,
  limit = 10,
  name,
}: GetColorsPagedListParams) {
  const response = await api.get<PaginatedResponse<ColorResponse>>(
    `stores/${storeId}/colors`,
    {
      params: {
        page,
        limit,
        name,
      },
    }
  );

  return response.data;
}
