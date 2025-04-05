import { api } from "@/lib/axios";
import { SizeResponse } from "./responses/size-response";

export interface GetSizeByIdParams {
    storeId: string;
    sizeId: string;
}

export async function getSizeById({ storeId, sizeId }: GetSizeByIdParams) {
    const response = await api.get<SizeResponse>(`stores/${storeId}/sizes/${sizeId}`);

    return response.data;
}