import { api } from "@/lib/axios";

export interface GetAllBillboardsParams {
  storeId: string;
}

export interface GetAllBillboardsResponse {
  id: string;
  label: string;
}

export async function getAllBillboards({ storeId }: GetAllBillboardsParams) {
  const response = await api.get<{ data: GetAllBillboardsResponse[] }>(
    `/stores/${storeId}/billboards`,
    {
      params: {
        limit: 1000,
        page: 1,
      },
    }
  );

  return response.data.data;
}
