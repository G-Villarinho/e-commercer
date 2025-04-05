import { api } from "@/lib/axios";

export interface CreateStoreRequest {
    name: string;
}

export interface CreateStoreResponse {
    storeId: string;
}

export async function createStore({ name }: CreateStoreRequest) {
    const response = await api.post<CreateStoreResponse>("/stores", { name });
    return response.data;
}  