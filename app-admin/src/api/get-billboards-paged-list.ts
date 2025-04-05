import { api } from "@/lib/axios";
import { BillboardResponse } from "./responses/billboard-response";
import { PaginatedResponse } from "./responses/paginated-response";

export interface GetBillboardsPagedListParams {
  storeId: string;
  page?: number | null;
  limit?: number | null;
  label?: string | null;
}

export async function getBillboardsPagedList({
  page,
  limit = 10,
  label,
  storeId,
}: GetBillboardsPagedListParams) {
  const response = await api.get<PaginatedResponse<BillboardResponse>>(
    `/stores/${storeId}/billboards`,
    {
      params: {
        page,
        limit,
        label,
      },
    }
  );

  return response.data;
}
