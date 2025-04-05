import { api } from "@/lib/axios";
import { StoreResponse } from "@/api/responses/store-response";

export interface GetStoreByIdParams {
    storeId: string;
}

export async function getStoreById(params: GetStoreByIdParams) {
    const response = await api.get<StoreResponse>(`/stores/${params.storeId}`);

    return response.data;
}