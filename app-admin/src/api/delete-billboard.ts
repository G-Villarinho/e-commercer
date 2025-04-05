import { api } from "@/lib/axios";

export interface DeleteBillboardParams {
    storeId: string;
    billboardId: string;
}

export async function  deleteBillboard({ storeId, billboardId }: DeleteBillboardParams) {
    await api.delete(`/stores/${storeId}/billboards/${billboardId}`);
}