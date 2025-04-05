import { api } from "@/lib/axios";
import { BillboardResponse } from "./responses/billboard-response";

export interface GetBillboardParams {
    storeId: string;
    billboardId: string;
}

export async function getBillboard({ storeId, billboardId }: GetBillboardParams) {
    const response =  await api.get<BillboardResponse>(`/stores/${storeId}/billboards/${billboardId}`);

    return response.data;
}