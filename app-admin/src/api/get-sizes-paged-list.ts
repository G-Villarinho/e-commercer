import { api } from "@/lib/axios";
import { SizeResponse } from "./responses/size-response";
import { PaginatedResponse } from "./responses/paginated-response";

export interface GetSizesPagedListQueyParams {
  storeId: string;
  page?: number | null;
  limit?: number | null;
  name?: string | null;
}

export async function getSizesPagedList({
  storeId,
  page,
  limit,
  name,
}: GetSizesPagedListQueyParams) {
  const response = await api.get<PaginatedResponse<SizeResponse>>(
    `stores/${storeId}/sizes`,
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
